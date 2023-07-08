#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Builds the iOS framework and application

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1
[ "$TGT" ] || TGT="0.0.0"

echo "building gomobile for iOS..."
GOARCH=arm64 time gomobile bind -o build/dist/mobile_ios_arm64/adminiServer.xcframework -target=ios admini.dev/admini/app/cmd
echo "gomobile for iOS completed successfully, building distribution..."
cd "build/dist/mobile_ios_arm64/adminiServer.xcframework"
zip --symlinks -r "../../admini_${TGT}_ios_framework.zip" .

echo "Building iOS app..."
cd $dir/../../tools/ios

rm -rf ../../build/dist/mobile_ios_app_arm64
mkdir -p ../../build/dist/mobile_ios_app_arm64

xcodegen generate --spec xcodegen.yml --project ../../build/dist/mobile_ios_app_arm64

mv Info.plist ../../build/dist/mobile_ios_app_arm64
cd ../../build/dist/mobile_ios_app_arm64

xcodebuild -project "Admini.xcodeproj" -allowProvisioningUpdates
zip -r "$dir/../../build/dist/admini_${TGT}_ios_app.zip" .

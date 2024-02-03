#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Meant to be run as part of the release process, builds desktop apps

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

docker build -f tools/desktop/Dockerfile.desktop -t admini .

rm -rf tmp/release
mkdir -p tmp/release

cd "tmp/release"

id=$(docker create admini)
docker cp $id:/dist - > ./desktop.tar
docker rm -v $id
tar -xvf "desktop.tar"
rm "desktop.tar"

mv dist/darwin_amd64/admini "admini.darwin"
mv dist/darwin_arm64/admini "admini.darwin.arm64"
mv dist/linux_amd64/admini "admini"
mv dist/windows_amd64/admini "admini.exe"
rm -rf "dist"

# darwin amd64
cp -R "../../tools/desktop/template" .

mkdir -p "./Admini.app/Contents/Resources"
mkdir -p "./Admini.app/Contents/MacOS"

cp -R "./template/darwin/Info.plist" "./Admini.app/Contents/Info.plist"
cp -R "./template/darwin/icons.icns" "./Admini.app/Contents/Resources/icons.icns"

cp "admini.darwin" "./Admini.app/Contents/MacOS/admini"

echo "signing amd64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app/Contents/MacOS/admini"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app"

cp "./template/darwin/appdmg.config.json" "./appdmg.config.json"

echo "building macOS amd64 DMG..."
appdmg "appdmg.config.json" "./admini_${TGT}_darwin_amd64_desktop.dmg"
zip -r "admini_${TGT}_darwin_amd64_desktop.zip" "./Admini.app"

# darwin arm64
cp "admini.darwin.arm64" "./Admini.app/Contents/MacOS/admini"

echo "signing arm64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app/Contents/MacOS/admini"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app"

echo "building macOS arm64 DMG..."
appdmg "appdmg.config.json" "./admini_${TGT}_darwin_arm64_desktop.dmg"
zip -r "admini_${TGT}_darwin_arm64_desktop.zip" "./Admini.app"

# macOS universal
rm "./Admini.app/Contents/MacOS/admini"
lipo -create -output "./Admini.app/Contents/MacOS/admini" admini.darwin admini.darwin.arm64

echo "signing universal desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app/Contents/MacOS/admini"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app"

echo "building macOS universal DMG..."
appdmg "appdmg.config.json" "./admini_${TGT}_darwin_all_desktop.dmg"
zip -r "admini_${TGT}_darwin_all_desktop.zip" "./Admini.app"

# linux
echo "building Linux zip..."
zip "admini_${TGT}_linux_amd64_desktop.zip" "./admini"

#windows
echo "building Windows zip..."
curl -L -o webview.dll https://github.com/webview/webview/raw/master/dll/x64/webview.dll
curl -L -o WebView2Loader.dll https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll
zip "admini_${TGT}_windows_amd64_desktop.zip" "./admini.exe" "./webview.dll" "./WebView2Loader.dll"

mkdir -p "../../build/dist"
mv "./admini_${TGT}_darwin_amd64_desktop.dmg" "../../build/dist"
mv "./admini_${TGT}_darwin_amd64_desktop.zip" "../../build/dist"
mv "./admini_${TGT}_darwin_arm64_desktop.dmg" "../../build/dist"
mv "./admini_${TGT}_darwin_arm64_desktop.zip" "../../build/dist"
mv "./admini_${TGT}_darwin_all_desktop.dmg" "../../build/dist"
mv "./admini_${TGT}_darwin_all_desktop.zip" "../../build/dist"
mv "./admini_${TGT}_linux_amd64_desktop.zip" "../../build/dist"
mv "./admini_${TGT}_windows_amd64_desktop.zip" "../../build/dist"

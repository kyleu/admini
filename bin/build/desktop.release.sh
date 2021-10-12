#!/bin/bash

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

docker build -f tools/desktop/Dockerfile -t admini .

rm -rf tmp/release
mkdir -p tmp/release

cd "tmp/release"

id=$(docker create admini)
docker cp $id:/dist - > ./desktop.tar
docker rm -v $id
tar -xvf "desktop.tar"
rm "desktop.tar"

mv dist/darwin_amd64/admini "admini.macos"
mv dist/linux_amd64/admini "admini"
mv dist/windows_amd64/admini "admini.exe"
rm -rf "dist"

# macOS
cp -R "../../tools/desktop/template" .

mkdir -p "./Admini.app/Contents/Resources"
mkdir -p "./Admini.app/Contents/MacOS"

cp -R "./template/macos/Info.plist" "./Admini.app/Contents/Info.plist"
cp -R "./template/macOS/icons.icns" "./Admini.app/Contents/Resources/icons.icns"

cp "admini.macos" "./Admini.app/Contents/MacOS/admini"

echo "signing desktop binary..."
codesign  -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app/Contents/MacOS/admini"
codesign  -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app"

cp "./template/macos/appdmg.config.json" "./appdmg.config.json"

echo "building macOS DMG..."
appdmg "appdmg.config.json" "./admini_${TGT}_macos_x86_64_desktop.dmg"
zip -r "admini_${TGT}_macos_x86_64_desktop.zip" "./Admini.app"

echo "building Linux zip..."
zip "admini_${TGT}_linux_x86_64_desktop.zip" "./admini"

echo "building Windows zip..."
curl -L -o webview.dll https://github.com/webview/webview/raw/master/dll/x64/webview.dll
curl -L -o WebView2Loader.dll https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll
zip "admini_${TGT}_windows_x86_64_desktop.zip" "./admini.exe" "./webview.dll" "./WebView2Loader.dll"

mkdir -p "../../build/dist"
mv "./admini_${TGT}_macos_x86_64_desktop.dmg" "../../build/dist"
mv "./admini_${TGT}_macos_x86_64_desktop.zip" "../../build/dist"
mv "./admini_${TGT}_linux_x86_64_desktop.zip" "../../build/dist"
mv "./admini_${TGT}_windows_x86_64_desktop.zip" "../../build/dist"

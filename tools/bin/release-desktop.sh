#!/bin/bash

set -eou pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1

#docker build -f desktop.Dockerfile -t admini .

rm -rf tmp/release
mkdir -p tmp/release

id=$(docker create admini)
docker cp $id:/dist - > ./tmp/release/desktop.tar
docker rm -v $id

cd "tmp/release"

tar -xvf "desktop.tar"
rm "desktop.tar"

cp -R "../../tools/desktop/template" .

mv dist/darwin_amd64/admini "admini.macos"
mv dist/linux_amd64/admini "admini.linux"
mv dist/windows_amd64/admini "admini.windows.exe"

rm -rf "dist"

mkdir -p "./Admini.app/Contents/Resources"
mkdir -p "./Admini.app/Contents/MacOS"

cp -R "./template/macos/Info.plist" "./Admini.app/Contents/Info.plist"
cp -R "./template/macOS/icons.icns" "./Admini.app/Contents/Resources/icons.icns"

cp "admini.macos" "./Admini.app/Contents/MacOS/admini"

cp "./template/macos/appdmg.config.json" "./appdmg.config.json"

appdmg "appdmg.config.json" "./admini_${TGT}_macos_amd64.dmg"
zip -r "admini_${TGT}_desktop_macos_amd64.zip" "./Admini.app"

mkdir -p "../../build/dist"
mv "./admini_${TGT}_macos_amd64.dmg" "../../build/dist"
mv "./admini_${TGT}_desktop_macos_amd64.zip" "../../build/dist"

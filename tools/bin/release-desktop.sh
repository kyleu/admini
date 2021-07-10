#!/bin/bash

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

if [ "$XSKIP_DESKTOP" != "true" ]
then
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

#  echo "signing desktop binary..."
#  codesign  -f --options=runtime -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app/Contents/MacOS/admini"
#  codesign  -f --options=runtime -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Admini.app"

  cp "./template/macos/appdmg.config.json" "./appdmg.config.json"

  echo "building macOS DMG..."
  appdmg "appdmg.config.json" "./admini_desktop_${TGT}_macos_x86_64.dmg"
  zip -r "admini_desktop_${TGT}_macos_x86_64.zip" "./Admini.app"

  echo "building Linux zip..."
  zip "admini_desktop_${TGT}_linux_x86_64.zip" "./admini"

  echo "building Windows zip..."
  zip "admini_desktop_${TGT}_windows_x86_64.zip" "./admini.exe"

  mkdir -p "../../build/dist"
  mv "./admini_desktop_${TGT}_macos_x86_64.dmg" "../../build/dist"
  mv "./admini_desktop_${TGT}_macos_x86_64.zip" "../../build/dist"
  mv "./admini_desktop_${TGT}_linux_x86_64.zip" "../../build/dist"
  mv "./admini_desktop_${TGT}_windows_x86_64.zip" "../../build/dist"
fi

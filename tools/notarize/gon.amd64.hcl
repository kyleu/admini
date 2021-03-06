# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_amd64_v1/admini"]
bundle_id = "com.kyleu.admini"

//notarize {
//  path = "./build/dist/admini_0.2.10_macos_x86_64_desktop.dmg"
//  bundle_id = "com.kyleu.admini"
//}

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/admini_0.2.10_macos_x86_64.dmg"
  volume_name = "Admini"
}

zip {
  output_path = "./build/dist/admini_0.2.10_macos_x86_64_notarized.zip"
}

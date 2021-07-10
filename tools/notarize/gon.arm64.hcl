source = ["./build/dist/darwin_arm64_darwin_arm64/admini"]
bundle_id = "com.kyleu.admini"

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/admini_0.0.0_macos_arm64.dmg"
  volume_name = "Admini"
}

zip {
  output_path = "./build/dist/admini_0.0.0_macos_arm64_notarized.zip"
}
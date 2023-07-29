# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_all/admini"]
bundle_id = "com.kyleu.admini"

notarize {
  path = "./build/dist/admini_0.2.20_darwin_all_desktop.dmg"
  bundle_id = "com.kyleu.admini"
}

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/admini_0.2.20_darwin_all.dmg"
  volume_name = "Admini"
}

zip {
  output_path = "./build/dist/admini_0.2.20_darwin_all_notarized.zip"
}

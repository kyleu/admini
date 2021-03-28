go-embed -input web/assets -output app/assets/assets.go

md build
md build\release

go build -o build/release github.com/kyleu/admini

git checkout app/assets/assets.go

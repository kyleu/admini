package download

type Link struct {
	URL  string `json:"url"`
	Mode string `json:"mode"`
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

type Links []*Link

func DownloadLinks() Links {
	ret := Links{}
	add := func(url string, mode string, os string, arch string) {
		ret = append(ret, &Link{URL: url, Mode: mode, OS: os, Arch: arch})
	}
	add("/foo", "desktop", "os", "arch")
	return ret
}

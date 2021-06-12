package export

type Result struct {
	Key string `json:"key"`
	Out *File  `json:"out"`
}

type Results []*Result

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"mime"
	"path/filepath"
)

func Asset(base, path string) ([]byte, string, string, error) {
	var b bytes.Buffer

	file := base + path

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, "", "", fmt.Errorf("error reading asset at ["+path+"]: %w", err)
	}

	if data != nil {
		w := gzip.NewWriter(&b)
		_, _ = w.Write(data)
		_ = w.Close()
		data = b.Bytes()
	}

	sum := md5.Sum(data)

	return data, hex.EncodeToString(sum[1:]), mime.TypeByExtension(filepath.Ext(file)), nil
}

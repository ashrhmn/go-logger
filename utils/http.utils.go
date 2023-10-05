package utils

import (
	"io"
	"net/http"
)

func HttpGet(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []byte{}, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

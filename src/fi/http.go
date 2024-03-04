package fi

import (
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(url string, headers string) (io.Reader, error) {

	URL := url + "?f=" + headers
	req, _ := http.NewRequest("GET", URL, nil)
	fmt.Println("Requesting.." + URL)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res.StatusCode)

	peabody := res.Body
	return peabody, nil
}

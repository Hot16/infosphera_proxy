package send_request

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Credentials struct {
	BaseUrl   string
	Method    string
	Headers   map[string]string
	GetParams map[string]string
}

func (c *Credentials) SendRequest() {
	req, err := http.NewRequest(c.Method, c.BaseUrl, nil)
	if err != nil {
		log.Println(err)
	}

	query_params := req.URL.Query()
	for k, v := range c.GetParams {
		query_params.Add(k, v)
	}
	

	req.URL.RawQuery = query_params.Encode()

	for k, v := range c.Headers {
		fmt.Println(k + " " + v)
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

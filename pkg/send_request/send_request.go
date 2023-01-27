package send_request

import (
	"bytes"
	"fmt"
	"infoSfera_proxy/pkg/save_file"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Credentials struct {
	BaseUrl    string
	Method     string
	Headers    map[string]string
	GetParams  map[string]string
	PostFields []byte
}

func (c *Credentials) SendRequest() {
	req, err := http.NewRequest(c.Method, c.BaseUrl, bytes.NewBuffer(c.PostFields))
	if err != nil {
		log.Println(err)
	}

	queryParams := req.URL.Query()
	for k, v := range c.GetParams {
		queryParams.Add(k, v)
	}

	req.URL.RawQuery = queryParams.Encode()

	for k, v := range c.Headers {
		fmt.Println(k + " " + v)
		req.Header.Add(k, v)
	}

	//req.Body.Read(c.PostFields)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)

	fileData := save_file.SaveFileData{
		IsRequest:  false,
		FileName:   "data-" + strconv.FormatInt(time.Now().UnixNano(), 10),
		StringData: string(body),
	}
	go fileData.SaveFile()
}

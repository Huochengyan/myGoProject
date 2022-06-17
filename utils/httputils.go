package utils

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "PostmanRuntime/7.19.0")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "b4d11fb6-11bc-4041-8a7f-fd65a2800be3,351f0ace-b5fe-4e3e-8bb2-930091954d24")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

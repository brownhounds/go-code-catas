package httpclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Resp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Payload any    `json:"payload,omitempty"`
}

type Req struct {
	Name string `json:"name"`
}

func Run(client *http.Client, url string) {
	req := Req{
		Name: "Mr Helope",
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		panic(err)
	}

	res, err := client.Post(url, "application/json", buf)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var resData Resp
	if err := json.NewDecoder(res.Body).Decode(&resData); err != nil {
		panic(err)
	}

	log.Println(resData)
}

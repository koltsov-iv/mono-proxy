package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type MonoHandler struct {
	apiKey  string
	apiUrl  string
	account string
}

func (h *MonoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/jar" {
		client := &http.Client{}
		url := fmt.Sprintf("%s/personal/statement/%s/%s", h.apiUrl, h.account, "1678844532")
		monoRequest, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			writer.WriteHeader(500)
			_, err := writer.Write([]byte(err.Error()))
			if err != nil {
				log.Fatalln(err)
			}
			return
		}
		monoRequest.Header.Set("X-Token", h.apiKey)
		resp, err := client.Do(monoRequest)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		for k, v := range resp.Header {
			writer.Header().Set(k, v[0])
		}
		_, err = io.Copy(writer, resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
	}
	http.NotFound(writer, request)
}

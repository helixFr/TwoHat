package main

import (
  "fmt"
  "github.com/valyala/fasthttp"
  "encoding/json"
)

func doRequest(url string) {
        req := fasthttp.AcquireRequest()
        resp := fasthttp.AcquireResponse()
        defer fasthttp.ReleaseRequest(req)
        defer fasthttp.ReleaseResponse(resp)
        jsonContent := map[string]string{"text": "test 1"}
        jsonReq, _ := json.Marshal(jsonContent)

        req.Header.SetContentType("application/json")
        req.Header.SetMethod("POST")
        req.SetBody(jsonReq)
        req.SetRequestURI(url)

        fasthttp.Do(req, resp)
        fmt.Println(resp)
}

func main() {
  doRequest("http://localhost:8080/")
}

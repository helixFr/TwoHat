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

  fmt.Println(req)
  fmt.Println(resp)
}

func main() {
  fmt.Println(":5050")
  doRequest("http://localhost:5050/")
}

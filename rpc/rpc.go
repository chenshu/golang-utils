package rpc

import (
    "io"
    "bytes"
    "strings"
    "net/url"
    "net/http"
    "mime/multipart"
)

type Client struct {
    *http.Client
}

func (r Client) PostWithForm(req_url string, headers map[string][]string, params map[string][]string) (resp *http.Response, err error) {
    var body string = url.Values(params).Encode()
    headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
    return r.Post(req_url, headers, strings.NewReader(body), int64(len(body)))
}

func (r Client) PostWithData(req_url string, headers map[string][]string, body io.Reader, content_length int64) (resp *http.Response, err error) {
    return r.Post(req_url, headers, body, content_length)
}

func (r Client) PostWithMultiPartData(req_url string, headers map[string][]string, params map[string][]string, body io.Reader, filename string) (resp *http.Response, err error) {
    var buffer *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
    var writer *multipart.Writer = multipart.NewWriter(buffer)
    for k, v := range params {
        for _, field := range v {
            err1 := writer.WriteField(k, field)
            if err1 != nil {
                err = err1
                return
            }
        }
    }
    w, err := writer.CreateFormFile("file", filename)
    if err != nil {
        return
    }
    _, err = io.Copy(w, body)
    if err != nil {
        return
    }
    writer.Close()
    headers["Content-Type"] = []string{writer.FormDataContentType()}
    return r.Post(req_url, headers, buffer, int64(buffer.Len()))
}

func (r Client) Post(req_url string, headers map[string][]string, body io.Reader, content_length int64) (resp *http.Response, err error) {
    req, err := http.NewRequest("POST", req_url, body)
    if err != nil {
        return
    }
    for k, v := range headers {
        for _, h := range v {
            req.Header.Set(k, h)
        }
    }
    req.ContentLength = content_length
    return r.Do(req)
}

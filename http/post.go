package http

import (
    "io"
    "os"
    "fmt"
    "bytes"
    "errors"
    "net/url"
    "net/http"
    "io/ioutil"
    "mime/multipart"
)

func Post(req_url string, params map[string]string) (body string, err error) {
    value := url.Values{}
    for k, v := range params {
        value.Add(k, v)
    }

    var buffer *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
    buffer.WriteString(value.Encode())

    var client *http.Client = new(http.Client)
    var req *http.Request
    req, err = http.NewRequest("POST", req_url, buffer)
    if err != nil {
        return
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; rv:17.0) Gecko/20100101 Firefox/17.0")
    var resp *http.Response
    resp, err = client.Do(req)
    if err != nil {
        return
    }
    defer resp.Body.Close()
    resp_body, err := ioutil.ReadAll(resp.Body)
    body = string(resp_body)
    if resp.StatusCode != 200 {
        err = errors.New(fmt.Sprintf("%d", resp.StatusCode))
        return
    }
    return
}

func PostFile(req_url, path, filename string, start, end, filesize int, params map[string]string) (body string, err error) {
    var buffer *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
    var writer *multipart.Writer = multipart.NewWriter(buffer)

    for k, v := range params {
        err = writer.WriteField(k, v)
        if err != nil {
            return
        }
    }

    var file *os.File
    if file, err = os.OpenFile(path, os.O_RDONLY, 0644); err != nil {
        return
    }
    defer file.Close()

    w, err := writer.CreateFormFile("file", filename)
    if err != nil {
        return
    }

    var size = end - start + 1
    if size == filesize {
        _, err = io.Copy(w, file)
        if err != nil {
            return
        }
    } else {
        var section []byte = make([]byte, size)
        var n int
        n, err = file.ReadAt(section, int64(start))
        if err != nil || n != size {
            return
        }
        var buf *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
        buf.Write(section)
        _, err = io.Copy(w, buf)
        if err != nil {
            return
        }
    }
    writer.Close()

    var client *http.Client = new(http.Client)
    var req *http.Request
    req, err = http.NewRequest("POST", req_url, buffer)
    if err != nil {
        return
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; rv:17.0) Gecko/20100101 Firefox/17.0")
    var resp *http.Response
    resp, err = client.Do(req)
    if err != nil {
        return
    }
    defer resp.Body.Close()
    resp_body, err := ioutil.ReadAll(resp.Body)
    body = string(resp_body)
    if resp.StatusCode != 200 {
        err = errors.New(fmt.Sprintf("%d", resp.StatusCode))
        return
    }
    return
}

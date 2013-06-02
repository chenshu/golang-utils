package hash

import (
    "io"
    "os"
    "fmt"
    "crypto/md5"
)

func Md5file(path string) (hash_hex string, err error) {
    var file *os.File
    if file, err = os.OpenFile(path, os.O_RDONLY, 0644); err != nil {
        return
    }
    defer file.Close()

    h := md5.New()
    //io.Copy(h, file)
    var data []byte = make([]byte, 8192)
    var size int
    for {
        size, err = file.Read(data)
        if err != nil {
            break
        }
        h.Write(data[:size])
        h.Sum(nil)
    }
    if err != nil && err != io.EOF {
        return
    }
    hash_hex = fmt.Sprintf("%x", h.Sum(nil))
    err = nil
    return
}

func Md5string(str string) (hash_hex string, err error) {
    h := md5.New()
    io.WriteString(h, str)
    hash_hex = fmt.Sprintf("%x", h.Sum(nil))
    return
}

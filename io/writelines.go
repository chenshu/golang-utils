package io

import (
    "os"
    "bufio"
    "strings"
)

func WriteLines(lines []string, path string) (err error) {
    var file *os.File
    if file, err = os.Create(path); err != nil {
        return
    }
    defer file.Close()

    var writer *bufio.Writer = bufio.NewWriter(file)
    var content string = strings.Join(lines, "\n")
    _, err = writer.WriteString(content + "\n")
    if err != nil {
        return
    }
    err = writer.Flush()
    return
}

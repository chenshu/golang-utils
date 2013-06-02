package hash

import "testing"

func TestMd5string(t *testing.T) {
    const in, out = "hello world", "5eb63bbbe01eeed093cb22bb8f5acdc3"
    if x, err := Md5string(in); x != out || err != nil {
        t.Errorf("md5sum(%v) = %v, want %v", in, x, out)
    }
}

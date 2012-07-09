package parser

import (
    "fmt"
)

func Byte(b byte) Parser {
    return func(r *Reader) (interface{}, *Reader, error) {
        bytes, rd, err := r.read(1)
        if err != nil {
            return nil, r, fmt.Errorf("Failed to parse %#v, read failed: %s", b, err)
        }

        if bytes[0] == b {
            return b, rd, err
        }

        return nil, r, fmt.Errorf("Failed to read %#v, got %#v", b, bytes[0])
    }
}

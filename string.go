package parser

import (
    "fmt"
)

func String(s string) Parser {
    length := len(s)
    return func(r *Reader) (interface{}, *Reader, error) {
        bytes, r2, err := r.read(length)
        if err != nil {
            return nil, r, fmt.Errorf("Failed to parse %v, read failed: %s", s, err)
        }

        parsed := string(bytes)
        if parsed == s {
            return parsed, r2, nil
        }

        return nil, r, fmt.Errorf("Failed to parse %v, got %v instead", s, parsed)
    }
}

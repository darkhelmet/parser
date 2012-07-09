package parser

import (
    "io"
)

type Parser func(*Reader) (interface{}, *Reader, error)

func Parse(parser Parser, r io.Reader) (interface{}, error) {
    i, _, err := parser(NewReader(r))
    return i, err
}

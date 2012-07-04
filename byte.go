package parser

import (
    "fmt"
)

type byteParser struct {
    byte byte
}

func (p *byteParser) String() string {
    return fmt.Sprintf("%#v", p.byte)
}

func (p *byteParser) Parse(reader Reader) (error, bool, interface{}, Reader) {
    bytes, err := reader.Peek(1)
    if err != nil {
        return fmt.Errorf("Failed to parse %s, read failed: %s", p, err), false, nil, reader
    }
    if bytes[0] == p.byte {
        b, err := reader.ReadByte()
        if err != nil {
            return fmt.Errorf("Failed to parse %s, read failed: %s", p, err), false, nil, reader
        }
        return nil, true, b, reader
    }

    return fmt.Errorf("Failed to read %s", p), false, nil, reader
}

func Byte(b byte) Parser {
    return &byteParser{byte: b}
}

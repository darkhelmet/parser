package parser

import (
    "fmt"
)

type stringParser struct {
    str string
}

func (p *stringParser) String() string {
    return fmt.Sprintf("%v", p.str)
}

func (p *stringParser) Parse(reader Reader) (error, bool, interface{}, Reader) {
    input, err := reader.Peek(len(p.str))
    if err != nil {
        return fmt.Errorf("Failed to parse %s, read failed: %s", p, err), false, nil, reader
    }

    parsed := string(input)
    if parsed == p.str {
        if _, err := reader.Read(input); err != nil {
            return fmt.Errorf("Failed to parse %s, read failed: %s", p, err), false, nil, reader
        }
        return nil, true, parsed, reader
    }
    return fmt.Errorf("Failed to parse %s", p), false, nil, reader
}

func String(s string) Parser {
    return &stringParser{str: s}
}

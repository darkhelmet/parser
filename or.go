package parser

import (
    "fmt"
    "strings"
)

func Or(parsers ...Parser) Parser {
    return func(r *Reader) (interface{}, *Reader, error) {
        errors := make([]string, 0, len(parsers))
        for _, parser := range parsers {
            result, rd, err := parser(r)
            if err == nil {
                return result, rd, nil
            }
            errors = append(errors, err.Error())
        }
        return nil, r, fmt.Errorf("Failed parsing Or: [%s]", strings.Join(errors, "; "))
    }
}

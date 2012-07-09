package parser

import (
    "fmt"
)

func And(parsers ...Parser) Parser {
    return func(r *Reader) (interface{}, *Reader, error) {
        results := make([]interface{}, 0, len(parsers))
        rd := r
        for _, parser := range parsers {
            var (
                err    error
                result interface{}
            )
            result, rd, err = parser(rd)
            if err != nil {
                return nil, r, fmt.Errorf("Failed to parse: %s", err)
            }
            results = append(results, result)
        }
        return results, rd, nil
    }
}

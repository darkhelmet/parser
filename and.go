package parser

import (
    "fmt"
)

type andParser struct {
    parsers []Parser
}

func (p *andParser) String() string {
    return fmt.Sprintf("%v", p.parsers)
}

func (p *andParser) Parse(reader Reader) (error, bool, interface{}, Reader) {
    results := make([]interface{}, 0, len(p.parsers))
    rd := reader
    for _, parser := range p.parsers {
        err, consumed, result, rd := parser.Parse(rd)
        if consumed {
            if err != nil {
                return fmt.Errorf("Parser %v consumed input, but returned error: %s", parser, err), true, nil, rd
            }
            results = append(results, result)
        }
    }
    return nil, true, results, rd
}

func And(parsers ...Parser) Parser {
    return &andParser{parsers}
}

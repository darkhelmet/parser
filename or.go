package parser

import (
    "fmt"
)

type orParser struct {
    parsers []Parser
}

func (p *orParser) String() string {
    return fmt.Sprintf("%#v", p.parsers)
}

func (p *orParser) Parse(reader Reader) (error, bool, interface{}, Reader) {
    for _, parser := range p.parsers {
        err, consumed, result, rd := parser.Parse(reader)
        if consumed {
            if err != nil {
                // Bail if the parser error'd but consumed input
                return fmt.Errorf("Parser %v consumed input, but returned error: %s", parser, err), true, nil, rd
            }
            return err, consumed, result, rd
        }
    }
    return fmt.Errorf("Failed parsing Or, none of %s parsed", p), false, nil, reader
}

func Or(parsers ...Parser) Parser {
    return &orParser{parsers: parsers}
}

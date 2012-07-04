package parser

import (
    "fmt"
)

type betweenParser struct {
    left, right string
    inside      Parser
}

func (p *betweenParser) String() string {
    return fmt.Sprintf("%v%s%v", p.left, p.inside, p.right)
}

func (p *betweenParser) Parse(reader Reader) (error, bool, interface{}, Reader) {
    err, consumed, results, rd := And(String(p.left), p.inside, String(p.right)).Parse(reader)
    if err != nil {
        return err, consumed, results, rd
    }
    return err, consumed, results.([]interface{})[1], rd
}

func Between(left, right string, inside Parser) Parser {
    return &betweenParser{left: left, right: right, inside: inside}
}

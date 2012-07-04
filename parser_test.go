package parser_test

import (
    "bufio"
    P "github.com/darkhelmet/parser"
    . "launchpad.net/gocheck"
    "strings"
    "testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func BufferedString(s string) *bufio.Reader {
    return bufio.NewReader(strings.NewReader(s))
}

func (s *S) TestStringFailure(c *C) {
    p := P.String("null")
    err, consumed, value, _ := p.Parse(BufferedString("facebook"))
    c.Assert(err, NotNil)
    c.Assert(consumed, Equals, false)
    c.Assert(value, Equals, nil)
}

func (s *S) TestStringSuccess(c *C) {
    v := "null"
    p := P.String(v)
    err, consumed, value, _ := p.Parse(BufferedString(v))
    c.Assert(err, IsNil)
    c.Assert(consumed, Equals, true)
    c.Assert(value, FitsTypeOf, v)
    c.Assert(value, Equals, "null")
}

func (s *S) TestOrFailure(c *C) {
    p := P.Or(P.String("batman"), P.String("robin"))
    err, consumed, value, _ := p.Parse(BufferedString("joker"))
    c.Assert(err, NotNil)
    c.Assert(consumed, Equals, false)
    c.Assert(value, Equals, nil)
}

func (s *S) TestOrSuccessFirst(c *C) {
    p := P.Or(P.String("batman"), P.String("robin"))
    err, consumed, value, _ := p.Parse(BufferedString("batman"))
    c.Assert(err, IsNil)
    c.Assert(consumed, Equals, true)
    c.Assert(value, Equals, "batman")
}

func (s *S) TestOrSuccessRest(c *C) {
    p := P.Or(P.String("batman"), P.String("robin"))
    err, consumed, value, _ := p.Parse(BufferedString("robin"))
    c.Assert(err, IsNil)
    c.Assert(consumed, Equals, true)
    c.Assert(value, Equals, "robin")
}

func (s *S) TestByteFailure(c *C) {
    p := P.Byte(':')
    err, consumed, value, _ := p.Parse(BufferedString("batman"))
    c.Assert(err, NotNil)
    c.Assert(consumed, Equals, false)
    c.Assert(value, Equals, nil)
}

func (s *S) TestByteSuccess(c *C) {
    p := P.Byte(':')
    err, consumed, value, _ := p.Parse(BufferedString(":symbol"))
    c.Assert(err, IsNil)
    c.Assert(consumed, Equals, true)
    c.Assert(value, Equals, byte(':'))
}

var (
    lBrace = P.Byte('{')
    rBrace = P.Byte('}')
    colon  = P.Byte(':')
    quote  = `"`
)

type JsonObject struct{}

func (jo *JsonObject) Parse(reader P.Reader) (err error, consumed bool, result interface{}, rd P.Reader) {
    key := P.Between(quote, quote, P.String("key"))
    value := P.Between(quote, quote, P.String("value"))
    p := P.And(lBrace, key, colon, value, rBrace)
    err, consumed, result, rd = p.Parse(reader)
    if err != nil {
        return
    }

    res := result.([]interface{})
    return nil, true, map[string]interface{}{res[1].(string): res[3]}, rd
}

func (s *S) TestBasicJson(c *C) {
    p := &JsonObject{}
    err, consumed, value, _ := p.Parse(BufferedString(`{"key":"value"}`))
    c.Assert(err, IsNil)
    c.Assert(consumed, Equals, true)
    c.Assert(value, DeepEquals, map[string]interface{}{"key": "value"})
}

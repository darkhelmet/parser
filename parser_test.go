package parser_test

import (
    P "github.com/darkhelmet/parser"
    . "launchpad.net/gocheck"
)

func (s *S) TestStringFailure(c *C) {
    value, err := P.Parse(P.String("null"), srd("facebook"))
    c.Assert(err, NotNil)
    c.Assert(value, Equals, nil)
}

func (s *S) TestStringSuccess(c *C) {
    v := "null"
    value, err := P.Parse(P.String(v), srd(v))
    c.Assert(err, IsNil)
    c.Assert(value, FitsTypeOf, v)
    c.Assert(value, Equals, "null")
}

func (s *S) TestOrFailure(c *C) {
    p := P.Or(P.String("batman"), P.String("robin"))
    value, err := P.Parse(p, srd("joker"))
    c.Assert(err, NotNil)
    c.Assert(value, Equals, nil)
}

func (s *S) TestOrSuccessFirst(c *C) {
    p := P.Or(P.String("batman"), P.String("robin"))
    value, err := P.Parse(p, srd("batman"))
    c.Assert(err, IsNil)
    c.Assert(value, Equals, "batman")
}

func (s *S) TestOrSuccessRest(c *C) {
    p := P.Or(P.String("batman"), P.String("robin"))
    value, err := P.Parse(p, srd("robin"))
    c.Assert(err, IsNil)
    c.Assert(value, Equals, "robin")
}

func (s *S) TestByteFailure(c *C) {
    value, err := P.Parse(P.Byte(':'), srd("batman"))
    c.Assert(err, NotNil)
    c.Assert(value, Equals, nil)
}

func (s *S) TestByteSuccess(c *C) {
    value, err := P.Parse(P.Byte(':'), srd(":symbol"))
    c.Assert(err, IsNil)
    c.Assert(value, Equals, byte(':'))
}

var (
    lBrace = P.Byte('{')
    rBrace = P.Byte('}')
    colon  = P.Byte(':')
    quote  = `"`
)

func JsonObject(r *P.Reader) (result interface{}, rd *P.Reader, err error) {
    key := P.Between(quote, quote, P.String("key"))
    value := P.Between(quote, quote, P.String("value"))
    p := P.And(lBrace, key, colon, value, rBrace)

    result, rd, err = p(r)
    if err != nil {
        return
    }

    res := result.([]interface{})
    return map[string]interface{}{res[1].(string): res[3]}, rd, nil
}

func (s *S) TestBasicJson(c *C) {
    value, err := P.Parse(JsonObject, srd(`{"key":"value"}`))
    c.Assert(err, IsNil)
    c.Assert(value, DeepEquals, map[string]interface{}{"key": "value"})
}

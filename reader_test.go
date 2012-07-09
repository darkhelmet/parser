package parser

import (
    "io"
    . "launchpad.net/gocheck"
    "strings"
    "testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func srd(s string) io.Reader {
    return strings.NewReader(s)
}

func (s *S) TestReadZero(c *C) {
    r := NewReader(srd("batman"))
    c.Assert(len(r.buf), Equals, 0)
    c.Assert(r.cur, Equals, 0)
}

func (s *S) TestRead(c *C) {
    r := NewReader(srd("batman"))

    data, r2, err := r.read(3)
    c.Assert(err, IsNil)
    c.Assert(string(data), Equals, "bat")
    c.Assert(len(r2.buf), Equals, 6)
    c.Assert(r2.cur, Equals, 3)
    c.Assert(r.cur, Equals, 0)

    data2, r3, err := r2.read(3)
    c.Assert(err, IsNil)
    c.Assert(string(data2), Equals, "man")
    c.Assert(len(r3.buf), Equals, 6)
    c.Assert(r3.cur, Equals, 6)

    _, _, err = r3.read(3)
    c.Assert(err, NotNil)
}

func (s *S) TestOverRead(c *C) {
    r := NewReader(srd("joker"))
    data, r2, err := r.read(6)
    c.Assert(data, IsNil)
    c.Assert(r2, DeepEquals, r)
    c.Assert(err, FitsTypeOf, ioerror{})
}

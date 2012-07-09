package parser_test

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

package parser

type Reader interface {
    Peek(n int) ([]byte, error)
    Read(p []byte) (int, error)
    ReadByte() (byte, error)
}

type Parser interface {
    Parse(Reader) (error, bool, interface{}, Reader)
}

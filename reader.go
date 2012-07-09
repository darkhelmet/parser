package parser

import (
    "io"
)

const readSize = 1024

type ioerror struct {
    err error
}

func (ioe ioerror) String() string {
    return ioe.err.Error()
}

func (ioe ioerror) Error() string {
    return ioe.err.Error()
}

type Reader struct {
    r   io.Reader
    buf []byte
    cur int
}

func NewReader(r io.Reader) *Reader {
    return &Reader{r: r}
}

func (r *Reader) read(n int) ([]byte, *Reader, error) {
    for r.needToRead(n) {
        fresh := make([]byte, 1024)
        c, err := r.r.Read(fresh)
        r.buf = append(r.buf, fresh[0:c]...)
        if err != nil {
            // Don't worry about cleanup on an ioerror
            return nil, r, ioerror{err}
        }
    }
    low, high := r.cur, r.cur+n
    bytes := r.buf[low:high]
    return bytes, r.clone(high), nil
}

func (r *Reader) needToRead(desired int) bool {
    return r.cur+desired > len(r.buf)
}

func (r *Reader) clone(cur int) *Reader {
    return &Reader{
        r:   r.r,
        buf: r.buf,
        cur: cur,
    }
}

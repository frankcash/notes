Encrypting the Internet with Go

Filippo Valsorda

##  What is TLS?

- TLS is a shared key encryption model.
- TLS 1.3 - Brings lots of changes, handshake down to one
- `crypto/tls += 1.3`

## `crypto/tls`

- Simplified using TLS code
- [0 Round Trip Time Primer](https://blog.cloudflare.com/introducing-0-rtt/)
- Buffering is trash, TLS Lib used to have to read through all the data before reaching fin
- TLS 1.3 is about `safe defaults`
- `golang.org/cl/33776`
- Very thoroughly tested
- `github.com/cloudfare/gokeyless/client`
- cloudfare uses crypto primitives that are implemented in ASM (amd64) only
- do not use default timouts on `net/http`, stuff will just not survive
- Go 1.8 + fixes some timeouts
- Should keep an eye on open connections
- `panick(nil)` (why)

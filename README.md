# ipfixmessage


[![GoDoc](https://godoc.org/github.com/adrianuswarmenhoven/ipfixmessage?status.svg)](https://godoc.org/github.com/adrianuswarmenhoven/ipfixmessage)

Full implementation of IPFIX in Go. RFC7011 and RFC7012 (basiclist, subtemplatelist, subtemplatemultilist)

Before compiling, use go generate to pull in some field id's.

Still WIP, but correctly creating and parsing IPFIX messages.

TODO:
    - Sessions (IPFIX has some knowledge about sessions)
    - Nice marshalling of structs
    - Examples
# go-capnproto

## Getting started
See: https://github.com/capnproto/go-capnproto2
```
$ go get -u -t zombiezen.com/go/capnproto2/...
$ go test -v zombiezen.com/go/capnproto2/...
```
Next: https://github.com/capnproto/go-capnproto2/wiki/Getting-Started

## Create schema
```
mkdir books
cd books

vi books.capnp

capnp -I /home/andi/go/src/zombiezen.com/go/capnproto2/std compile books.capnp -o go

# Now we have books.capnp.go :-)
```
Next: [main.go](main.go)

## Build go-capnproto
```
go build
```

## Run go-capnproto
```
./go-capnproto

# Now we have a capnproto encoded file (`book.msg`) containing schema and data
$ capnp decode --help
Usage: capnp decode [<option>...] <schema-file> <type>

Decodes one or more encoded Cap'n Proto messages as text.  The messages have
root type <type> defined in <schema-file>.  Messages are read from standard
input and by default are expected to be in standard Cap'n Proto serialization
format.

Options:
    -I<dir>, --import-path=<dir>
        Add <dir> to the list of directories searched for non-relative imports
        (ones that start with a '/').
    --flat
        Interpret the input as one large single-segment message rather than a
        stream in standard serialization format.  (Rarely used.)
    --no-standard-import
        Do not add any default import paths; use only those specified by -I.
        Otherwise, typically /usr/include and /usr/local/include are added by
        default.
    -p, --packed
        Expect the input to be packed using standard Cap'n Proto packing, which
        deflates zero-valued bytes.  (This reads messages written with
        capnp::writePackedMessage*() from <capnp/serialize-packed.h>.  Do not
        use this for messages written with capnp::writeMessage*() from
        <capnp/serialize.h>.)
    --quiet
        Do not print warning messages about the input being in the wrong format.
        Use this if you find the warnings are wrong (but also let us know so we
        can improve them).
    --short
        Print in short (non-pretty) format.  Each message will be printed on one
        line, without using whitespace to improve readability.
    --verbose
        Log informational messages to stderr; useful for debugging.
    --version
        Print version information and exit.
    --help
        Display this help text and exit.

$ cat book.msg | capnp decode books/books.capnp Book
(title = "War and Peace", pageCount = 1440)
```

## TODOs
* Is schema included in book.msg? How to use 'capnp decode' without schema-file?

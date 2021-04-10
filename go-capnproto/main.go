package main

import (
	"bytes"
	"io/ioutil"
	"log"

	capnp "zombiezen.com/go/capnproto2"

	"github.com/eg5846/getting-started-with-capnproto/go-capnproto/books"
)

func encodeBookMessage(title string, pageCount int32) ([]byte, error) {
	// Make a brand new empty message.  A Message allocates Cap'n Proto structs.
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	// Create a new Book struct.  Every message must have a root struct.
	book, err := books.NewRootBook(seg)
	if err != nil {
		return nil, err
	}
	if err := book.SetTitle(title); err != nil {
		return nil, err
	}
	book.SetPageCount(pageCount)

	// Encode the message.
	var buf bytes.Buffer
	err = capnp.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decodeBookMessage(data []byte) (string, int32, error) {
	// Decode the message.
	r := bytes.NewReader(data)
	msg, err := capnp.NewDecoder(r).Decode()
	if err != nil {
		return "", 0, err
	}

	// Extract the root struct from the message.
	book, err := books.ReadRootBook(msg)
	if err != nil {
		return "", 0, err
	}

	// Access fields from the struct.
	title, err := book.Title()
	if err != nil {
		return "", 0, err
	}
	pageCount := book.PageCount()

	return title, pageCount, nil
}

func main() {
	title := "War and Peace"
	pageCount := int32(1440)

	log.Printf("Encoding book message: (title: '%s', pageCount: %d) ...", title, pageCount)
	data, err := encodeBookMessage(title, pageCount)
	if err != nil {
		log.Fatalf("Encoding book message: (title: '%s', pageCount: %d) failed: %s", title, pageCount, err)
	}

	log.Printf("Book message: %#v, len: %d", data, len(data))

	filePath := "book.msg"
	log.Printf("Writing book message to %s ...", filePath)
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		log.Fatalf("Writing book message to %s failed: %s", filePath, err)
	}

	log.Print("Decoding book message ...")
	title, pageCount, err = decodeBookMessage(data)
	if err != nil {
		log.Printf("Decoding book message failed: %s", err)
	}

	log.Printf("Decoded book message: (title: '%s', pageCount: %d)", title, pageCount)
}

package main

import (
	"encoding/json"
	"testing"

	"github.com/mussyaroslav/home_work_basic/hw09_serialize/pb"
	"google.golang.org/protobuf/proto"
)

func TestJSONSerialization(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling",
		Year:   1997,
		Size:   223,
		Rate:   4.46,
	}

	jsonData, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Ошибка при сериализации в JSON: %v", err)
	}

	var unmarshaledBook Book
	err = json.Unmarshal(jsonData, &unmarshaledBook)
	if err != nil {
		t.Fatalf("Ошибка при десериализации JSON: %v", err)
	}

	if book != unmarshaledBook {
		t.Errorf("Ожидали %+v, но получили %+v", book, unmarshaledBook)
	}
}

func TestProtobufSerialization(t *testing.T) {
	protoBook := &pb.Book{
		Id:     1,
		Title:  "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling",
		Year:   1997,
		Size:   223,
		Rate:   4.46,
	}

	data, err := proto.Marshal(protoBook)
	if err != nil {
		t.Fatalf("Ошибка при сериализации в Protobuf: %v", err)
	}

	newProtoBook := new(pb.Book)
	err = proto.Unmarshal(data, newProtoBook)
	if err != nil {
		t.Fatalf("Ошибка при десериализации Protobuf: %v", err)
	}

	if protoBook.String() != newProtoBook.String() {
		t.Errorf("Ожидали %+v, но получили %+v", protoBook, newProtoBook)
	}
}

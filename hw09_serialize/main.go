package main

import (
	"encoding/json"
	"fmt"

	"github.com/mussyaroslav/home_work_basic/hw09_serialize/pb"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int32   `json:"id" protobuf:"varint,1,opt,name=id,proto3"`
	Title  string  `json:"title" protobuf:"bytes,2,opt,name=title,proto3"`
	Author string  `json:"author" protobuf:"bytes,3,opt,name=author,proto3"`
	Year   int32   `json:"year" protobuf:"varint,4,opt,name=year,proto3"`
	Size   int32   `json:"size" protobuf:"varint,5,opt,name=size,proto3"`
	Rate   float64 `json:"rate" protobuf:"fixed64,6,opt,name=rate,proto3"`
}

func (b Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(b),
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux)
}

func (b *Book) Reset() {
	*b = Book{}
}

func (b *Book) String() string {
	return fmt.Sprintf("Book{ID: %d, Title: %s, Author: %s, Year: %d, Size: %d, Rate: %.2f}",
		b.ID, b.Title, b.Author, b.Year, b.Size, b.Rate)
}

func (b *Book) ProtoMessage() {}

func SerializeBooksToJSON(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func DeserializeBooksFromJSON(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	return books, err
}

type Books struct {
	Books []Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books"`
}

func main() {
	book1 := Book{
		ID:     1,
		Title:  "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling",
		Year:   1997,
		Size:   223,
		Rate:   4.46,
	}
	book2 := Book{
		ID:     2,
		Title:  "Harry Potter and the Chamber of Secrets",
		Author: "J.K. Rowling",
		Year:   1998,
		Size:   251,
		Rate:   4.41,
	}

	books := []Book{book1, book2}

	// Сериализация массива книг в JSON
	jsonData, err := SerializeBooksToJSON(books)
	if err != nil {
		fmt.Println("Ошибка при сериализации в JSON:", err)
		return
	}
	fmt.Println("JSON данные:", string(jsonData))

	unmarshaledBooks, err := DeserializeBooksFromJSON(jsonData)
	if err != nil {
		fmt.Println("Ошибка при десериализации JSON:", err)
		return
	}
	fmt.Printf("Десериализованные книги: %+v\n", unmarshaledBooks)

	// Пример работы с Protobuf
	protoBooks := &pb.Books{
		Books: []*pb.Book{
			{
				Id:     1,
				Title:  "Harry Potter and the Philosopher's Stone",
				Author: "J.K. Rowling",
				Year:   1997,
				Size:   223,
				Rate:   4.46,
			},
			{
				Id:     2,
				Title:  "Harry Potter and the Chamber of Secrets",
				Author: "J.K. Rowling",
				Year:   1998,
				Size:   251,
				Rate:   4.41,
			},
		},
	}

	protoData, err := proto.Marshal(protoBooks)
	if err != nil {
		fmt.Println("Ошибка при сериализации в Protobuf:", err)
		return
	}

	newProtoBooks := new(pb.Books)
	err = proto.Unmarshal(protoData, newProtoBooks)
	if err != nil {
		fmt.Println("Ошибка при десериализации Protobuf:", err)
		return
	}

	fmt.Println("Protobuf данные:", newProtoBooks)
	fmt.Println("Protobuf байты:", protoData)
}

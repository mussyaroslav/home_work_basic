package main

import (
	"fmt"
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) Rate() float64 {
	return b.rate
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

type ComparisonMather int

const (
	ByYear ComparisonMather = iota
	BySize
	ByRate
)

type BookComparator struct {
	comparison ComparisonMather
}

func NewBookComparator(comparison ComparisonMather) *BookComparator {
	return &BookComparator{comparison: comparison}
}

func (bc *BookComparator) Compare(book1, book2 *Book) bool {
	switch bc.comparison {
	case ByYear:
		return book1.Year() > book2.Year()
	case BySize:
		return book1.Size() > book2.Size()
	case ByRate:
		return book1.Rate() > book2.Rate()
	default:
		return false
	}
}

func main() {
	book1 := Book{
		1,
		"Harry Potter and the Philosopher's Stone",
		"J.K. Rowling",
		1997,
		223,
		4.46,
	}
	book2 := Book{
		2,
		"Harry Potter and the Chamber of Secrets",
		"J.K. Rowling",
		1998,
		251,
		4.41,
	}

	comparator := NewBookComparator(0)
	fmt.Println("Сравнение по году:", comparator.Compare(&book1, &book2))

	comparator = NewBookComparator(1)
	fmt.Println("Сравнение по страницам:", comparator.Compare(&book1, &book2))

	comparator = NewBookComparator(2)
	fmt.Println("Сравнение по рейтингу:", comparator.Compare(&book1, &book2))
}

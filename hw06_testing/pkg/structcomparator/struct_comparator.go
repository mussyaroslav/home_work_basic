package structcomparator

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float64
}

type BookComparator struct {
	comparison func(book1, book2 *Book) bool
}

func NewBookComparator(comparison func(book1, book2 *Book) bool) *BookComparator {
	return &BookComparator{comparison: comparison}
}

func (bc *BookComparator) Compare(book1, book2 *Book) bool {
	return bc.comparison(book1, book2)
}

func ByYear(book1, book2 *Book) bool {
	return book1.Year > book2.Year
}

func BySize(book1, book2 *Book) bool {
	return book1.Size > book2.Size
}

func ByRate(book1, book2 *Book) bool {
	return book1.Rate > book2.Rate
}

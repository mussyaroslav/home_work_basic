package structcomparator

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float64
}

func (b *Book) ValID() int {
	return b.ID
}

func (b *Book) SetID(id int) {
	b.ID = id
}

func (b *Book) ValTitle() string {
	return b.Title
}

func (b *Book) SetTitle(title string) {
	b.Title = title
}

func (b *Book) ValAuthor() string {
	return b.Author
}

func (b *Book) SetAuthor(author string) {
	b.Author = author
}

func (b *Book) ValYear() int {
	return b.Year
}

func (b *Book) SetYear(year int) {
	b.Year = year
}

func (b *Book) ValSize() int {
	return b.Size
}

func (b *Book) SetSize(size int) {
	b.Size = size
}

func (b *Book) ValRate() float64 {
	return b.Rate
}

func (b *Book) SetRate(rate float64) {
	b.Rate = rate
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
		return book1.ValYear() > book2.ValYear()
	case BySize:
		return book1.ValSize() > book2.ValSize()
	case ByRate:
		return book1.ValRate() > book2.ValRate()
	default:
		return false
	}
}

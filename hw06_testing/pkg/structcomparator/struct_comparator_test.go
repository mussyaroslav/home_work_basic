package structcomparator

import (
	"testing"
)

func TestBookComparator(t *testing.T) {
	book1 := &Book{
		ID:     1,
		Title:  "Гарри Поттер и философский камень",
		Author: "Дж. К. Роулинг",
		Year:   1997,
		Size:   223,
		Rate:   4.46,
	}
	book2 := &Book{
		ID:     2,
		Title:  "Гарри Поттер и тайная комната",
		Author: "Дж. К. Роулинг",
		Year:   1998,
		Size:   251,
		Rate:   4.41,
	}

	tests := []struct {
		name       string
		book1      *Book
		book2      *Book
		comparison ComparisonMather
		expected   bool
	}{
		{"По Году - Книга1 новее", book1, book2, ByYear, false},
		{"По Году - Книга2 новее", book2, book1, ByYear, true},
		{"По Страницам - Книга1 больше", book1, book2, BySize, false},
		{"По Страницам - Книга2 больше", book2, book1, BySize, true},
		{"По Рейтингу - Книга1 выше", book1, book2, ByRate, true},
		{"По Рейтингу - Книга2 выше", book2, book1, ByRate, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comparator := NewBookComparator(tt.comparison)
			result := comparator.Compare(tt.book1, tt.book2)
			if result != tt.expected {
				t.Errorf("Сравнение неудачно для %s: ожидалось %t, получено %t", tt.name, tt.expected, result)
			}
		})
	}
}

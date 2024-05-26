package shapes

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name    string
		shape   Shape
		want    float64
		wantErr bool
	}{
		{"Правильный круг", Circle{Radius: 5}, 78.53981633974483, false},
		{"Круг с отрицательным радиусом", Circle{Radius: -5}, 0, true},
		{"Правильный прямоугольник", Rectangle{Width: 10, Height: 5}, 50, false},
		{"Прямоугольник с отрицательной шириной", Rectangle{Width: -10, Height: 5}, 0, true},
		{"Правильный треугольник", Triangle{Base: 8, Height: 6}, 24, false},
		{"Треугольник с отрицательным основанием", Triangle{Base: -8, Height: 6}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateArea(tt.shape)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateArea() ошибка = %v, ожидаемая ошибка %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateArea() = %v, ожидаемое %v", got, tt.want)
			}
		})
	}
}

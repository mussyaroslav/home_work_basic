package chessboard

import "testing"

func TestGenerateBoard(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		want    string
		wantErr bool
	}{
		{"Допустимый размер 3", 3, "# #\n # \n# #\n", false},
		{"Допустимый размер 4", 4, "# # \n # #\n# # \n # #\n", false},
		{"Недопустимый размер 0", 0, "", true},
		{"Недопустимый размер -1", -1, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateBoard(tt.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateBoard() ошибка = %v, ожидается ошибка %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateBoard() = %v, ожидается %v", got, tt.want)
			}
		})
	}
}

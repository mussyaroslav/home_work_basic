package chessboard

import (
	"errors"
	"fmt"
)

func GenerateBoard(size int) (string, error) {
	if size < 1 {
		return "", errors.New("размер шахматной доски должен быть больше нуля")
	}

	board := ""
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board += "#"
			} else {
				board += " "
			}
		}
		board += "\n"
	}
	return board, nil
}

func PrintBoard(size int) error {
	board, err := GenerateBoard(size)
	if err != nil {
		return err
	}
	fmt.Print(board)
	return nil
}

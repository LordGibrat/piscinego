package piscine

import "github.com/01-edu/z01"

// EightQueens находит и выводит все решения задачи о 8 ферзях.
func EightQueens() {
	var board [8]int
	solve(0, &board)
}

// solve - рекурсивная функция для поиска решений с помощью возврата (backtracking).
func solve(col int, board *[8]int) {
	// Если расставили ферзей на все 8 колонок, выводим результат
	if col == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(board[i] + '0'))
		}
		z01.PrintRune('\n')
		return
	}

	// Пробуем поставить ферзя на каждую строку (от 1 до 8)
	for row := 1; row <= 8; row++ {
		if isValid(board, col, row) {
			board[col] = row
			solve(col+1, board) // Переходим к следующей колонке
		}
	}
}

// isValid проверяет, можно ли поставить ферзя на позицию (col, row).
func isValid(board *[8]int, col int, row int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := board[prevCol]

		// Проверка: на той же строке
		if prevRow == row {
			return false
		}

		// Проверка: по диагонали
		diffRow := row - prevRow
		if diffRow < 0 {
			diffRow = -diffRow
		}
		diffCol := col - prevCol

		if diffRow == diffCol {
			return false
		}
	}
	return true
}

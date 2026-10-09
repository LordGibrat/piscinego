package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// Вычисляем общую длину результирующей строки (включая символы \n)
	totalLen := 0
	for _, s := range args {
		totalLen += len(s)
	}
	totalLen += len(args) - 1

	// Создаем буфер нужного размера для эффективности
	result := make([]byte, 0, totalLen)

	for i, s := range args {
		if i > 0 {
			result = append(result, '\n')
		}
		result = append(result, s...)
	}

	return string(result)
}


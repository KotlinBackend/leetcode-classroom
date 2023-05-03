package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	curRow, goingDown := 0, false

	for _, ch := range s {
		rows[curRow] += string(ch)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	res := ""
	for _, row := range rows {
		res += row
	}
	return res
}

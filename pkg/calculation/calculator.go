package calculator

// TODO добавь коммент
func Calc(expression string) (float64, error) {
	// Проверим что выражение не пустое
	if len(expression) == 0 {
		return 0, ErrEmptyExpression
	}
	expression = simplifySpaces(expression)

	simple, err := simplifyParentheses(expression)
	if err != nil {
		return 0, err
	}
	res, err := calculateSimple(simple)
	if err != nil {
		return 0, err
	}

	return res, nil
}

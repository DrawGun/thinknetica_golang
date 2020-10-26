package stub

// Scanner  - заглушка для тестов.
type Scanner struct{}

// Scan - заглушка для тестов.
func (m *Scanner) Scan() (data map[string]string, err error) {
	data = map[string]string{
		"Test url one":   "Test title one",
		"Test url two":   "Test title two",
		"Test url three": "Test title three",
	}

	return data, nil
}

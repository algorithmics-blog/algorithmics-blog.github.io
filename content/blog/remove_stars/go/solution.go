package remove_stars

func removeStars(s string) string {
	res := make([]rune, 0, len(s))

	for _, char := range []rune(s) {
		if char == '*' {
			res = res[:len(res)-1]
		} else {
			res = append(res, char)
		}
	}

	return string(res)
}

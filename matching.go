package main

// match search for regexp anywhere in text
func match(regexp []rune, text []rune) bool {
	if regexp[0] == '^' {
		return matchhere(regexp[1:], text)
	}
	// must look even if string is empty
	if matchhere(regexp, text) {
		return true
	}
	for i := range text {
		if matchhere(regexp, text[i:]) {
			return true
		}
	}
	return false
}

// matchhere search for regexp at beginning of text
func matchhere(regexp []rune, text []rune) bool {
	if len(regexp) == 0 {
		return true
	}
	if len(regexp) > 1 && regexp[1] == '*' {
		return matchstar(regexp[0], regexp[2:], text)
	}
	if regexp[0] == '$' && len(regexp) == 1 {
		return len(text) == 0
	}
	if len(text) > 0 && (regexp[0] == '.' || regexp[0] == text[0]) {
		return matchhere(regexp[1:], text[1:])
	}
	return false
}

// matchstar search for c*regexp at beginning of text
func matchstar(c rune, regexp []rune, text []rune) bool {
	// a * matches zero or more instances
	if matchhere(regexp, text) {
		return true
	}
	for len(text) > 0 && (text[0] == c || c == '.') {
		text = text[1:]
		if matchhere(regexp, text) {
			return true
		}
	}
	return false
}

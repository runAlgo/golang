package greet

import "strings"

// exported functions starts with Capital letter
// other packages can call it -> p1, p2
func Hello(name string) string {
	normalized := normalizeName(name)
	return "Hello " + normalized
}

func normalizeName(name string) string {
	n := strings.TrimSpace(name)
	return strings.ToUpper(n)
}
package string

// Reverse returns the reverse of the passed in string
func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// Reversible is a string that is Reversible
type Reversible string

// Reverse reverse a Reversible
func (s Reversible) Reverse() Reversible {
	return Reversible(Reverse(string(s)))
}

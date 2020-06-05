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

// Reversable is a string that is reversable
type Reversable string

// Reverse reverse a reversable
func (s Reversable) Reverse() Reversable {
	return Reversable(Reverse(string(s)))
}

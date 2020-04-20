package strings

// Reverses a string
/*
	Since strings in Go are immutable, we first convert the string to a mutable array of runes ([]rune), 
	perform the reverse operation on that, and then re-cast to a string.
*/


func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
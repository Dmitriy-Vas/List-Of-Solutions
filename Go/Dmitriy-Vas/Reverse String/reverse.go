package Reverse_String

func Reverse(s string) string {
	o := []rune(s)
	l := len(s)
	for i := 0; i < l/2; i++ {
		w := l - i -1
		o[w], o[i] = o[i], o[w]
	}
	return string(o)
}

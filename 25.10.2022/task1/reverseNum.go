package test

func Reverse(n int) int {
	reverse := 0
	var reminder int
	for n != 0 {
		reminder = n % 10
		reverse = reverse*10 + reminder
		n /= 10
	}
	return reverse
}

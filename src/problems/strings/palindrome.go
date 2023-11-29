package strings

func PalindromeString(str string) bool {
	n := len(str)
	if n <= 0 {
		return false
	}
	start := 0
	end := n - 1
	chars := []rune(str)

	for start < end {
		if chars[start] != chars[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func PalindromeNumber(num int) bool {
	if num < 0 {
		return false
	}
	tmp := num
	sum := 0

	for tmp != 0 {
		sum = sum*10 + tmp%10 // sum will have the reversed value of num
		tmp = tmp / 10
	}
	return sum == num
}

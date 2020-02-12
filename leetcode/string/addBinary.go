package string

import "strconv"

// 67. 二进制求和
func addBinary(a string, b string) (sum string) {

	i := len(a) - 1
	j := len(b) - 1

	var carry uint8

	for i >= 0 || j >= 0 {

		var s uint8

		if i >= 0 && j >= 0 {
			s = (a[i] - '0') + (b[j] - '0') + carry
		} else if i < 0 {
			s = (b[j] - '0') + carry
		} else {
			s = (a[i] - '0') + carry
		}

		this := s & 1
		carry = (s & 2) >> 1

		sum = strconv.Itoa(int(this)) + sum

		i--
		j--
	}

	if carry != 0 {
		sum = "1" + sum
	}

	return
}

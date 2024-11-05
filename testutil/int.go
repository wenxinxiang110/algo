package testutil

// IterDigest 取数值各个位上的单数
func IterDigest(n int) (digest []int) {
	for n != 0 {
		digest = append(digest, n%10)
		n /= 10
	}
	return
}

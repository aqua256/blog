package p0136

// x^x=0，x^0=x，异或满足交换律：成对的数全部抵消，剩下的就是只出现一次的。
func singleNumber(nums []int) int {
	ans := 0
	for _, x := range nums {
		ans ^= x
	}
	return ans
}

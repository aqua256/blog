package p0055

func canJump(nums []int) bool {
	farthest := 0 // 目前能到达的最远下标
	for i, n := range nums {
		if i > farthest {
			return false // 走不到 i
		}
		farthest = max(farthest, i+n)
	}
	return true
}

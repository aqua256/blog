package p0108

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 取中点作根，左右两半大小差 ≤ 1，递归下去自然平衡
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

package p0287

// 把 i -> nums[i] 看成链表的 next 指针：值域 [1,n] 而下标 [0,n]，
// 0 号节点没有入边是链头，重复值 = 有两条入边的节点 = 环的入口。
// 于是套 142 的 Floyd 判圈：快慢指针相遇后，一个回到起点，同速再走，相遇处即入口。
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

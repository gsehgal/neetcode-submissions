func findMin(nums []int) int {
	l:=0
	r:=len(nums)-1
	for {
		mid := (l + r) / 2
		if nums[l] <= nums[r] {
			return nums[l]
		}
		if nums[l] > nums[mid] {
			r=mid
		} else {
			l=mid+1
		}
	}
}

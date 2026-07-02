func threeSum(nums []int) [][]int {
	var res [][]int
	res1 := make(map[[3]int]struct{}) 
	n := len(nums)
	sort.Ints(nums)
	for i:=0;i<n-2;i++ {
		k:=len(nums)-1
		for j:=i+1;j<k; {
			sum := nums[i]+nums[j]+nums[k]
			if sum == 0 {
				key := [3]int{nums[i],nums[j],nums[k]}
				res1[key] = struct{}{}
				j++
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	for key := range res1 {
		res = append(res,key[:])
	}	
	return res
}

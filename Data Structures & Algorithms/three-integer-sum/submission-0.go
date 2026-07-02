func threeSum(nums []int) [][]int {
	var res [][]int
	vals := make(map[int]int)
	n := len(nums)
	for i:=0;i<n;i++ {
		vals[nums[i]] = i
	}
	res1 := make(map[[3]int]struct{})
	for i:=0;i<n-1;i++ {
		for j:=i+1;j<n;j++ {
			sum := nums[i] + nums[j]
			if idx,ok := vals[-sum]; ok {
				if idx!=i && idx!=j {
					key := [3]int{nums[i],nums[j],nums[idx]}
					if key[0] > key[1] { key[0], key[1] = key[1], key[0] }
					if key[1] > key[2] { key[1], key[2] = key[2], key[1] }
					if key[0] > key[1] { key[0], key[1] = key[1], key[0] }
					if _,exists := res1[key]; !exists {
						res1[key] = struct{}{}
					}
				}
			}			
		}
	}
	for key := range res1 {
		res = append(res,key[:])
	}		
	return res
}

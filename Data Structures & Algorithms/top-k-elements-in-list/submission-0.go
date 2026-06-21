func topKFrequent(nums []int, k int) []int {
   freq := make(map[int]int)
   for _, num := range nums{
		val,ok := freq[num]
		if !ok {
			freq[num]=1
		} else {
			freq[num]=val+1
		}
   }
   type Pair struct {
	 num int
	 count int
   }
   var occur []Pair
   for num,count := range freq {
	occur = append(occur,Pair{num,count})
   }
   sort.Slice(occur, func(i, j int) bool {
    return occur[i].count > occur[j].count
	})
	var res []int
	for i:=0;i<k;i++ {
		res = append(res,occur[i].num)
	}
	return res

}

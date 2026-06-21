func groupAnagrams(strs []string) [][]string {
     	anagram := make(map[string][]string)
		for _,str := range strs {
			bytes := []byte(str)
			sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
			key := string(bytes)
			anagram[key] = append(anagram[key],str)
		}
		var res [][]string
		for _,val := range anagram{
			res = append(res,val)
		}
		return res
}

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]struct{})
	if len(s)<2 {
	   return len(s)
	}
	i:=0
	j:=1
	maxLen:=1
	curLen:=1
	seen[s[i]]=struct{}{}
	for j<len(s) {
		if _,exists := seen[s[j]] ; exists {
			for {
				delete(seen,s[i])
				i++
				if _,exists := seen[s[j]] ; !exists {
					break
				} 
			}
			curLen = j-i+1
			fmt.Printf("%v %v %v",curLen,i,j)
		} else {
			fmt.Printf("%v %v %v",curLen,i,j)
			seen[s[j]]=struct{}{}
			curLen = j-i+1
			j++
		}
		maxLen=max(maxLen,curLen)
	}
	return maxLen
	
}

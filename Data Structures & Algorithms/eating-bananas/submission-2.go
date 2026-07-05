func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	l := 1
	r := piles[len(piles)-1]
	minBanana := piles[len(piles)-1]
	for l<r {
		curBanana := (l+r)/2
		totalHours := 0
		for i:=0;i<len(piles);i++ {
			totalHours += piles[i] / curBanana
			if piles[i] % curBanana !=0 {
				totalHours +=1
			}
		}
		if (totalHours<=h){
			minBanana = min(minBanana,curBanana)
			r=curBanana
		} else {
			l=curBanana+1
		}
	}
	return minBanana
}

func dailyTemperatures(temperatures []int) []int {
	temp := []int{temperatures[len(temperatures)-1]}
	index := []int{len(temperatures)-1}
	res := make([]int, len(temperatures))
	res[len(res)-1]=0
	temp = append(temp,temperatures[len(temperatures)-1])
	index = append(index,len(temperatures)-1)
	for i:=len(temperatures)-2;i>=0;i-- {
		if len(temp)>0 && temperatures[i]>=temp[len(temp)-1] {
			for {
				if len(temp)>0 && temperatures[i]>=temp[len(temp)-1] {
					temp = temp[:len(temp)-1]
					index = index[:len(index)-1]
				} else {
					break
				}
			}
			if (len(index)>0) {
				res[i]=index[len(index)-1]-i
			} else {
				res[i]=0
			}
			index = append(index,i)
			temp = append(temp,temperatures[i])		
		} else {
			index = append(index,i)
			temp = append(temp,temperatures[i])		
			res[i]=1
		} 
	}
	return res
}

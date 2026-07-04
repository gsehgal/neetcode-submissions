func carFleet(target int, position []int, speed []int) int {

       pos := make([][2]int,len(position))  
	   for i:=0;i<len(position);i++ {
		pos[i][0] = position[i]
		pos[i][1] = speed[i]
	   }
	   sort.Slice(pos, func(i, j int) bool {
  		  return pos[i][0] > pos[j][0]
		})
		arr := make([]float64,len(position))
		for i:=0;i<len(pos);i++ {
			dist := float64(target - pos[i][0])
			arr[i] =  math.Round((dist / float64(pos[i][1]))*100)/100	
		}

	   var stack []float64
	   stack = append(stack,arr[0])
	   for i:= 1 ; i<len(arr);i++{
		 if (arr[i]>stack[len(stack)-1]) {
			stack = append(stack,arr[i])
		 }
	   }
	   return len(stack)
}

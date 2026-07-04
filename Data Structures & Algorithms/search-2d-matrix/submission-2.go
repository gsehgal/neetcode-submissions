func searchRow(matrix [][]int,target int) int {
	first := 0
	last := len(matrix)
	for first<last {
		mid := (first+last) / 2
		if matrix[mid][0] <= target && matrix[mid][len(matrix[0])-1]>=target {
			return mid
		}
		if matrix[mid][0]>target {
			last = mid;
		}else if matrix[mid][0]<target {
			first = mid+1
		} 
	}
	return -1
}

func searchValue(col []int,target int) bool {
	first := 0
	last := len(col)
	for first<last {
		mid := (first+last)/2
		if col[mid] == target {
			return true
		}
		if col[mid] > target {
			last = mid
		} else {
			first = mid+1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
    row := searchRow(matrix,target)
	if row == -1 {
		return false
	}
	fmt.Printf("Row %v \n",row)
	return searchValue(matrix[row],target)
}

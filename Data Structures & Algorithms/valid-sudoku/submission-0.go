func isValidSudoku(board [][]byte) bool {
	rowDup := make(map[int]map[string]struct{})
    colDup := make(map[int]map[string]struct{})
    
    subBox := make([][]map[string]struct{}, 3)
    for i := range subBox {
        subBox[i] = make([]map[string]struct{}, 3)
        for j := range subBox[i] {
            subBox[i][j] = make(map[string]struct{})
        }
    }

    for i := 0; i < len(board); i++ {
        if _, ok := rowDup[i]; !ok {
            rowDup[i] = make(map[string]struct{})
        }
        
        for j := 0; j < len(board[0]); j++ {
            if _, cok := colDup[j]; !cok {
                colDup[j] = make(map[string]struct{})
            }
            
            val := string(board[i][j]) 
            if val == "." {
                continue
            }
            
            if _, exists := rowDup[i][val]; exists {
                return false
            }
            rowDup[i][val] = struct{}{}
        
            if _, exists := colDup[j][val]; exists {
                return false
            }
            colDup[j][val] = struct{}{}
            
            if _, exists := subBox[i/3][j/3][val]; exists {
                return false
            }
            subBox[i/3][j/3][val] = struct{}{} // Added missing {}
        }
    }
	return true
	}

class NumMatrix:

    def __init__(self, matrix: List[List[int]]):
        self.matrix = matrix
        ROWS,COLS = len(matrix),len(matrix[0])
        self.prefix = [[0]*(COLS+1) for _ in range(ROWS+1)]
        for i in range(ROWS):
            sum=0
            for j in range(COLS):
                sum+=self.matrix[i][j]
                self.prefix[i][j]=sum

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        sum=0    
        for r in range(row1,row2+1):
            if (col1==0):
                sum+=self.prefix[r][col2]
            else:
                sum+=(self.prefix[r][col2]-self.prefix[r][col1-1])
        return sum

# Your NumMatrix object will be instantiated and called as such:
# obj = NumMatrix(matrix)
# param_1 = obj.sumRegion(row1,col1,row2,col2)
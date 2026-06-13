class NumMatrix:

    def __init__(self, matrix: List[List[int]]):
        self.matrix = matrix
        ROW,COL = len(matrix),len(matrix[0])
        self.prefix = matrix[:]
        for i in range(1,COL):
            self.prefix[0][i]=self.prefix[0][i-1]+self.matrix[0][i]
        for i in range(1,ROW):
            self.prefix[i][0]=self.prefix[i-1][0]+self.matrix[i][0]
        for r in range(1,ROW):
            for c in range(1,COL):
                self.prefix[r][c]=self.matrix[r][c] + self.prefix[r-1][c] + self.prefix[r][c-1] - self.prefix[r-1][c-1]

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        sum = self.prefix[row2][col2]
        if (row1>0):
            sum-=self.prefix[row1-1][col2]
        if (col1>0):
            sum-=self.prefix[row2][col1-1]
        if (row1>0 and col1>0):
            sum+=self.prefix[row1-1][col1-1]
        return sum


# Your NumMatrix object will be instantiated and called as such:
# obj = NumMatrix(matrix)
# param_1 = obj.sumRegion(row1,col1,row2,col2)
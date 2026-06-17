class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        trie = {}
        def addToTrie(word:string):
            child=trie
            for c in word:
                if c not in child:
                    child[c]={}
                child = child.get(c)    
            child['*']=True
        for word in words:
            addToTrie(word)
        
        def dfs(node,visited , curWord: List[str],result: List[str],row:int,col:int):
            if not node:
                return
            if board[row][col] not in node:
                return
            curWord.append(board[row][col])    
            child = node.get(board[row][col])
            if '*' in child and child.get('*'):
                result.append(''.join(curWord))
                child['*']=False
            visited[row][col]=True
            neighbours = [(-1,0),(0,1),(1,0),(0,-1)]
            for neighbour in neighbours:
                i , j = neighbour
                nrow = row+i
                ncol = col +j
                if not (nrow<0 or nrow>=len(board) or ncol<0 or ncol>=len(board[0]) or visited[nrow][ncol]):
                    dfs(child, visited,curWord,result,nrow,ncol)
            curWord.pop()
            visited[row][col]=False

        result=[]
        visited = [[False] * len(board[0]) for _ in range(len(board))]
        for row in range(len(board)):
            for col in range(len(board[0])):
                dfs(trie,visited,[],result,row,col)
        return result
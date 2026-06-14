class WordDictionary:

    def __init__(self):
        self.trie={}        

    def addWord(self, word: str) -> None:
        child=self.trie
        for c in word:
            if c not in child:
                child[c]={}
            child=child.get(c)
        child['*']=True

    def helper(self,node,si,word) -> bool:
        for i in range(si,len(word)):
            if word[i]=='.':
                for v in node.keys():
                    if v!='*':
                        if self.helper(node[v],i+1,word):
                            return True
            if word[i] not in node:
                return False
            node=node.get(word[i])
        return node.get('*',False)
        

    def search(self, word: str) -> bool:
        return self.helper(self.trie,0,word)

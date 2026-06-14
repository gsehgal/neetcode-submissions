class PrefixTree:

    def __init__(self):
        self.prefix={}

    def insert(self, word: str) -> None:
        child=self.prefix
        for c in word:
            if c not in child:
                child[c]={}
            child = child.get(c)
        child['*']=True
 

    def search(self, word: str) -> bool:
        child=self.prefix
        for c in word:
            if c not in child:
                return False
            child=child.get(c)
        return '*' in child

    def startsWith(self, prefix: str) -> bool:
        child=self.prefix
        for c in prefix:
            if c not in child:
                return False
            child=child.get(c)
        return True        
        
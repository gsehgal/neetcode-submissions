class Solution:

    def encode(self, strs: List[str]) -> str:
        result=""
        for s in strs:
            l = len(s)
            result+=(str(l)+";"+s)
        return result

    def decode(self, s: str) -> List[str]:
        result=[]
        while len(s)>0:
            i=s.find(';')
            l=int(s[:i])
            result.append(s[i+1:i+1+l])
            s=s[i+1+l:]
        return result



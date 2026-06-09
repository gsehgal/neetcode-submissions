class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        maxLen=0
        L=0
        numOccur={}
        maxf=0
        for R in range(len(s)):
            numOccur[s[R]]=numOccur.get(s[R],0)+1
            maxf = max(maxf,numOccur.get(s[R]))
            replacement = (R-L+1)-maxf
            if (replacement>k):
                numOccur[s[L]]-=1
                L+=1
            maxLen=max(maxLen,R-L+1)
        return maxLen
            

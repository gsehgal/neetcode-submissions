class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        window = set()
        L=0
        maxLen=0
        curLen=0
        for R in range(len(s)):
            if s[R] in window:
                while (s[L] != s[R]):
                    window.remove(s[L])
                    L+=1
                L+=1
                curLen=R-L+1
            else:
               curLen+=1
               window.add(s[R])
            maxLen=max(maxLen,curLen)
        return maxLen         
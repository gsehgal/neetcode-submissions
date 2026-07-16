class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        s1dict = {}
        s2dict = {}
        for i in range(len(s1)):
            c = s1dict.get(s1[i],0)
            s1dict[s1[i]]=c+1
        l=0
        r=0
        while r<len(s2):
            if s2[r] in s1dict:
                c=s2dict.get(s2[r],0)
                s2dict[s2[r]]=c+1
                while s2dict[s2[r]] > s1dict[s2[r]]:
                     print(s2[l] + "\n") 
                     print("l  "+str(l)+"\n")
                     c=s2dict.get(s2[l])
                     s2dict[s2[l]] = c-1
                     l+=1
                if (r-l+1)==len(s1):
                    return True
            else:
                s2dict={}
                while r<len(s2) and s2[r] not in s1dict:
                    r+=1
                if r<len(s2):
                    s2dict[s2[r]] = 1
                l=r
            r+=1
        return False

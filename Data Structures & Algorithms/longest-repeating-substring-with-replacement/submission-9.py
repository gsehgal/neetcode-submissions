class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        hash = {}
        maxCount = 0
        l=0
        r=0
        maxLen=0    
        while r<len(s):
            count = hash.get(s[r],0)
            hash[s[r]] = count+1
            maxCount = max(maxCount,count+1)
            length = r-l+1
            if (length-maxCount <=k):
                maxLen = max(maxLen,length)
            else:
                count = hash.get(s[l])
                hash[s[l]]=count-1
                for key, value in hash.items():
                        maxCount = max(value,maxCount)
                
                l+=1
            r+=1

        return maxLen
                            
 


            
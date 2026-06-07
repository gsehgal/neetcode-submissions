class Solution:
    def numOfSubarrays(self, arr: List[int], k: int, threshold: int) -> int:
        result = 0
        curSum=0
        L=0

        for R in range(len(arr)):
            curSum+=arr[R]
            if (R-L+1==k):
                if (curSum/k >= threshold):
                    result+=1
                curSum-=arr[L]
                L+=1
        return result
                

        
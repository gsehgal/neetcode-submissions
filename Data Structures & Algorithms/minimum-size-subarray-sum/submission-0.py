class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        L=0
        curSum=0
        minLen=2*len(nums)
        for R in range(len(nums)):
            curSum+=nums[R]
            while (curSum>=target):
                minLen=min(R-L+1,minLen)
                curSum-=nums[L]
                L+=1
        if (minLen==2*len(nums)):
            return 0
        return minLen
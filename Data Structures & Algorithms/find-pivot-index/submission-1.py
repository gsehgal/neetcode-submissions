class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        pSum = [0] * len(nums)
        for i in range(len(nums)):
            pSum[i]=pSum[i-1]+nums[i]
        totalSum=pSum[len(pSum)-1]
        if (totalSum-nums[0]==0):
            return 0
        for i in range(1,len(nums)):
            rsum = totalSum - pSum[i-1]-nums[i]
            print("i = "+str(i) + " pSum = "+str(pSum[i-1])+ " num= "+str(nums[i]))
            if (rsum==pSum[i-1]):
                return i
        return -1 
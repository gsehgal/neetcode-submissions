class Solution:
    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        maxSum = nums[0]
        minSum = nums[0]
        curMaxSum = 0
        curMinSum = 0
        totalSum=0
        for n in nums:
            curMaxSum = max(curMaxSum,0)
            curMinSum = min(curMinSum,0)
            curMaxSum+=n
            curMinSum+=n
            totalSum+=n
            maxSum=max(maxSum,curMaxSum)
            minSum=min(minSum,curMinSum)
        if (maxSum<0):
            return maxSum
        return max(maxSum,totalSum-minSum)
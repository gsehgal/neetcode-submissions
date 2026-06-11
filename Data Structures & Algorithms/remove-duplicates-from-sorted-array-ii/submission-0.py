class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        L=0
        occur=0
        for R in range(1,len(nums)):
            if (nums[R]!=nums[R-1]):
                L+=1
                nums[L]=nums[R]
                occur=0
            elif (nums[R]==nums[R-1] and  occur==0):
                L+=1
                nums[L]=nums[R]
                occur+=1
            else:
                occur+=1
        return L+1
        
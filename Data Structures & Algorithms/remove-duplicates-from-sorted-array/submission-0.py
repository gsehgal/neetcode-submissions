class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if (len(nums)==1):
            return 1
        L=0
        for R in range(1,len(nums)):
            if (nums[R]!=nums[R-1]):
                L+=1
                nums[L]=nums[R]
        return L+1
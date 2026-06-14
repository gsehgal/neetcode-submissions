class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        zero = False
        product = 1
        output=[0]*len(nums)
        for num in nums:
            if (num==0):
                if (zero):
                    return output
                zero=True
            else:
                product*=num
        for i in range(len(nums)):
            if not zero:
                output[i] = product//nums[i]
            elif nums[i] == 0:
                output[i] = product
                break
        return output
            
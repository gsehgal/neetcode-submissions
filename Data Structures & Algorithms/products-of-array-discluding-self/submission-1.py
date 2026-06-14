class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix = [0]*len(nums)
        postfix = [0]*len(nums)
        output = [0]*len(nums)
        product=1
        prefix[0]=1
        zero = (nums[0] == 0)
        for i in range(1,len(nums)):
            if nums[i]==0:
                if zero:
                    return output
                zero=True
            prefix[i]=prefix[i-1]*nums[i-1]
        postfix[len(nums)-1]=1
        for i in range(len(nums)-2,-1,-1):
            postfix[i]=postfix[i+1]*nums[i+1]
        for i in range(len(nums)):
            output[i]=prefix[i]*postfix[i]
        return output
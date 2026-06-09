class Solution:
    def maxTurbulenceSize(self, arr: List[int]) -> int:
        if (len(arr)==1):
            return 1
        if (len(arr)==2 and arr[0]!=arr[1]):
            return 2
        maxLen=0
        curLen=0
        L=0
        def Sign(a,b)->int:
            if (a<b):
                return -1
            elif (a>b):
                return 1
            else:
                return 0
        psign=Sign(arr[1],arr[0])
        for R in range(1,len(arr)):
            sign = Sign(arr[R],arr[R-1])
            if (sign==0):
                L=R
            elif (sign==psign):
                  L=R-1
            else:
                psign=sign         
            maxLen=max(maxLen,R-L+1)
        return maxLen
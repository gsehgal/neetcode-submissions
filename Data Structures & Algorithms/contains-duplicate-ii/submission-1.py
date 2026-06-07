class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        duplicate  = {}
        for i in range(len(nums)):
            if nums[i] in duplicate:
                if (abs(i - duplicate[nums[i]])<=k):
                    return True
            duplicate[nums[i]]=i
        return False

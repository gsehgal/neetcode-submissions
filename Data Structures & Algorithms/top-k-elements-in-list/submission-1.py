class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq=defaultdict(int)
        for num in nums:
            freq[num]+=1
        arr = []
        for key in freq:
            arr.append((-1*freq[key],key))

        heapq.heapify(arr)
        output=[]
        for i in range(1,k+1):
            _,val  = heapq.heappop(arr)
            output.append(val)
        return output

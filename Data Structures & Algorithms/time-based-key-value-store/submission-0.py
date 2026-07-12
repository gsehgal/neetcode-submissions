class TimeMap:

    def __init__(self):
        self.timemap = defaultdict(list)

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.timemap[key].append((timestamp,value))

    def get(self, key: str, timestamp: int) -> str:
        if key not in self.timemap:
            return ""
        values = self.timemap[key]
        l = 0
        r = len(values)-1
        res=""
        while l <= r :
            mid = (l+r)//2
            if timestamp == values[mid][0]:
                return values[mid][1]
            if timestamp < values[mid][0]:
                r = mid-1
            else:
               res = values[mid][1]
               l=mid+1
        return res
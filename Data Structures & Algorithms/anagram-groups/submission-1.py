class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagrams = {}
        for word in strs:
            key = "".join(sorted(word))
            arr = anagrams.get(key,[])
            arr.append(word)
            anagrams[key] = arr

        output=[]
        for key in anagrams:
            output.append(anagrams[key])

        return output
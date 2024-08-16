class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        
        anahrams = {}

        for str in strs:
            sorted_str = self.sort_string(str)
            if sorted_str in anahrams:
                anahrams[sorted_str].append(str)
            else:
                anahrams[sorted_str] = [str]
        
        return list(anahrams.values())

    
    def sort_string(self, s):
        return ''.join(sorted(s))
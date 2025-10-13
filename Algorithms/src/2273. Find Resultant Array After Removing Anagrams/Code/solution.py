from typing import List

class Solution:
    def removeAnagrams(self, words: List[str]) -> List[str]:
        result = []
        prev_sig = ""
        for w in words:
            # signature by sorted characters
            sig = ''.join(sorted(w))
            if sig != prev_sig:
                result.append(w)
                prev_sig = sig
        return result

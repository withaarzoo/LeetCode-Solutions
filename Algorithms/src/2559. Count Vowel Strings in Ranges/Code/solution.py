from typing import List

class Solution:
    def vowelStrings(self, words: List[str], queries: List[List[int]]) -> List[int]:
        vowels = {'a', 'e', 'i', 'o', 'u'}
        n = len(words)
        prefix = [0] * n

        # Precompute the prefix sum
        for i in range(n):
            if words[i][0] in vowels and words[i][-1] in vowels:
                prefix[i] = 1
            if i > 0:
                prefix[i] += prefix[i - 1]

        # Answer the queries
        result = []
        for l, r in queries:
            result.append(prefix[r] - (prefix[l - 1] if l > 0 else 0))
        return result
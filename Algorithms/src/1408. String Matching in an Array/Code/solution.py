class Solution:
    def stringMatching(self, words: List[str]) -> List[str]:
        # Sort words by length
        words.sort(key=len)

        result = []

        # Check for substrings
        for i in range(len(words)):
            for j in range(i + 1, len(words)):
                if words[i] in words[j]:
                    result.append(words[i])
                    break

        return result

class Solution:
    def prefixCount(self, words: List[str], pref: str) -> int:
        count = 0
        for word in words:
            # Check if the word starts with the prefix
            if word.startswith(pref):
                count += 1
        return count

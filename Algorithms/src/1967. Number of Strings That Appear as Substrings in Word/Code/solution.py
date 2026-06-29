class Solution:
    def numOfStrings(self, patterns: List[str], word: str) -> int:

        # Store the number of matching patterns
        count = 0

        # Check every pattern
        for pattern in patterns:

            # If pattern exists inside word, increase the answer
            if pattern in word:
                count += 1

        # Return the total number of matches
        return count
class Solution:
    def minimumPushes(self, word: str) -> int:

        # Store the final answer
        pushes = 0

        # Visit every character
        for i in range(len(word)):

            # Every group of 8 letters has the same typing cost
            pushes += (i // 8) + 1

        # Return the minimum number of pushes
        return pushes
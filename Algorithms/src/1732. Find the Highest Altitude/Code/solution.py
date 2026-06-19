class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        # Current altitude starts at 0
        currentAltitude = 0

        # Highest altitude seen so far
        maxAltitude = 0

        # Process every altitude change
        for change in gain:
            # Apply the altitude change
            currentAltitude += change

            # Update highest altitude if needed
            maxAltitude = max(maxAltitude, currentAltitude)

        # Return the answer
        return maxAltitude
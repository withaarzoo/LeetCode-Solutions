class Solution {
    public int largestAltitude(int[] gain) {
        // Current altitude starts at 0
        int currentAltitude = 0;

        // Highest altitude seen so far
        int maxAltitude = 0;

        // Process every gain value
        for (int change : gain) {
            // Apply altitude change
            currentAltitude += change;

            // Update highest altitude if needed
            maxAltitude = Math.max(maxAltitude, currentAltitude);
        }

        // Return the answer
        return maxAltitude;
    }
}
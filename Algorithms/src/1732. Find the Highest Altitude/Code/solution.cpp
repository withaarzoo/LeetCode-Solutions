class Solution {
public:
    int largestAltitude(vector<int>& gain) {
        // Current altitude starts at 0
        int currentAltitude = 0;
        
        // Highest altitude seen so far
        int maxAltitude = 0;

        // Process every gain value
        for (int change : gain) {
            // Move to the next point by applying altitude change
            currentAltitude += change;

            // Update highest altitude if current altitude is greater
            maxAltitude = max(maxAltitude, currentAltitude);
        }

        // Return the highest altitude reached
        return maxAltitude;
    }
};
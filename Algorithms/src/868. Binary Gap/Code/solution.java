class Solution {
    public int binaryGap(int n) {
        int lastPosition = -1; // last index where 1 was found
        int maxDistance = 0; // maximum distance
        int currentPosition = 0; // bit index

        while (n > 0) {
            // Check if last bit is 1
            if ((n & 1) == 1) {
                // If previous 1 exists
                if (lastPosition != -1) {
                    maxDistance = Math.max(maxDistance, currentPosition - lastPosition);
                }
                lastPosition = currentPosition;
            }

            n >>= 1; // shift right
            currentPosition++;
        }

        return maxDistance;
    }
}
class Solution {
public:
    int minimumPushes(string word) {
        // Store the final minimum number of pushes
        int pushes = 0;

        // Visit every distinct letter
        for (int i = 0; i < word.size(); i++) {

            // Every group of 8 letters increases the push count by 1
            // i = 0..7   -> cost = 1
            // i = 8..15  -> cost = 2
            // i = 16..23 -> cost = 3
            // i = 24..25 -> cost = 4
            pushes += (i / 8) + 1;
        }

        // Return the minimum total pushes
        return pushes;
    }
};
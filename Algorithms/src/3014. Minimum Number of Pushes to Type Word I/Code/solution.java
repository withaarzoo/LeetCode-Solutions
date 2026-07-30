class Solution {
    public int minimumPushes(String word) {

        // Store the final answer
        int pushes = 0;

        // Traverse every character
        for (int i = 0; i < word.length(); i++) {

            // Every block of 8 letters has the same cost
            pushes += (i / 8) + 1;
        }

        // Return the minimum number of pushes
        return pushes;
    }
}
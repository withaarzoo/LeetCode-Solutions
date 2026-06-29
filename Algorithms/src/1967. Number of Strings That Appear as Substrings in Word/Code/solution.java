class Solution {
    public int numOfStrings(String[] patterns, String word) {

        // Store the number of matching patterns
        int count = 0;

        // Check every pattern
        for (String pattern : patterns) {

            // indexOf() returns -1 if the substring does not exist
            if (word.indexOf(pattern) != -1) {
                count++;
            }
        }

        // Return the final answer
        return count;
    }
}
class Solution {
    public int countCommas(int n) {
        long answer = 0; // Stores the total number of commas.
        long threshold = 1000; // First number that contains a comma.

        while (threshold <= n) {
            // Every number from threshold through n contributes
            // one comma for this comma group.
            answer += n - threshold + 1;

            // Move to the next comma level.
            threshold *= 1000;
        }

        // The answer fits in an int for the given constraints.
        return (int) answer;
    }
}
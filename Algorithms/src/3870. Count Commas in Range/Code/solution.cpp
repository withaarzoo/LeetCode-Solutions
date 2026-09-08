class Solution {
public:
    int countCommas(int n) {
        long long answer = 0; // Stores the total number of commas.
        long long threshold = 1000; // First number that contains a comma.

        while (threshold <= n) {
            // Every number from threshold through n has one comma
            // corresponding to this comma group.
            answer += n - threshold + 1;

            // Move to the next comma level: 1,000 -> 1,000,000 -> 1,000,000,000.
            threshold *= 1000;
        }

        // The problem's answer fits in an integer for the given constraints.
        return static_cast<int>(answer);
    }
};
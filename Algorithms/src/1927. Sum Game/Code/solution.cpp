class Solution {
public:
    bool sumGame(string num) {
        // Find the middle because the string always has even length.
        int mid = num.size() / 2;

        // Store the sums of known digits in both halves.
        int leftSum = 0, rightSum = 0;

        // Count how many '?' characters exist in both halves.
        int leftQuestion = 0, rightQuestion = 0;

        // Scan every character once and update the correct half.
        for (int i = 0; i < num.size(); i++) {
            if (i < mid) {
                // This character belongs to the left half.
                if (num[i] == '?')
                    leftQuestion++;
                else
                    leftSum += num[i] - '0';
            } else {
                // This character belongs to the right half.
                if (num[i] == '?')
                    rightQuestion++;
                else
                    rightSum += num[i] - '0';
            }
        }

        // Bob can force equality only when this exact condition is true.
        // If equality is impossible, Alice can force the sums to be different.
        return 2 * (leftSum - rightSum) !=
               9 * (rightQuestion - leftQuestion);
    }
};
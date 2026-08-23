class Solution {
    public boolean sumGame(String num) {
        // Find the starting index of the right half.
        int mid = num.length() / 2;

        // Store the sums of known digits in both halves.
        int leftSum = 0;
        int rightSum = 0;

        // Count '?' characters in both halves.
        int leftQuestion = 0;
        int rightQuestion = 0;

        // Visit every character and update its corresponding half.
        for (int i = 0; i < num.length(); i++) {
            if (i < mid) {
                // This character is in the left half.
                if (num.charAt(i) == '?') {
                    leftQuestion++;
                } else {
                    leftSum += num.charAt(i) - '0';
                }
            } else {
                // This character is in the right half.
                if (num.charAt(i) == '?') {
                    rightQuestion++;
                } else {
                    rightSum += num.charAt(i) - '0';
                }
            }
        }

        // Bob wins only when he can make both final sums exactly equal.
        // Otherwise, Alice can force the sums to stay different.
        return 2 * (leftSum - rightSum) !=
               9 * (rightQuestion - leftQuestion);
    }
}
class Solution {
    public int countOdds(int low, int high) {
        // Helper function: count of odd numbers from 1 to x
        // Using a private method for clarity
        return oddsUpTo(high) - oddsUpTo(low - 1);
    }

    private int oddsUpTo(int x) {
        // Integer division in Java automatically floors the result
        return (x + 1) / 2;
    }
}

class Solution {
    public int findTheWinner(int n, int k) {
        // Calls the recursive Josephus function and adjusts the result from 0-indexed
        // to 1-indexed
        return josephus(n, k) + 1;
    }

    private int josephus(int n, int k) {
        // Base case: when there's only one person left, they are the winner
        if (n == 1) {
            return 0;
        }
        // Recursive case: calculate the position of the winner for the current state
        // josephus(n - 1, k) finds the winner position in the reduced group of (n-1)
        // people
        // (josephus(n - 1, k) + k) % n adjusts the position for the current group size
        // n
        return (josephus(n - 1, k) + k) % n;
    }
}

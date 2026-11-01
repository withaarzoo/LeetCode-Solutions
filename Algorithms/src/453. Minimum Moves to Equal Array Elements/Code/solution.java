class Solution {
    public int minMoves(int[] nums) {
        int mn = Integer.MAX_VALUE;
        for (int x : nums) if (x < mn) mn = x;
        long moves = 0;
        for (int x : nums) moves += x - mn;
        return (int) moves;
    }
}
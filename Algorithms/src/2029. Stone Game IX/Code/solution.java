class Solution {
    public boolean stoneGameIX(int[] stones) {
        // cnt[r] stores how many stones have remainder r when divided by 3.
        int[] cnt = new int[3];

        // Count the stones in each remainder group.
        for (int stone : stones) {
            // Only the remainder affects the game.
            cnt[stone % 3]++;
        }

        // With an even number of remainder-0 stones,
        // Alice needs both a remainder-1 and a remainder-2 stone.
        if (cnt[0] % 2 == 0) {
            return cnt[1] > 0 && cnt[2] > 0;
        }

        // With an odd number of remainder-0 stones,
        // Alice wins when the two useful groups differ by more than 2.
        return Math.abs(cnt[1] - cnt[2]) > 2;
    }
}
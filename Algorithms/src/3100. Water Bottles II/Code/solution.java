class Solution {
    public int maxBottlesDrunk(int numBottles, int numExchange) {
        int full = numBottles;    // current full bottles
        int empty = 0;            // current empty bottles
        int ans = 0;              // total bottles drunk
        int curEx = numExchange;  // current exchange threshold

        while (full > 0) {
            // drink all full bottles
            ans += full;
            empty += full;
            full = 0;

            // perform one-by-one exchanges while possible
            while (empty >= curEx) {
                empty -= curEx;  // use curEx empty bottles
                full += 1;       // receive one full bottle
                curEx += 1;      // increase exchange requirement
            }
        }
        return ans;
    }
}

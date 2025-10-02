class Solution {
public:
    int maxBottlesDrunk(int numBottles, int numExchange) {
        int full = numBottles;    // current full bottles
        int empty = 0;            // current empty bottles
        int ans = 0;              // total bottles drunk
        int curEx = numExchange;  // current exchange threshold (increases after each exchange)

        // keep going while I have any full bottles to drink
        while (full > 0) {
            // drink everything I have now
            ans += full;
            empty += full;
            full = 0;

            // exchange empties for full bottles one-by-one
            while (empty >= curEx) {
                empty -= curEx;  // spend curEx empties
                full += 1;       // get one full bottle
                curEx += 1;      // increase threshold after each exchange
            }
        }
        return ans;
    }
};

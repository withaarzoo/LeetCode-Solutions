class Solution {
public:
    int numWaterBottles(int numBottles, int numExchange) {
        // total bottles I can drink so far
        int total = numBottles;
        // empties I have after drinking the initial bottles
        int empties = numBottles;

        // keep exchanging empties for new full bottles while possible
        while (empties >= numExchange) {
            int newFull = empties / numExchange;           // how many new full bottles I get
            total += newFull;                             // drink them (count them)
            empties = newFull + (empties % numExchange);  // update empties: leftover + newly emptied bottles
        }

        return total;
    }
};

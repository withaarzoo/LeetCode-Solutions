class Solution {
    public int numWaterBottles(int numBottles, int numExchange) {
        int total = numBottles;
        int empties = numBottles;

        while (empties >= numExchange) {
            int newFull = empties / numExchange;
            total += newFull;
            empties = newFull + (empties % numExchange);
        }

        return total;
    }
}

class Solution {
    public int maxIceCream(int[] costs, int coins) {

        // Maximum possible cost according to constraints
        final int MAX_COST = 100000;

        // Frequency array to count ice cream bars of each cost
        int[] freq = new int[MAX_COST + 1];

        // Count occurrences of every cost
        for (int cost : costs) {
            freq[cost]++;
        }

        // Stores total ice cream bars purchased
        int answer = 0;

        // Process costs from smallest to largest
        for (int cost = 1; cost <= MAX_COST; cost++) {

            // Skip unavailable costs
            if (freq[cost] == 0) {
                continue;
            }

            // Maximum bars affordable at this cost
            int canBuy = Math.min(freq[cost], coins / cost);

            // Increase purchased count
            answer += canBuy;

            // Deduct spent coins
            coins -= canBuy * cost;
        }

        return answer;
    }
}
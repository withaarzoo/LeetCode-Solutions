class Solution {
public:
    int minMoves(vector<int>& nums, int limit) {
        
        int n = nums.size();

        // Difference array to track move changes
        vector<int> diff(2 * limit + 2, 0);

        // Process every pair
        for (int i = 0; i < n / 2; i++) {

            int a = min(nums[i], nums[n - 1 - i]);
            int b = max(nums[i], nums[n - 1 - i]);

            // By default every sum needs 2 moves
            // We subtract moves where improvement is possible

            // For sums in [a+1, b+limit], only 1 move is needed
            diff[a + 1] -= 1;
            diff[b + limit + 1] += 1;

            // For exact sum a+b, 0 moves are needed
            diff[a + b] -= 1;
            diff[a + b + 1] += 1;
        }

        int pairs = n / 2;

        // Initially every pair contributes 2 moves
        int current = pairs * 2;

        int answer = INT_MAX;

        // Build prefix sums to compute moves for every target sum
        for (int sum = 2; sum <= 2 * limit; sum++) {

            current += diff[sum];

            answer = min(answer, current);
        }

        return answer;
    }
};
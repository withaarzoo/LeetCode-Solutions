class Solution {
public:
    int minimumPairRemoval(vector<int>& nums) {
        int operations = 0;

        // Function to check if array is non-decreasing
        auto isSorted = [&]() {
            for (int i = 1; i < nums.size(); i++) {
                if (nums[i] < nums[i - 1]) return false;
            }
            return true;
        };

        // Keep performing operations until array becomes non-decreasing
        while (!isSorted()) {
            int minSum = INT_MAX;
            int index = 0;

            // Find adjacent pair with minimum sum
            for (int i = 0; i + 1 < nums.size(); i++) {
                int sum = nums[i] + nums[i + 1];
                if (sum < minSum) {
                    minSum = sum;
                    index = i;
                }
            }

            // Replace the pair with their sum
            nums[index] = minSum;
            nums.erase(nums.begin() + index + 1);

            operations++;
        }

        return operations;
    }
};

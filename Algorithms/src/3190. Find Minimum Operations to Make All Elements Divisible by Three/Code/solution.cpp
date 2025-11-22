class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        int operations = 0;
        
        // Traverse each number in the array
        for (int x : nums) {
            // If x is not divisible by 3, we need exactly 1 operation
            if (x % 3 != 0) {
                operations++;
            }
        }
        
        return operations;
    }
};

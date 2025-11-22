class Solution {
    public int minimumOperations(int[] nums) {
        int operations = 0;
        
        // Go through each element in the array
        for (int x : nums) {
            // If x % 3 is not 0, one operation is needed
            if (x % 3 != 0) {
                operations++;
            }
        }
        
        return operations;
    }
}

class Solution {
    public int[] resultArray(int[] nums) {
        // I create arr1 and arr2 to simulate the two arrays from the problem.
        int[] arr1 = new int[nums.length];
        int[] arr2 = new int[nums.length];

        // I keep separate sizes because normal Java arrays do not grow automatically.
        int size1 = 1;
        int size2 = 1;

        // The first number must go into arr1.
        arr1[0] = nums[0];

        // The second number must go into arr2.
        arr2[0] = nums[1];

        // I process every remaining number starting from index 2.
        for (int i = 2; i < nums.length; i++) {
            // I compare the current last elements of arr1 and arr2.
            if (arr1[size1 - 1] > arr2[size2 - 1]) {
                // I add the number to arr1 and increase its size.
                arr1[size1++] = nums[i];
            } else {
                // Otherwise, I add the number to arr2 and increase its size.
                arr2[size2++] = nums[i];
            }
        }

        // I create the final array because LeetCode expects an int[] result.
        int[] result = new int[nums.length];

        // I copy all used elements of arr1 first.
        System.arraycopy(arr1, 0, result, 0, size1);

        // I copy all used elements of arr2 immediately after arr1.
        System.arraycopy(arr2, 0, result, size1, size2);

        // I return the concatenated result.
        return result;
    }
}
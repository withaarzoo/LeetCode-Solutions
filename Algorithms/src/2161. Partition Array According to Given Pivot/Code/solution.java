class Solution {
    public int[] pivotArray(int[] nums, int pivot) {

        // Lists to store three groups
        List<Integer> smaller = new ArrayList<>();
        List<Integer> equal = new ArrayList<>();
        List<Integer> greater = new ArrayList<>();

        // Classify each element
        for (int num : nums) {
            if (num < pivot) {
                smaller.add(num);
            } else if (num == pivot) {
                equal.add(num);
            } else {
                greater.add(num);
            }
        }

        // Result array of same size
        int[] result = new int[nums.length];
        int index = 0;

        // Add smaller elements
        for (int num : smaller) {
            result[index++] = num;
        }

        // Add equal elements
        for (int num : equal) {
            result[index++] = num;
        }

        // Add greater elements
        for (int num : greater) {
            result[index++] = num;
        }

        // Return final answer
        return result;
    }
}
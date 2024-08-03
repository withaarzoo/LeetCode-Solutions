import java.util.Arrays;

class Solution {
    /**
     * This method checks if two arrays can be made equal by sorting them.
     * 
     * @param target The target array we want to compare to.
     * @param arr    The array that we want to manipulate and compare to the target
     *               array.
     * @return true if both arrays can be made equal by sorting, otherwise false.
     */
    public boolean canBeEqual(int[] target, int[] arr) {
        // Step 1: Sort the target array
        Arrays.sort(target);

        // Step 2: Sort the arr array
        Arrays.sort(arr);

        // Step 3: Compare the sorted target and arr arrays
        // If they are equal, it means the original arrays can be rearranged to match
        // each other
        return Arrays.equals(target, arr);
    }
}

class Solution {
    public int maximumElementAfterDecrementingAndRearranging(int[] arr) {

        // Sort the array in ascending order.
        Arrays.sort(arr);

        // The first element must be 1.
        arr[0] = 1;

        // Make every element as large as possible.
        for (int i = 1; i < arr.length; i++) {

            // Limit the current value to previous + 1.
            arr[i] = Math.min(arr[i], arr[i - 1] + 1);
        }

        // The last element is the answer.
        return arr[arr.length - 1];
    }
}
import java.util.*;

class Solution {
    public int[] sortByBits(int[] arr) {

        // Convert to Integer array for custom sorting
        Integer[] nums = Arrays.stream(arr).boxed().toArray(Integer[]::new);

        Arrays.sort(nums, (a, b) -> {

            int bitsA = Integer.bitCount(a);
            int bitsB = Integer.bitCount(b);

            if (bitsA != bitsB)
                return bitsA - bitsB;

            return a - b;
        });

        // Convert back to int[]
        for (int i = 0; i < arr.length; i++) {
            arr[i] = nums[i];
        }

        return arr;
    }
}
class Solution {
    public int missingInteger(int[] nums) {
        // I start the sum with the first element because
        // nums[0] always belongs to the sequential prefix.
        int sum = nums[0];

        // I check each following element to find the longest sequential prefix.
        for (int i = 1; i < nums.length; i++) {
            // The sequence continues only when the current value
            // is exactly one greater than the previous value.
            if (nums[i] == nums[i - 1] + 1) {
                sum += nums[i];
            } else {
                // The sequence breaks, so I stop checking the prefix.
                break;
            }
        }

        // I put all array values into a HashSet for O(1) average lookup.
        java.util.HashSet<Integer> seen = new java.util.HashSet<>();

        // I add every value from the array to the set.
        for (int num : nums) {
            seen.add(num);
        }

        // I start looking for the missing number from the prefix sum.
        int answer = sum;

        // If the current candidate exists, I move to the next integer.
        while (seen.contains(answer)) {
            answer++;
        }

        // The first value not present in the array is the answer.
        return answer;
    }
}
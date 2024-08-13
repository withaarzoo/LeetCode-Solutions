import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution {
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        // Step 1: Sort the array to manage duplicates and make it easier to handle the
        // combination logic.
        Arrays.sort(candidates);

        // Step 2: Initialize the result list which will store all valid combinations.
        List<List<Integer>> result = new ArrayList<>();

        // Step 3: Initialize a temporary list to store the current combination.
        List<Integer> current = new ArrayList<>();

        // Step 4: Start the backtracking process from the 0th index of candidates.
        backtrack(candidates, target, 0, current, result);

        // Step 5: Return the result containing all valid combinations.
        return result;
    }

    // Helper method to perform backtracking.
    private void backtrack(int[] candidates, int target, int start, List<Integer> current, List<List<Integer>> result) {
        // Step 6: Base case - if the target becomes 0, it means we have found a valid
        // combination.
        if (target == 0) {
            // Add a copy of the current combination to the result.
            result.add(new ArrayList<>(current));
            return;
        }

        // Step 7: Iterate through the candidates starting from the 'start' index.
        for (int i = start; i < candidates.length; i++) {
            // Step 8: Skip duplicates - if the current candidate is the same as the
            // previous one and it's not the first element in the loop.
            if (i > start && candidates[i] == candidates[i - 1]) {
                continue; // Move to the next candidate.
            }

            // Step 9: If the current candidate exceeds the remaining target, no need to
            // proceed further as all further elements will also exceed.
            if (candidates[i] > target) {
                break; // Exit the loop as further candidates are too large.
            }

            // Step 10: Include the current candidate in the current combination.
            current.add(candidates[i]);

            // Step 11: Recursively call the backtrack method with the remaining target and
            // the next starting index.
            backtrack(candidates, target - candidates[i], i + 1, current, result);

            // Step 12: Backtrack - remove the last candidate from the current combination
            // to explore other possibilities.
            current.remove(current.size() - 1);
        }
    }
}

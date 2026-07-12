class Solution {
    public int[] arrayRankTransform(int[] arr) {
        // Create a copy of the original array
        int[] sorted = arr.clone();

        // Sort the copied array
        Arrays.sort(sorted);

        // Store each unique value and its rank
        HashMap<Integer, Integer> rank = new HashMap<>();
        int currentRank = 1;

        // Assign ranks to unique values only
        for (int num : sorted) {
            if (!rank.containsKey(num)) {
                rank.put(num, currentRank++);
            }
        }

        // Replace every element with its rank
        for (int i = 0; i < arr.length; i++) {
            arr[i] = rank.get(arr[i]);
        }

        // Return the transformed array
        return arr;
    }
}
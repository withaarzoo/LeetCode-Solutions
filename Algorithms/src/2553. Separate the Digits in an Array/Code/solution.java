class Solution {
    public int[] separateDigits(int[] nums) {

        // Dynamic list to store digits temporarily
        List<Integer> list = new ArrayList<>();

        // Traverse every number
        for (int num : nums) {

            // Convert number to string
            String s = String.valueOf(num);

            // Traverse each character of the string
            for (char ch : s.toCharArray()) {

                // Convert character into integer digit
                list.add(ch - '0');
            }
        }

        // Convert List<Integer> into int[]
        int[] result = new int[list.size()];

        for (int i = 0; i < list.size(); i++) {
            result[i] = list.get(i);
        }

        // Return final answer
        return result;
    }
}
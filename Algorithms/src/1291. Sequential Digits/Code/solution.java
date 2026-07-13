class Solution {
    public List<Integer> sequentialDigits(int low, int high) {

        // String containing all consecutive digits
        String digits = "123456789";

        // Result list
        List<Integer> ans = new ArrayList<>();

        // Number of digits in low and high
        int minLen = String.valueOf(low).length();
        int maxLen = String.valueOf(high).length();

        // Try every possible length
        for (int len = minLen; len <= maxLen; len++) {

            // Generate every substring of current length
            for (int start = 0; start + len <= 9; start++) {

                // Convert substring into an integer
                int num = Integer.parseInt(digits.substring(start, start + len));

                // Keep only numbers inside the range
                if (num >= low && num <= high) {
                    ans.add(num);
                }
            }
        }

        return ans;
    }
}
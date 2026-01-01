class Solution {
    public int[] plusOne(int[] digits) {
        // Traverse from last digit
        for (int i = digits.length - 1; i >= 0; i--) {
            digits[i]++;

            if (digits[i] < 10) { // No carry
                return digits;
            }

            digits[i] = 0; // Carry continues
        }

        // All digits were 9
        int[] result = new int[digits.length + 1];
        result[0] = 1;
        return result;
    }
}

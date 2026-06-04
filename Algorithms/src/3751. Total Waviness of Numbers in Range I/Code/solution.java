class Solution {
    public int totalWaviness(int num1, int num2) {
        int answer = 0;

        // Check every number in the range
        for (int num = num1; num <= num2; num++) {
            String s = String.valueOf(num);

            // Numbers with fewer than 3 digits have waviness 0
            if (s.length() < 3) {
                continue;
            }

            // Check every middle digit
            for (int i = 1; i < s.length() - 1; i++) {
                // Peak condition
                if (s.charAt(i) > s.charAt(i - 1) &&
                        s.charAt(i) > s.charAt(i + 1)) {
                    answer++;
                }
                // Valley condition
                else if (s.charAt(i) < s.charAt(i - 1) &&
                        s.charAt(i) < s.charAt(i + 1)) {
                    answer++;
                }
            }
        }

        return answer;
    }
}
class Solution {
    public int[] minOperations(String boxes) {
        int n = boxes.length();
        int[] answer = new int[n];

        // Left-to-right pass
        int balls = 0, operations = 0;
        for (int i = 0; i < n; i++) {
            answer[i] += operations;
            balls += (boxes.charAt(i) == '1' ? 1 : 0); // Count balls
            operations += balls; // Add the current number of balls to operations
        }

        // Right-to-left pass
        balls = 0;
        operations = 0;
        for (int i = n - 1; i >= 0; i--) {
            answer[i] += operations;
            balls += (boxes.charAt(i) == '1' ? 1 : 0); // Count balls
            operations += balls; // Add the current number of balls to operations
        }

        return answer;
    }
}

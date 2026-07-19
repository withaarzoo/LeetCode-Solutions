class Solution {
    public String smallestSubsequence(String s) {

        // Count remaining occurrences of every character
        int[] freq = new int[26];
        for (char c : s.toCharArray()) {
            freq[c - 'a']++;
        }

        // Check whether a character is already inside the stack
        boolean[] inStack = new boolean[26];

        // StringBuilder is used as a stack
        StringBuilder stack = new StringBuilder();

        for (char c : s.toCharArray()) {

            // Current occurrence has been processed
            freq[c - 'a']--;

            // Skip duplicate characters
            if (inStack[c - 'a']) {
                continue;
            }

            // Remove larger characters that will appear again
            while (stack.length() > 0 &&
                    stack.charAt(stack.length() - 1) > c &&
                    freq[stack.charAt(stack.length() - 1) - 'a'] > 0) {

                inStack[stack.charAt(stack.length() - 1) - 'a'] = false;
                stack.deleteCharAt(stack.length() - 1);
            }

            // Add current character
            stack.append(c);
            inStack[c - 'a'] = true;
        }

        return stack.toString();
    }
}
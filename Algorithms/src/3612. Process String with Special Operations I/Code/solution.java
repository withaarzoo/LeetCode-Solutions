class Solution {
    public String processStr(String s) {
        // Stores the current result being built
        StringBuilder result = new StringBuilder();

        for (char c : s.toCharArray()) {
            // Lowercase letter -> append to result
            if (c >= 'a' && c <= 'z') {
                result.append(c);
            }
            // Remove last character if it exists
            else if (c == '*') {
                if (result.length() > 0) {
                    result.deleteCharAt(result.length() - 1);
                }
            }
            // Duplicate current result
            else if (c == '#') {
                result.append(result.toString());
            }
            // Reverse current result
            else if (c == '%') {
                result.reverse();
            }
        }

        return result.toString();
    }
}
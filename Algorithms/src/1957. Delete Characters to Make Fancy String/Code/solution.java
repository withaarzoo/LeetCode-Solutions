class Solution {
    public String makeFancyString(String s) {
        StringBuilder result = new StringBuilder();
        for (char c : s.toCharArray()) {
            int n = result.length();
            if (n < 2 || !(result.charAt(n - 1) == c && result.charAt(n - 2) == c)) {
                result.append(c);
            }
        }
        return result.toString();
    }
}
class Solution {
    public boolean rotateString(String s, String goal) {
        // If lengths are different, rotation is impossible
        if (s.length() != goal.length()) return false;

        // Concatenate s with itself
        String doubled = s + s;

        // Check if goal exists in doubled string
        return doubled.contains(goal);
    }
}
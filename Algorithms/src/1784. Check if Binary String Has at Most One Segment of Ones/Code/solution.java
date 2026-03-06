class Solution {
    public boolean checkOnesSegment(String s) {
        // If "01" appears, there are multiple segments of 1s
        return !s.contains("01");
    }
}
class Solution {
    public int countBinarySubstrings(String s) {
        int n = s.length();
        
        int prevGroup = 0;   // previous group length
        int currGroup = 1;   // current group length
        int result = 0;      // answer
        
        for (int i = 1; i < n; i++) {
            
            if (s.charAt(i) == s.charAt(i - 1)) {
                currGroup++;
            } else {
                result += Math.min(prevGroup, currGroup);
                prevGroup = currGroup;
                currGroup = 1;
            }
        }
        
        result += Math.min(prevGroup, currGroup);
        
        return result;
    }
}

class Solution {
    public int distinctSubseqII(String s) {
        final long MOD = 1000000007L; // I use this modulo because the answer can be very large.
        
        long dp = 1; // I start with the empty subsequence as the only subsequence.
        long[] last = new long[26]; // last[c] stores dp from before the previous occurrence of character c.
        
        for (char c : s.toCharArray()) { // I process every character exactly once.
            int index = c - 'a'; // I convert the character into an index from 0 to 25.
            
            long oldDp = dp; // I save the old count before changing dp.
            
            dp = (2 * dp - last[index] + MOD) % MOD; // I add both choices and remove duplicate subsequences.
            
            last[index] = oldDp; // I store the old count for the next occurrence of this character.
        }
        
        return (int)((dp - 1 + MOD) % MOD); // I remove the empty subsequence from the final answer.
    }
}
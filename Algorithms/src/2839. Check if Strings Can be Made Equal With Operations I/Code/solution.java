import java.util.Arrays;

class Solution {
    public boolean canBeEqual(String s1, String s2) {
        // Create arrays for even and odd indexed characters
        char[] even1 = { s1.charAt(0), s1.charAt(2) };
        char[] even2 = { s2.charAt(0), s2.charAt(2) };

        char[] odd1 = { s1.charAt(1), s1.charAt(3) };
        char[] odd2 = { s2.charAt(1), s2.charAt(3) };

        // Sort the arrays
        Arrays.sort(even1);
        Arrays.sort(even2);
        Arrays.sort(odd1);
        Arrays.sort(odd2);

        // Compare both even and odd groups
        return Arrays.equals(even1, even2) && Arrays.equals(odd1, odd2);
    }
}
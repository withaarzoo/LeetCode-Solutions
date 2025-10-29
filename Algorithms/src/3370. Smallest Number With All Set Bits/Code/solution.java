class Solution {
    public int smallestNumber(int n) {
        // Find smallest k with (1 << k) - 1 >= n
        int k = 1;
        while (true) {
            long val = (1L << k) - 1; // use long to be safe
            if (val >= n)
                return (int) val;
            k++;
        }
    }
}

// Java (O(n log n) with TreeSet)
class Solution {
    public int[] avoidFlood(int[] rains) {
        int n = rains.length;
        int[] ans = new int[n];
        Arrays.fill(ans, 1);
        Map<Integer,Integer> last = new HashMap<>(); // lake -> last day it rained
        TreeSet<Integer> dry = new TreeSet<>();      // available dry days (sorted)

        for (int i = 0; i < n; ++i) {
            if (rains[i] > 0) {
                int lake = rains[i];
                ans[i] = -1;
                if (last.containsKey(lake)) {
                    int prev = last.get(lake);
                    Integer dryDay = dry.higher(prev); // first dry day > prev
                    if (dryDay == null) return new int[0];
                    ans[dryDay] = lake; // dry that lake on dryDay
                    dry.remove(dryDay);
                }
                last.put(lake, i);
            } else {
                dry.add(i); // collect dry day
            }
        }
        return ans;
    }
}

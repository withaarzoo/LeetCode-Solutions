class Solution {
    public int maxActiveSectionsAfterTrade(String s) {
        // Count active sections
        int ones = 0;
        for (char c : s.toCharArray()) {
            if (c == '1')
                ones++;
        }

        // Add virtual boundaries
        String t = "1" + s + "1";

        ArrayList<Character> type = new ArrayList<>();
        ArrayList<Integer> len = new ArrayList<>();

        // Run-length encoding
        for (char c : t.toCharArray()) {
            if (type.isEmpty() || type.get(type.size() - 1) != c) {
                type.add(c);
                len.add(1);
            } else {
                len.set(len.size() - 1, len.get(len.size() - 1) + 1);
            }
        }

        int best = 0;

        // Check every valid 1-block
        for (int i = 1; i + 1 < type.size(); i++) {
            if (type.get(i) == '1' &&
                    type.get(i - 1) == '0' &&
                    type.get(i + 1) == '0') {

                best = Math.max(best, len.get(i - 1) + len.get(i + 1));
            }
        }

        return ones + best;
    }
}
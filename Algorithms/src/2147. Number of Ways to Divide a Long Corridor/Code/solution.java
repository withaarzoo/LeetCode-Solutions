class Solution {
    public int numberOfWays(String corridor) {
        final int MOD = 1_000_000_007;
        List<Integer> seats = new ArrayList<>();

        for (int i = 0; i < corridor.length(); i++) {
            if (corridor.charAt(i) == 'S') {
                seats.add(i);
            }
        }

        int total = seats.size();
        if (total == 0 || total % 2 != 0)
            return 0;

        long ways = 1;
        for (int i = 2; i < total; i += 2) {
            ways = (ways * (seats.get(i) - seats.get(i - 1))) % MOD;
        }

        return (int) ways;
    }
}

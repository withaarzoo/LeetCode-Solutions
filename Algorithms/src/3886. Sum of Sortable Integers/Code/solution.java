class Solution {
    public int sortableIntegers(int[] nums) {
        int n = nums.length;
        int[] sorted = nums.clone();
        Arrays.sort(sorted);
        List<Integer> divs = new ArrayList<>();
        for (int i = 1; i * i <= n; i++) {
            if (n % i == 0) {
                divs.add(i);
                if (i != n / i) divs.add(n / i);
            }
        }
        int ans = 0;
        for (int k : divs) {
            boolean ok = true;
            for (int s = 0; s < n && ok; s += k) {
                int desc = 0, pos = -1;
                for (int j = 0; j < k; j++) {
                    int a = nums[s + j];
                    int b = nums[s + (j + 1) % k];
                    if (a > b) {
                        desc++;
                        pos = j;
                        if (desc > 1) break;
                    }
                }
                if (desc > 1) {
                    ok = false;
                    break;
                }
                int r = (desc == 0) ? 0 : (pos + 1) % k;
                for (int j = 0; j < k; j++) {
                    if (nums[s + (r + j) % k] != sorted[s + j]) {
                        ok = false;
                        break;
                    }
                }
            }
            if (ok) ans += k;
        }
        return ans;
    }
}
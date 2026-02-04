class Solution {

    static class Block {
        int l, r;
        long sum;

        Block(int l, int r, long sum) {
            this.l = l;
            this.r = r;
            this.sum = sum;
        }
    }

    List<Block> decompose(int[] nums) {
        int n = nums.length;
        List<Block> list = new ArrayList<>();

        int l = 0;
        long s = nums[0];

        for (int i = 1; i < n; i++) {
            if (nums[i - 1] <= nums[i]) {
                list.add(new Block(l, i - 1, s));
                l = i;
                s = 0;
            }
            s += nums[i];
        }
        list.add(new Block(l, n - 1, s));
        return list;
    }

    public long maxSumTrionic(int[] nums) {
        int n = nums.length;

        long[] left = new long[n];
        long[] right = new long[n];

        for (int i = 0; i < n; i++) {
            left[i] = nums[i];
            if (i > 0 && nums[i - 1] < nums[i] && left[i - 1] > 0) {
                left[i] += left[i - 1];
            }
        }

        for (int i = n - 1; i >= 0; i--) {
            right[i] = nums[i];
            if (i + 1 < n && nums[i] < nums[i + 1] && right[i + 1] > 0) {
                right[i] += right[i + 1];
            }
        }

        long ans = Long.MIN_VALUE;
        for (Block b : decompose(nums)) {
            if (b.l > 0 && b.r < n - 1 &&
                    nums[b.l - 1] < nums[b.l] &&
                    nums[b.r] < nums[b.r + 1] &&
                    b.l < b.r) {
                ans = Math.max(ans, left[b.l - 1] + b.sum + right[b.r + 1]);
            }
        }
        return ans;
    }
}

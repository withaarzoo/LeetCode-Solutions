class Solution {

    private int getMaxGap(int[] bars) {
        Arrays.sort(bars);

        int maxLen = 1;
        int curLen = 1;

        for (int i = 1; i < bars.length; i++) {
            if (bars[i] == bars[i - 1] + 1) {
                curLen++;
            } else {
                curLen = 1;
            }
            maxLen = Math.max(maxLen, curLen);
        }
        return maxLen;
    }

    public int maximizeSquareHoleArea(int n, int m, int[] hBars, int[] vBars) {
        int hGap = getMaxGap(hBars) + 1;
        int vGap = getMaxGap(vBars) + 1;

        int side = Math.min(hGap, vGap);
        return side * side;
    }
}

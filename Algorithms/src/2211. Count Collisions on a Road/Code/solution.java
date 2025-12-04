class Solution {
    public int countCollisions(String directions) {
        int n = directions.length();
        int i = 0, j = n - 1;

        // Skip all leading 'L' cars (safe, no collision)
        while (i < n && directions.charAt(i) == 'L') {
            i++;
        }

        // Skip all trailing 'R' cars (safe, no collision)
        while (j >= 0 && directions.charAt(j) == 'R') {
            j--;
        }

        int collisions = 0;
        // Count all moving cars ('L' or 'R') in the middle part
        for (int k = i; k <= j; k++) {
            if (directions.charAt(k) != 'S') {
                collisions++;
            }
        }

        return collisions;
    }
}

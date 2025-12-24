class Solution {
    public int minimumBoxes(int[] apple, int[] capacity) {
        // Step 1: Calculate total apples
        int totalApples = 0;
        for (int a : apple) {
            totalApples += a;
        }

        // Step 2: Sort capacities
        Arrays.sort(capacity);

        // Step 3: Pick boxes from largest to smallest
        int usedCapacity = 0;
        int boxes = 0;

        for (int i = capacity.length - 1; i >= 0; i--) {
            usedCapacity += capacity[i];
            boxes++;
            if (usedCapacity >= totalApples) {
                return boxes;
            }
        }

        return boxes;
    }
}

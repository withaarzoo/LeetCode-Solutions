class Solution {
    public List<Integer> survivedRobotsHealths(int[] positions, int[] healths, String directions) {
        int n = positions.length;

        Integer[] indices = new Integer[n];
        for (int i = 0; i < n; i++) {
            indices[i] = i;
        }

        // Sort indices based on positions
        Arrays.sort(indices, (a, b) -> positions[a] - positions[b]);

        Stack<Integer> stack = new Stack<>();

        for (int idx : indices) {
            // Robot moving right
            if (directions.charAt(idx) == 'R') {
                stack.push(idx);
            } else {
                // Robot moving left
                while (!stack.isEmpty() && healths[idx] > 0) {
                    int topIdx = stack.peek();

                    if (healths[topIdx] < healths[idx]) {
                        stack.pop();
                        healths[idx]--;
                        healths[topIdx] = 0;
                    } else if (healths[topIdx] == healths[idx]) {
                        stack.pop();
                        healths[topIdx] = 0;
                        healths[idx] = 0;
                    } else {
                        healths[topIdx]--;
                        healths[idx] = 0;
                    }
                }
            }
        }

        List<Integer> result = new ArrayList<>();

        for (int health : healths) {
            if (health > 0) {
                result.add(health);
            }
        }

        return result;
    }
}
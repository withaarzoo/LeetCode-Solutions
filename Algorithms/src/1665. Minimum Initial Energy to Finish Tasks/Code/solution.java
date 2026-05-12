class Solution {
    public int minimumEffort(int[][] tasks) {

        // Sort by (minimum - actual) in descending order
        Arrays.sort(tasks, (a, b) -> (b[1] - b[0]) - (a[1] - a[0]));

        int answer = 0; // Minimum starting energy
        int energy = 0; // Current energy

        // Process every task
        for (int[] task : tasks) {

            int actual = task[0];
            int minimum = task[1];

            // Add extra energy if needed
            if (energy < minimum) {

                int need = minimum - energy;

                answer += need;
                energy += need;
            }

            // Finish the task
            energy -= actual;
        }

        return answer;
    }
}
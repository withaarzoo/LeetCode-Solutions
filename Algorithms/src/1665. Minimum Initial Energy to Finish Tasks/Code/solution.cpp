class Solution {
public:
    int minimumEffort(vector<vector<int>>& tasks) {

        // Sort tasks by (minimum - actual) in descending order
        sort(tasks.begin(), tasks.end(), [](vector<int>& a, vector<int>& b) {
            return (a[1] - a[0]) > (b[1] - b[0]);
        });

        int answer = 0; // Minimum initial energy required
        int energy = 0; // Current available energy

        // Process tasks one by one
        for (auto& task : tasks) {

            int actual = task[0];
            int minimum = task[1];

            // If current energy is not enough,
            // add the missing amount
            if (energy < minimum) {

                int need = minimum - energy;

                answer += need;
                energy += need;
            }

            // Complete the task
            energy -= actual;
        }

        return answer;
    }
};
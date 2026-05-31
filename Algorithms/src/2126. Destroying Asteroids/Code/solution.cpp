class Solution {
public:
    bool asteroidsDestroyed(int mass, vector<int>& asteroids) {
        // Sort asteroids from smallest to largest
        sort(asteroids.begin(), asteroids.end());

        // Use long long because mass can become very large
        long long currentMass = mass;

        // Try destroying asteroids one by one
        for (int asteroid : asteroids) {

            // If planet is too small, it gets destroyed
            if (currentMass < asteroid) {
                return false;
            }

            // Gain the asteroid's mass
            currentMass += asteroid;
        }

        // All asteroids were destroyed
        return true;
    }
};
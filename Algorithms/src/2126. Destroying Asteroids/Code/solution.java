class Solution {
    public boolean asteroidsDestroyed(int mass, int[] asteroids) {

        // Sort asteroids from smallest to largest
        Arrays.sort(asteroids);

        // Use long because mass can exceed int range
        long currentMass = mass;

        // Process each asteroid
        for (int asteroid : asteroids) {

            // Planet cannot destroy this asteroid
            if (currentMass < asteroid) {
                return false;
            }

            // Gain asteroid mass
            currentMass += asteroid;
        }

        // Successfully destroyed all asteroids
        return true;
    }
}
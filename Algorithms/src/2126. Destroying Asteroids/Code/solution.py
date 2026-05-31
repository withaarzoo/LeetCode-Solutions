class Solution:
    def asteroidsDestroyed(self, mass: int, asteroids: List[int]) -> bool:

        # Sort asteroids from smallest to largest
        asteroids.sort()

        # Current planet mass
        current_mass = mass

        # Try destroying each asteroid
        for asteroid in asteroids:

            # Planet is too small
            if current_mass < asteroid:
                return False

            # Gain asteroid mass
            current_mass += asteroid

        # All asteroids destroyed
        return True
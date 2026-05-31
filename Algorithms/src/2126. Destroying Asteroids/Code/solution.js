/**
 * @param {number} mass
 * @param {number[]} asteroids
 * @return {boolean}
 */
var asteroidsDestroyed = function (mass, asteroids) {
  // Sort asteroids in ascending order
  asteroids.sort((a, b) => a - b);

  // Store current planet mass
  let currentMass = mass;

  // Process every asteroid
  for (const asteroid of asteroids) {
    // Cannot destroy this asteroid
    if (currentMass < asteroid) {
      return false;
    }

    // Gain its mass
    currentMass += asteroid;
  }

  // All asteroids destroyed
  return true;
};

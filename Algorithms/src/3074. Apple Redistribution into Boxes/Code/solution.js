var minimumBoxes = function (apple, capacity) {
  // Step 1: Total apples
  let totalApples = apple.reduce((sum, a) => sum + a, 0);

  // Step 2: Sort capacities in descending order
  capacity.sort((a, b) => b - a);

  // Step 3: Use boxes greedily
  let usedCapacity = 0;
  let boxes = 0;

  for (let cap of capacity) {
    usedCapacity += cap;
    boxes++;
    if (usedCapacity >= totalApples) {
      return boxes;
    }
  }

  return boxes;
};

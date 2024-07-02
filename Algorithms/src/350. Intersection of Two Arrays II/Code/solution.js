var intersect = function (nums1, nums2) {
  // Step 1: Initialize a Map to store counts of each number in nums1
  let countMap = new Map();
  // Step 2: Initialize an array to store the intersection result
  let result = [];

  // Step 3: Count occurrences of each number in nums1 using forEach
  nums1.forEach((num) => {
    // Increase the count of num in countMap by 1
    countMap.set(num, (countMap.get(num) || 0) + 1);
  });

  // Step 4: Iterate through nums2 to find intersecting elements
  nums2.forEach((num) => {
    // Check if num exists in countMap and its count is greater than 0
    if (countMap.has(num) && countMap.get(num) > 0) {
      // Add num to result since it's an intersecting element
      result.push(num);
      // Decrease the count of num in countMap by 1 to mark its usage
      countMap.set(num, countMap.get(num) - 1);
    }
  });

  // Step 5: Return the resulting array containing intersecting elements
  return result;
};

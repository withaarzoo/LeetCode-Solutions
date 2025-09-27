/**
 * @param {number[][]} points
 * @return {number}
 */
var largestTriangleArea = function (points) {
  const n = points.length;
  let maxArea = 0;
  // all triples i < j < k
  for (let i = 0; i < n - 2; i++) {
    for (let j = i + 1; j < n - 1; j++) {
      for (let k = j + 1; k < n; k++) {
        const [x1, y1] = points[i];
        const [x2, y2] = points[j];
        const [x3, y3] = points[k];
        // shoelace / cross-product formula (double area)
        const doubled = Math.abs(
          x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2)
        );
        const area = doubled * 0.5;
        if (area > maxArea) maxArea = area;
      }
    }
  }
  return maxArea;
};

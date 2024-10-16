var longestDiverseString = function (a, b, c) {
  // Priority queue to keep the characters sorted by count.
  const pq = [];
  if (a > 0) pq.push([a, "a"]);
  if (b > 0) pq.push([b, "b"]);
  if (c > 0) pq.push([c, "c"]);
  pq.sort((x, y) => y[0] - x[0]); // Sort in descending order.

  let result = "";

  while (pq.length > 0) {
    let [count1, char1] = pq.shift();

    // If last two characters are the same as char1, pick the next one.
    if (
      result.length >= 2 &&
      result[result.length - 1] === char1 &&
      result[result.length - 2] === char1
    ) {
      if (pq.length === 0) break;

      let [count2, char2] = pq.shift();
      result += char2;
      if (--count2 > 0) pq.push([count2, char2]);

      pq.push([count1, char1]); // Push char1 back for later use.
    } else {
      result += char1;
      if (--count1 > 0) pq.push([count1, char1]);
    }

    pq.sort((x, y) => y[0] - x[0]); // Sort after each modification.
  }

  return result;
};

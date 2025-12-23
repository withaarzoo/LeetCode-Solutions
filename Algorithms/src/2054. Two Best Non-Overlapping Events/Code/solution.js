var maxTwoEvents = function (events) {
  // Sort by start time
  events.sort((a, b) => a[0] - b[0]);

  // Sort another copy by end time
  const endSorted = [...events].sort((a, b) => a[1] - b[1]);

  const n = events.length;
  const maxValueTill = new Array(n);

  maxValueTill[0] = endSorted[0][2];
  for (let i = 1; i < n; i++) {
    maxValueTill[i] = Math.max(maxValueTill[i - 1], endSorted[i][2]);
  }

  let ans = 0;
  let j = 0;

  for (let i = 0; i < n; i++) {
    const [start, , value] = events[i];

    while (j < n && endSorted[j][1] < start) {
      j++;
    }

    ans = Math.max(ans, value);
    if (j > 0) {
      ans = Math.max(ans, value + maxValueTill[j - 1]);
    }
  }

  return ans;
};

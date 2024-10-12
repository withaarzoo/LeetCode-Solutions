/**
 * @param {number[][]} intervals
 * @return {number}
 */
var minGroups = function (intervals) {
  const events = [];

  // Separate start and end points into events
  for (const [start, end] of intervals) {
    events.push([start, 1]); // 1 indicates start of interval
    events.push([end + 1, -1]); // -1 indicates end of interval
  }

  // Sort events: by time first, then by type (-1 ends come before 1 starts)
  events.sort((a, b) => (a[0] === b[0] ? a[1] - b[1] : a[0] - b[0]));

  let maxGroups = 0,
    currentGroups = 0;

  // Process events
  for (const [time, type] of events) {
    currentGroups += type;
    maxGroups = Math.max(maxGroups, currentGroups);
  }

  return maxGroups;
};

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {number[]}
 */
var nodesBetweenCriticalPoints = function (head) {
  // Store the position of the first critical point.
  let first = -1;

  // Store the position of the most recent critical point.
  let last = -1;

  // Start with Infinity so the first valid distance can replace it.
  let minDistance = Infinity;

  // The current node starts at position 1 because head is position 0.
  let position = 1;

  // prev starts at the head because the current node needs a previous node.
  let prev = head;

  // curr starts from the second node.
  let curr = head.next;

  // The last node cannot be critical because it has no next node.
  while (curr !== null && curr.next !== null) {
    // Check whether curr is a local maximum or local minimum.
    const isCritical =
      (curr.val > prev.val && curr.val > curr.next.val) ||
      (curr.val < prev.val && curr.val < curr.next.val);

    // Process the node only when it is a critical point.
    if (isCritical) {
      // Save this position if it is the first critical point.
      if (first === -1) {
        first = position;
      } else {
        // Update the minimum distance from the previous critical point.
        minDistance = Math.min(minDistance, position - last);
      }

      // Store this critical point as the latest one.
      last = position;
    }

    // Move prev forward for the next comparison.
    prev = curr;

    // Move curr forward to the next node.
    curr = curr.next;

    // Move to the next position.
    position++;
  }

  // Fewer than two critical points means no valid answer exists.
  if (first === -1 || first === last) {
    return [-1, -1];
  }

  // The first-to-last distance is always the maximum distance.
  const maxDistance = last - first;

  // Return the minimum and maximum distances.
  return [minDistance, maxDistance];
};

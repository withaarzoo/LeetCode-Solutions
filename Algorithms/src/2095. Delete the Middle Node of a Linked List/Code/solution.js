/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteMiddle = function (head) {
  // If there is only one node, return empty list
  if (head.next === null) {
    return null;
  }

  // Slow finds middle, fast moves twice as fast
  let slow = head;
  let fast = head;

  // Node before slow
  let prev = null;

  while (fast !== null && fast.next !== null) {
    prev = slow; // Store previous node
    slow = slow.next; // Move slow by 1 step
    fast = fast.next.next; // Move fast by 2 steps
  }

  // Remove middle node
  prev.next = slow.next;

  return head;
};

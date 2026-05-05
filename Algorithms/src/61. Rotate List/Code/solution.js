/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
var rotateRight = function (head, k) {
  // edge case
  if (!head || !head.next || k === 0) return head;

  let n = 1;
  let tail = head;

  // find length
  while (tail.next) {
    tail = tail.next;
    n++;
  }

  // make circular
  tail.next = head;

  // reduce k
  k = k % n;

  let steps = n - k - 1;
  let newTail = head;

  // move to new tail
  while (steps-- > 0) {
    newTail = newTail.next;
  }

  let newHead = newTail.next;

  // break circle
  newTail.next = null;

  return newHead;
};

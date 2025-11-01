/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {number[]} nums
 * @param {ListNode} head
 * @return {ListNode}
 */
var modifiedList = function (nums, head) {
  // build a set for O(1) checks
  const toDelete = new Set(nums);

  // dummy node to handle head deletions easily
  const dummy = new ListNode(0, head);
  let prev = dummy;
  let curr = head;

  while (curr !== null) {
    if (toDelete.has(curr.val)) {
      // skip current node
      prev.next = curr.next;
    } else {
      // keep current node
      prev = curr;
    }
    curr = curr.next;
  }
  return dummy.next;
};

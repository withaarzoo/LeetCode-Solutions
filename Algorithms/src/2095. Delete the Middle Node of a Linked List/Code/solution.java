/**
 * Definition for singly-linked list.
 * public class ListNode {
 * int val;
 * ListNode next;
 * ListNode() {}
 * ListNode(int val) { this.val = val; }
 * ListNode(int val, ListNode next) { this.next = next; }
 * }
 */
class Solution {
    public ListNode deleteMiddle(ListNode head) {

        // If there is only one node, return an empty list
        if (head.next == null) {
            return null;
        }

        // Slow finds middle, fast moves twice as fast
        ListNode slow = head;
        ListNode fast = head;

        // Node before slow
        ListNode prev = null;

        while (fast != null && fast.next != null) {
            prev = slow; // Store previous node
            slow = slow.next; // Move slow by 1 step
            fast = fast.next.next; // Move fast by 2 steps
        }

        // Remove middle node
        prev.next = slow.next;

        return head;
    }
}
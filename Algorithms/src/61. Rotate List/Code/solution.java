/**
 * Definition for singly-linked list.
 * public class ListNode {
 * int val;
 * ListNode next;
 * ListNode() {}
 * ListNode(int val) { this.val = val; }
 * ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode rotateRight(ListNode head, int k) {
        // edge case
        if (head == null || head.next == null || k == 0)
            return head;

        int n = 1;
        ListNode tail = head;

        // find length
        while (tail.next != null) {
            tail = tail.next;
            n++;
        }

        // make circular
        tail.next = head;

        // reduce k
        k = k % n;

        int steps = n - k - 1;
        ListNode newTail = head;

        // move to new tail
        while (steps-- > 0) {
            newTail = newTail.next;
        }

        ListNode newHead = newTail.next;

        // break circle
        newTail.next = null;

        return newHead;
    }
}
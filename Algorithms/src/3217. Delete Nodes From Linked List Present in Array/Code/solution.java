
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
import java.util.HashSet;
import java.util.Set;

class Solution {
    public ListNode modifiedList(int[] nums, ListNode head) {
        // Build a set for quick membership checks
        Set<Integer> toDelete = new HashSet<>();
        for (int x : nums)
            toDelete.add(x);

        // Dummy node to simplify head deletions
        ListNode dummy = new ListNode(0, head);
        ListNode prev = dummy;
        ListNode curr = head;

        while (curr != null) {
            if (toDelete.contains(curr.val)) {
                // skip curr
                prev.next = curr.next;
            } else {
                // keep curr
                prev = curr;
            }
            curr = curr.next;
        }
        return dummy.next;
    }
}

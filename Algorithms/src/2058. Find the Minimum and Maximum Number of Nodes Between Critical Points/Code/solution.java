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
    public int[] nodesBetweenCriticalPoints(ListNode head) {
        // Store the position of the first critical point.
        int first = -1;

        // Store the position of the most recent critical point.
        int last = -1;

        // Start with the largest possible integer for the minimum distance.
        int minDistance = Integer.MAX_VALUE;

        // The current node starts at position 1 because head is position 0.
        int position = 1;

        // The previous node is initially the head.
        ListNode prev = head;

        // The current node starts from the second node.
        ListNode curr = head.next;

        // The last node cannot be critical because it has no next node.
        while (curr != null && curr.next != null) {
            // Check whether the current node is a local maximum or local minimum.
            boolean isCritical = (curr.val > prev.val && curr.val > curr.next.val) ||
                    (curr.val < prev.val && curr.val < curr.next.val);

            // Process the current node if it is a critical point.
            if (isCritical) {
                // Save the position when this is the first critical point.
                if (first == -1) {
                    first = position;
                } else {
                    // Update the minimum distance using consecutive critical points.
                    minDistance = Math.min(minDistance, position - last);
                }

                // Make the current critical point the latest critical point.
                last = position;
            }

            // Move prev to the current node for the next iteration.
            prev = curr;

            // Move curr to the next node.
            curr = curr.next;

            // Move to the next position.
            position++;
        }

        // If there are fewer than two critical points, no distance can be calculated.
        if (first == -1 || first == last) {
            return new int[] { -1, -1 };
        }

        // The distance between the first and last critical points is the maximum.
        int maxDistance = last - first;

        // Return the minimum and maximum distances.
        return new int[] { minDistance, maxDistance };
    }
}
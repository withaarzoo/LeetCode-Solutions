# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

from typing import Optional, List

class Solution:
    def nodesBetweenCriticalPoints(self, head: Optional[ListNode]) -> List[int]:
        # Store the position of the first critical point.
        first = -1

        # Store the position of the most recent critical point.
        last = -1

        # Start with infinity so the first valid distance can replace it.
        min_distance = float("inf")

        # The current node starts at position 1 because head is position 0.
        position = 1

        # prev starts at head because curr needs a previous node.
        prev = head

        # curr starts from the second node.
        curr = head.next

        # The last node cannot be critical because it has no next node.
        while curr is not None and curr.next is not None:
            # Check whether curr is a local maximum or local minimum.
            is_critical = (
                (curr.val > prev.val and curr.val > curr.next.val)
                or
                (curr.val < prev.val and curr.val < curr.next.val)
            )

            # Process the node only if it is a critical point.
            if is_critical:
                # Save this position if it is the first critical point.
                if first == -1:
                    first = position
                else:
                    # Update the minimum distance from the previous critical point.
                    min_distance = min(min_distance, position - last)

                # Store this critical point as the latest one.
                last = position

            # Move prev forward for the next comparison.
            prev = curr

            # Move curr forward to the next node.
            curr = curr.next

            # Move to the next position.
            position += 1

        # Fewer than two critical points means no valid answer exists.
        if first == -1 or first == last:
            return [-1, -1]

        # The first-to-last distance is always the maximum distance.
        max_distance = last - first

        # Return the minimum and maximum distances.
        return [min_distance, max_distance]
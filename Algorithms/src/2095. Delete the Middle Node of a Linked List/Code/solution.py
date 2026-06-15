# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:

        # If there is only one node, deleting it leaves an empty list
        if head.next is None:
            return None

        # Slow finds middle, fast moves twice as fast
        slow = head
        fast = head

        # Node before slow
        prev = None

        while fast and fast.next:
            prev = slow          # Store previous node
            slow = slow.next     # Move slow by 1 step
            fast = fast.next.next  # Move fast by 2 steps

        # Remove middle node
        prev.next = slow.next

        return head
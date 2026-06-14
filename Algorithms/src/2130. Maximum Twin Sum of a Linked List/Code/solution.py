# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def pairSum(self, head: Optional[ListNode]) -> int:

        # Find the middle of the linked list
        slow = head
        fast = head

        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

        # Reverse the second half
        prev = None

        while slow:
            next_node = slow.next
            slow.next = prev
            prev = slow
            slow = next_node

        # Calculate maximum twin sum
        ans = 0
        first = head
        second = prev

        while second:
            ans = max(ans, first.val + second.val)

            first = first.next
            second = second.next

        return ans
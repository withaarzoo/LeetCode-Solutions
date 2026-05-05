# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        # edge case
        if not head or not head.next or k == 0:
            return head

        n = 1
        tail = head

        # find length
        while tail.next:
            tail = tail.next
            n += 1

        # make circular
        tail.next = head

        # reduce k
        k = k % n

        steps = n - k - 1
        newTail = head

        # move to new tail
        while steps > 0:
            newTail = newTail.next
            steps -= 1

        newHead = newTail.next

        # break circle
        newTail.next = None

        return newHead
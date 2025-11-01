# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def modifiedList(self, nums: List[int], head: Optional[ListNode]) -> Optional[ListNode]:
        # Build a set for constant-time membership checks
        to_delete = set(nums)
        
        # Dummy node to simplify deletions at head
        dummy = ListNode(0, head)
        prev, curr = dummy, head
        
        while curr:
            if curr.val in to_delete:
                # Remove curr by skipping it
                prev.next = curr.next
            else:
                # Keep curr, move prev forward
                prev = curr
            curr = curr.next
        
        return dummy.next

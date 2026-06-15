/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {

    // If there is only one node, return an empty list
    if head.Next == nil {
        return nil
    }

    // Slow finds middle, fast moves twice as fast
    slow := head
    fast := head

    // Node before slow
    var prev *ListNode = nil

    for fast != nil && fast.Next != nil {
        prev = slow           // Store previous node
        slow = slow.Next      // Move slow by 1 step
        fast = fast.Next.Next // Move fast by 2 steps
    }

    // Remove middle node
    prev.Next = slow.Next

    return head
}
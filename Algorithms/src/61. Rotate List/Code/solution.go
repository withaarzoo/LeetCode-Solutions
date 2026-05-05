/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    // edge case
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    n := 1
    tail := head

    // find length
    for tail.Next != nil {
        tail = tail.Next
        n++
    }

    // make circular
    tail.Next = head

    // reduce k
    k = k % n

    steps := n - k - 1
    newTail := head

    // move to new tail
    for steps > 0 {
        newTail = newTail.Next
        steps--
    }

    newHead := newTail.Next

    // break circle
    newTail.Next = nil

    return newHead
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    // build a map[int]struct{} as a set
    toDelete := make(map[int]struct{}, len(nums))
    for _, v := range nums {
        toDelete[v] = struct{}{}
    }

    // dummy node to handle head deletions easily
    dummy := &ListNode{Val: 0, Next: head}
    prev := dummy
    curr := head

    for curr != nil {
        if _, ok := toDelete[curr.Val]; ok {
            // skip curr
            prev.Next = curr.Next
        } else {
            // keep curr
            prev = curr
        }
        curr = curr.Next
    }
    return dummy.Next
}

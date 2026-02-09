func balanceBST(root *TreeNode) *TreeNode {
    arr := []int{}

    // Inorder traversal
    var inorder func(*TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        arr = append(arr, node.Val)
        inorder(node.Right)
    }

    // Build balanced BST
    var build func(int, int) *TreeNode
    build = func(left, right int) *TreeNode {
        if left > right {
            return nil
        }

        mid := (left + right) / 2
        node := &TreeNode{Val: arr[mid]}
        node.Left = build(left, mid-1)
        node.Right = build(mid+1, right)

        return node
    }

    inorder(root)
    return build(0, len(arr)-1)
}

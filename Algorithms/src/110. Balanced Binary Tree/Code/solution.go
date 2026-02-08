func isBalanced(root *TreeNode) bool {

    var height func(node *TreeNode) int
    height = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftHeight := height(node.Left)
        if leftHeight == -1 {
            return -1
        }

        rightHeight := height(node.Right)
        if rightHeight == -1 {
            return -1
        }

        if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
            return -1
        }

        if leftHeight > rightHeight {
            return 1 + leftHeight
        }
        return 1 + rightHeight
    }

    return height(root) != -1
}

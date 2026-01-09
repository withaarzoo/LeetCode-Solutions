func subtreeWithAllDeepest(root *TreeNode) *TreeNode {

    var dfs func(*TreeNode) (int, *TreeNode)

    dfs = func(node *TreeNode) (int, *TreeNode) {
        if node == nil {
            return 0, nil
        }

        ld, ln := dfs(node.Left)
        rd, rn := dfs(node.Right)

        if ld == rd {
            return ld + 1, node
        } else if ld > rd {
            return ld + 1, ln
        } else {
            return rd + 1, rn
        }
    }

    _, ans := dfs(root)
    return ans
}

func sumRootToLeaf(root *TreeNode) int {
    
    var dfs func(node *TreeNode, current int) int
    
    dfs = func(node *TreeNode, current int) int {
        if node == nil {
            return 0
        }
        
        // Build binary number
        current = current*2 + node.Val
        
        // If leaf
        if node.Left == nil && node.Right == nil {
            return current
        }
        
        // Sum left and right
        return dfs(node.Left, current) + dfs(node.Right, current)
    }
    
    return dfs(root, 0)
}
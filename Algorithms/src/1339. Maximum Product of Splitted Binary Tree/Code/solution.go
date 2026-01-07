func maxProduct(root *TreeNode) int {
    const MOD int64 = 1e9 + 7
    var totalSum int64 = 0
    var maxProd int64 = 0

    var getTotalSum func(*TreeNode) int64
    getTotalSum = func(node *TreeNode) int64 {
        if node == nil {
            return 0
        }
        return int64(node.Val) + getTotalSum(node.Left) + getTotalSum(node.Right)
    }

    var dfs func(*TreeNode) int64
    dfs = func(node *TreeNode) int64 {
        if node == nil {
            return 0
        }

        left := dfs(node.Left)
        right := dfs(node.Right)

        subtreeSum := int64(node.Val) + left + right
        product := subtreeSum * (totalSum - subtreeSum)
        if product > maxProd {
            maxProd = product
        }

        return subtreeSum
    }

    totalSum = getTotalSum(root)
    dfs(root)

    return int(maxProd % MOD)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func averageOfSubtree(root *TreeNode) int {
    answer := 0 // Stores the number of nodes equal to their subtree average.

    var dfs func(*TreeNode) (int, int) // Defines a recursive function returning subtree sum and count.

    dfs = func(node *TreeNode) (int, int) {
        if node == nil {
            return 0, 0 // An empty subtree has sum 0 and contains 0 nodes.
        }

        leftSum, leftCount := dfs(node.Left) // Get the sum and count of the left subtree.
        rightSum, rightCount := dfs(node.Right) // Get the sum and count of the right subtree.

        sum := leftSum + rightSum + node.Val // Calculate the sum of the current subtree.
        count := leftCount + rightCount + 1 // Calculate the number of nodes in the current subtree.

        if node.Val == sum/count {
            answer++ // Count the node when its value equals the floored average.
        }

        return sum, count // Return the current subtree's sum and count to its parent.
    }

    dfs(root) // Process every node using postorder traversal.
    return answer // Return the final count.
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    
    // Stores value -> TreeNode mapping
    nodes := make(map[int]*TreeNode)
    
    // Stores all child values
    children := make(map[int]bool)
    
    for _, d := range descriptions {
        parent := d[0]
        child := d[1]
        isLeft := d[2]
        
        // Create parent node if needed
        if _, exists := nodes[parent]; !exists {
            nodes[parent] = &TreeNode{Val: parent}
        }
        
        // Create child node if needed
        if _, exists := nodes[child]; !exists {
            nodes[child] = &TreeNode{Val: child}
        }
        
        // Attach child to correct side
        if isLeft == 1 {
            nodes[parent].Left = nodes[child]
        } else {
            nodes[parent].Right = nodes[child]
        }
        
        // Mark child
        children[child] = true
    }
    
    // Find root node
    for value, node := range nodes {
        if !children[value] {
            return node
        }
    }
    
    return nil
}
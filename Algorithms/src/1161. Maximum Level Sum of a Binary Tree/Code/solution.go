func maxLevelSum(root *TreeNode) int {
    queue := []*TreeNode{root}
    level := 1
    answerLevel := 1
    maxSum := math.MinInt64

    for len(queue) > 0 {
        size := len(queue)
        levelSum := 0

        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]

            levelSum += node.Val

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        if levelSum > maxSum {
            maxSum = levelSum
            answerLevel = level
        }

        level++
    }

    return answerLevel
}

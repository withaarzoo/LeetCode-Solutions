class Solution:
    def maxLevelSum(self, root: Optional[TreeNode]) -> int:
        from collections import deque

        queue = deque([root])
        level = 1
        answer_level = 1
        max_sum = float('-inf')

        while queue:
            size = len(queue)
            level_sum = 0

            for _ in range(size):
                node = queue.popleft()
                level_sum += node.val

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

            if level_sum > max_sum:
                max_sum = level_sum
                answer_level = level

            level += 1

        return answer_level

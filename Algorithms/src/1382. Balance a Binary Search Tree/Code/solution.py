class Solution:
    def balanceBST(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        arr = []

        # Inorder traversal
        def inorder(node):
            if not node:
                return
            inorder(node.left)
            arr.append(node.val)
            inorder(node.right)

        # Build balanced BST
        def build(left, right):
            if left > right:
                return None

            mid = (left + right) // 2
            node = TreeNode(arr[mid])
            node.left = build(left, mid - 1)
            node.right = build(mid + 1, right)

            return node

        inorder(root)
        return build(0, len(arr) - 1)

class Solution:
    def subtreeWithAllDeepest(self, root):
        
        def dfs(node):
            if not node:
                return 0, None
            
            ld, ln = dfs(node.left)
            rd, rn = dfs(node.right)

            if ld == rd:
                return ld + 1, node
            elif ld > rd:
                return ld + 1, ln
            else:
                return rd + 1, rn

        return dfs(root)[1]

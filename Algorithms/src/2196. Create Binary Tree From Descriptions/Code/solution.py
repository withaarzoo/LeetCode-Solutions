# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def createBinaryTree(self, descriptions: List[List[int]]) -> Optional[TreeNode]:
        
        # Stores value -> TreeNode mapping
        nodes = {}
        
        # Stores all child values
        children = set()
        
        for parent, child, isLeft in descriptions:
            
            # Create parent node if needed
            if parent not in nodes:
                nodes[parent] = TreeNode(parent)
            
            # Create child node if needed
            if child not in nodes:
                nodes[child] = TreeNode(child)
            
            # Connect child to correct side
            if isLeft:
                nodes[parent].left = nodes[child]
            else:
                nodes[parent].right = nodes[child]
            
            # Mark child
            children.add(child)
        
        # Root never appears as a child
        for value, node in nodes.items():
            if value not in children:
                return node
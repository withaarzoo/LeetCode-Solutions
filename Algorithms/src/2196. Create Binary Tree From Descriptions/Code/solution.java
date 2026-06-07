/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode() {}
 * TreeNode(int val) { this.val = val; }
 * TreeNode(int val, TreeNode left, TreeNode right) {
 * this.val = val;
 * this.left = left;
 * this.right = right;
 * }
 * }
 */
class Solution {
    public TreeNode createBinaryTree(int[][] descriptions) {

        // Stores value -> TreeNode mapping
        Map<Integer, TreeNode> nodes = new HashMap<>();

        // Stores all child values
        Set<Integer> children = new HashSet<>();

        for (int[] d : descriptions) {
            int parent = d[0];
            int child = d[1];
            int isLeft = d[2];

            // Create parent node if needed
            nodes.putIfAbsent(parent, new TreeNode(parent));

            // Create child node if needed
            nodes.putIfAbsent(child, new TreeNode(child));

            // Attach child to correct side
            if (isLeft == 1)
                nodes.get(parent).left = nodes.get(child);
            else
                nodes.get(parent).right = nodes.get(child);

            // Mark child
            children.add(child);
        }

        // Find root
        for (int value : nodes.keySet()) {
            if (!children.contains(value))
                return nodes.get(value);
        }

        return null;
    }
}
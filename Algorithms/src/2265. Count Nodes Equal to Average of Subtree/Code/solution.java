/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    private int answer = 0; // Stores the number of nodes equal to their subtree average.

    private int[] dfs(TreeNode node) {
        if (node == null) {
            return new int[]{0, 0}; // An empty subtree has sum 0 and count 0.
        }

        int[] left = dfs(node.left); // Get the sum and count of the left subtree.
        int[] right = dfs(node.right); // Get the sum and count of the right subtree.

        int sum = left[0] + right[0] + node.val; // Combine both subtree sums with the current value.
        int count = left[1] + right[1] + 1; // Combine both subtree counts and include the current node.

        if (node.val == sum / count) {
            answer++; // Count the node when its value equals the floored average.
        }

        return new int[]{sum, count}; // Return this subtree's sum and count to its parent.
    }

    public int averageOfSubtree(TreeNode root) {
        dfs(root); // Traverse the entire tree and calculate each subtree's information.
        return answer; // Return the total number of valid nodes.
    }
}
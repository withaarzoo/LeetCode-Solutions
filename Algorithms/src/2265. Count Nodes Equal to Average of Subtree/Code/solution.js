/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

/**
 * @param {TreeNode} root
 * @return {number}
 */
var averageOfSubtree = function (root) {
  let answer = 0; // Stores the number of nodes whose value equals their subtree average.

  function dfs(node) {
    if (node === null) {
      return [0, 0]; // An empty subtree has sum 0 and contains 0 nodes.
    }

    const [leftSum, leftCount] = dfs(node.left); // Get sum and count from the left subtree.
    const [rightSum, rightCount] = dfs(node.right); // Get sum and count from the right subtree.

    const sum = leftSum + rightSum + node.val; // Calculate the sum of the current subtree.
    const count = leftCount + rightCount + 1; // Calculate the number of nodes in the current subtree.

    if (node.val === Math.floor(sum / count)) {
      answer++; // Count the node when its value equals the floored average.
    }

    return [sum, count]; // Return the current subtree's sum and count to its parent.
  }

  dfs(root); // Process every node using postorder traversal.
  return answer; // Return the final count.
};

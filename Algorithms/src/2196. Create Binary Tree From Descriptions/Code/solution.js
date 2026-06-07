/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[][]} descriptions
 * @return {TreeNode}
 */
var createBinaryTree = function (descriptions) {
  // Stores value -> TreeNode mapping
  const nodes = new Map();

  // Stores all child values
  const children = new Set();

  for (const [parent, child, isLeft] of descriptions) {
    // Create parent node if needed
    if (!nodes.has(parent)) nodes.set(parent, new TreeNode(parent));

    // Create child node if needed
    if (!nodes.has(child)) nodes.set(child, new TreeNode(child));

    // Connect child to parent
    if (isLeft === 1) nodes.get(parent).left = nodes.get(child);
    else nodes.get(parent).right = nodes.get(child);

    // Mark child
    children.add(child);
  }

  // Root never appears as a child
  for (const [value, node] of nodes) {
    if (!children.has(value)) return node;
  }

  return null;
};

var sumRootToLeaf = function (root) {
  function dfs(node, current) {
    if (!node) return 0;

    // Build binary value
    current = current * 2 + node.val;

    // If leaf
    if (!node.left && !node.right) {
      return current;
    }

    // Return sum of both sides
    return dfs(node.left, current) + dfs(node.right, current);
  }

  return dfs(root, 0);
};

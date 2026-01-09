var subtreeWithAllDeepest = function (root) {
  function dfs(node) {
    if (!node) return [0, null];

    const [ld, ln] = dfs(node.left);
    const [rd, rn] = dfs(node.right);

    if (ld === rd) {
      return [ld + 1, node];
    } else if (ld > rd) {
      return [ld + 1, ln];
    } else {
      return [rd + 1, rn];
    }
  }

  return dfs(root)[1];
};

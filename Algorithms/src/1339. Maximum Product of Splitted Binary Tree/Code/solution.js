var maxProduct = function (root) {
  const MOD = 1e9 + 7;
  let totalSum = 0;
  let maxProd = 0;

  function getTotalSum(node) {
    if (!node) return 0;
    return node.val + getTotalSum(node.left) + getTotalSum(node.right);
  }

  function dfs(node) {
    if (!node) return 0;

    let left = dfs(node.left);
    let right = dfs(node.right);

    let subtreeSum = node.val + left + right;
    maxProd = Math.max(maxProd, subtreeSum * (totalSum - subtreeSum));

    return subtreeSum;
  }

  totalSum = getTotalSum(root);
  dfs(root);

  return maxProd % MOD;
};

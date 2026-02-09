var balanceBST = function (root) {
  let arr = [];

  // Inorder traversal
  function inorder(node) {
    if (!node) return;
    inorder(node.left);
    arr.push(node.val);
    inorder(node.right);
  }

  // Build balanced BST
  function build(left, right) {
    if (left > right) return null;

    let mid = Math.floor((left + right) / 2);
    let node = new TreeNode(arr[mid]);

    node.left = build(left, mid - 1);
    node.right = build(mid + 1, right);

    return node;
  }

  inorder(root);
  return build(0, arr.length - 1);
};

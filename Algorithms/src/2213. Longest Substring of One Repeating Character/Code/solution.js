/**
 * @param {string} s
 * @param {string} queryCharacters
 * @param {number[]} queryIndices
 * @return {number[]}
 */
var longestRepeating = function (s, queryCharacters, queryIndices) {
  const n = s.length;

  // I use an array to represent the segment tree.
  const tree = new Array(4 * n);

  // Every node stores the same six pieces of information.
  const createNode = () => ({
    leftChar: "",
    rightChar: "",
    len: 0,
    prefix: 0,
    suffix: 0,
    best: 0,
  });

  // I merge two adjacent segments into one.
  const merge = (left, right) => {
    const res = createNode();

    // Set the boundaries and total length of the merged segment.
    res.leftChar = left.leftChar;
    res.rightChar = right.rightChar;
    res.len = left.len + right.len;

    // Initially, the best answer is inside one of the children.
    res.best = Math.max(left.best, right.best);

    // The prefix normally comes from the left child.
    res.prefix = left.prefix;

    // The complete left segment must be uniform before
    // its prefix can extend into the right segment.
    if (left.prefix === left.len && left.rightChar === right.leftChar) {
      res.prefix = left.len + right.prefix;
    }

    // The suffix normally comes from the right child.
    res.suffix = right.suffix;

    // The complete right segment must be uniform before
    // its suffix can extend into the left segment.
    if (right.suffix === right.len && left.rightChar === right.leftChar) {
      res.suffix = left.suffix + right.len;
    }

    // Check the substring that crosses the boundary.
    if (left.rightChar === right.leftChar) {
      res.best = Math.max(res.best, left.suffix + right.prefix);
    }

    return res;
  };

  // I build the tree from the original string.
  const build = (node, l, r) => {
    // A leaf represents one character.
    if (l === r) {
      tree[node] = {
        leftChar: s[l],
        rightChar: s[l],
        len: 1,
        prefix: 1,
        suffix: 1,
        best: 1,
      };
      return;
    }

    const mid = Math.floor((l + r) / 2);

    // Build both halves.
    build(node * 2, l, mid);
    build(node * 2 + 1, mid + 1, r);

    // Combine the two halves.
    tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
  };

  // I update only the path containing the changed index.
  const update = (node, l, r, idx, char) => {
    // At the target leaf, replace the old character.
    if (l === r) {
      tree[node] = {
        leftChar: char,
        rightChar: char,
        len: 1,
        prefix: 1,
        suffix: 1,
        best: 1,
      };
      return;
    }

    const mid = Math.floor((l + r) / 2);

    // Move into the correct child.
    if (idx <= mid) {
      update(node * 2, l, mid, idx, char);
    } else {
      update(node * 2 + 1, mid + 1, r, idx, char);
    }

    // Recalculate this node after the update.
    tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
  };

  // Build the tree once before processing the queries.
  build(1, 0, n - 1);

  const answer = [];

  for (let i = 0; i < queryCharacters.length; i++) {
    // Apply the current query's character update.
    update(1, 0, n - 1, queryIndices[i], queryCharacters[i]);

    // The root stores the answer for the whole string.
    answer.push(tree[1].best);
  }

  return answer;
};

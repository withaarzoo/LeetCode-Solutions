var maxLevelSum = function (root) {
  let queue = [root];
  let level = 1;
  let answerLevel = 1;
  let maxSum = -Infinity;

  while (queue.length > 0) {
    let size = queue.length;
    let levelSum = 0;

    for (let i = 0; i < size; i++) {
      let node = queue.shift();
      levelSum += node.val;

      if (node.left) queue.push(node.left);
      if (node.right) queue.push(node.right);
    }

    if (levelSum > maxSum) {
      maxSum = levelSum;
      answerLevel = level;
    }

    level++;
  }

  return answerLevel;
};

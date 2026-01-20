var minBitwiseArray = function (nums) {
  let ans = [];

  for (let p of nums) {
    let found = -1;

    for (let x = 0; x <= p; x++) {
      if ((x | (x + 1)) === p) {
        found = x;
        break;
      }
    }

    ans.push(found);
  }

  return ans;
};

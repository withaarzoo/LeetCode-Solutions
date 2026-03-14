var getHappyString = function (n, k) {
  let result = "";
  let count = 0;

  function dfs(curr) {
    if (result !== "") return;

    if (curr.length === n) {
      count++;
      if (count === k) result = curr;
      return;
    }

    for (let c of ["a", "b", "c"]) {
      if (curr.length > 0 && curr[curr.length - 1] === c) continue;

      dfs(curr + c);
    }
  }

  dfs("");

  return result;
};

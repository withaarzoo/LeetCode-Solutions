var maximizeSquareHoleArea = function (n, m, hBars, vBars) {
  const getMaxGap = (bars) => {
    bars.sort((a, b) => a - b);

    let maxLen = 1;
    let curLen = 1;

    for (let i = 1; i < bars.length; i++) {
      if (bars[i] === bars[i - 1] + 1) {
        curLen++;
      } else {
        curLen = 1;
      }
      maxLen = Math.max(maxLen, curLen);
    }
    return maxLen;
  };

  let hGap = getMaxGap(hBars) + 1;
  let vGap = getMaxGap(vBars) + 1;

  let side = Math.min(hGap, vGap);
  return side * side;
};

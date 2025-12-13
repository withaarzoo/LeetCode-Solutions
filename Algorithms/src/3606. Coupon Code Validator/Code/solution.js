var validateCoupons = function (code, businessLine, isActive) {
  const priority = {
    electronics: 0,
    grocery: 1,
    pharmacy: 2,
    restaurant: 3,
  };

  let valid = [];

  for (let i = 0; i < code.length; i++) {
    if (!isActive[i]) continue;
    if (!(businessLine[i] in priority)) continue;
    if (code[i].length === 0) continue;

    let ok = true;
    for (let ch of code[i]) {
      if (!/[a-zA-Z0-9_]/.test(ch)) {
        ok = false;
        break;
      }
    }
    if (!ok) continue;

    valid.push([priority[businessLine[i]], code[i]]);
  }

  // ✅ FIX: Use ASCII-based comparison instead of localeCompare
  valid.sort((a, b) => {
    if (a[0] === b[0]) {
      if (a[1] < b[1]) return -1;
      if (a[1] > b[1]) return 1;
      return 0;
    }
    return a[0] - b[0];
  });

  return valid.map((v) => v[1]);
};

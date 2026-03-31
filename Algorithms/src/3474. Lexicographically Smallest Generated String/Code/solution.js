/**
 * @param {string} str1
 * @param {string} str2
 * @return {string}
 */
var generateString = function (str1, str2) {
  const n = str1.length;
  const m = str2.length;
  const len = n + m - 1;

  const ans = Array(len).fill("?");
  const fixed = Array(len).fill(false);

  // Apply all 'T' constraints
  for (let i = 0; i < n; i++) {
    if (str1[i] === "T") {
      for (let j = 0; j < m; j++) {
        const pos = i + j;

        if (ans[pos] !== "?" && ans[pos] !== str2[j]) {
          return "";
        }

        ans[pos] = str2[j];
        fixed[pos] = true;
      }
    }
  }

  // Fill remaining positions with 'a'
  for (let i = 0; i < len; i++) {
    if (ans[i] === "?") ans[i] = "a";
  }

  // Process all 'F' constraints
  for (let i = 0; i < n; i++) {
    if (str1[i] === "F") {
      let same = true;

      for (let j = 0; j < m; j++) {
        if (ans[i + j] !== str2[j]) {
          same = false;
          break;
        }
      }

      if (!same) continue;

      let changed = false;

      for (let j = m - 1; j >= 0; j--) {
        const pos = i + j;

        if (fixed[pos]) continue;

        for (let c = 97; c <= 122; c++) {
          const ch = String.fromCharCode(c);

          if (ch !== ans[pos] && ch !== str2[j]) {
            ans[pos] = ch;
            changed = true;
            break;
          }
        }

        if (changed) break;
      }

      if (!changed) return "";
    }
  }

  return ans.join("");
};

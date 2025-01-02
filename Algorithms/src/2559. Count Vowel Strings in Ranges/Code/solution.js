var vowelStrings = function (words, queries) {
  const vowels = new Set(["a", "e", "i", "o", "u"]);
  const n = words.length;
  const prefix = Array(n).fill(0);

  // Precompute the prefix sum
  for (let i = 0; i < n; i++) {
    if (vowels.has(words[i][0]) && vowels.has(words[i][words[i].length - 1])) {
      prefix[i] = 1;
    }
    if (i > 0) {
      prefix[i] += prefix[i - 1];
    }
  }

  // Answer the queries
  return queries.map(([l, r]) => prefix[r] - (l > 0 ? prefix[l - 1] : 0));
};

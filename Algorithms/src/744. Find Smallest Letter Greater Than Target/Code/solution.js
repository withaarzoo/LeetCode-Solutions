var nextGreatestLetter = function (letters, target) {
  let left = 0;
  let right = letters.length - 1;
  let answer = letters[0]; // wrap-around case

  while (left <= right) {
    let mid = Math.floor(left + (right - left) / 2);

    if (letters[mid] > target) {
      answer = letters[mid];
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }

  return answer;
};

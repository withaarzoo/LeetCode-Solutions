var minimizeXor = function (num1, num2) {
  let count2 = num2.toString(2).split("1").length - 1; // Count of 1s in num2
  let count1 = num1.toString(2).split("1").length - 1; // Count of 1s in num1

  if (count1 === count2) return num1;

  let result = num1;
  if (count1 > count2) {
    for (let i = 0; i < 32 && count1 > count2; i++) {
      if ((result & (1 << i)) !== 0) {
        result &= ~(1 << i); // Clear bit
        count1--;
      }
    }
  } else {
    for (let i = 0; i < 32 && count1 < count2; i++) {
      if ((result & (1 << i)) === 0) {
        result |= 1 << i; // Set bit
        count1++;
      }
    }
  }
  return result;
};

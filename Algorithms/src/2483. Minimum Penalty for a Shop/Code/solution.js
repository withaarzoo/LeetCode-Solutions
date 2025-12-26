/**
 * @param {string} customers
 * @return {number}
 */
var bestClosingTime = function (customers) {
  let totalY = 0;
  for (let c of customers) {
    if (c === "Y") totalY++;
  }

  let openPenalty = 0;
  let closedPenalty = totalY;
  let minPenalty = closedPenalty;
  let answer = 0;

  for (let i = 0; i < customers.length; i++) {
    if (customers[i] === "N") {
      openPenalty++;
    } else {
      closedPenalty--;
    }

    let currentPenalty = openPenalty + closedPenalty;
    if (currentPenalty < minPenalty) {
      minPenalty = currentPenalty;
      answer = i + 1;
    }
  }

  return answer;
};

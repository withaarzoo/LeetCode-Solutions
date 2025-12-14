var numberOfWays = function (corridor) {
  const MOD = 1_000_000_007;
  const seats = [];

  for (let i = 0; i < corridor.length; i++) {
    if (corridor[i] === "S") seats.push(i);
  }

  if (seats.length === 0 || seats.length % 2 !== 0) return 0;

  let ways = 1;
  for (let i = 2; i < seats.length; i += 2) {
    ways = (ways * (seats[i] - seats[i - 1])) % MOD;
  }

  return ways;
};

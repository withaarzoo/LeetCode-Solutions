const primeSubOperation = function (nums) {
  const sieve = new Array(1001).fill(true);
  const primes = [];

  for (let i = 2; i <= 1000; i++) {
    if (sieve[i]) {
      primes.push(i);
      for (let j = i * 2; j <= 1000; j += i) sieve[j] = false;
    }
  }

  for (let i = nums.length - 2; i >= 0; i--) {
    if (nums[i] < nums[i + 1]) continue;

    let target = nums[i] - nums[i + 1];
    let prime = primes.find((p) => p > target);
    if (!prime) return false;

    nums[i] -= prime;
    if (nums[i] <= 0) return false;
  }
  return true;
};

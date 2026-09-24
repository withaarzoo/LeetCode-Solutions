function smallestIndex(nums: number[]): number {
    for (let i = 0; i < nums.length; ++i) {
        let x = nums[i];
        let sum = 0;
        while (x > 0) {
            sum += x % 10;
            x = Math.floor(x / 10);
        }
        if (sum === i) return i;
    }
    return -1;
};
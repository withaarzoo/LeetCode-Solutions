function targetIndices(nums: number[], target: number): number[] {
    let less = 0, equal = 0;
    for (const x of nums) {
        if (x < target) less++;
        else if (x === target) equal++;
    }
    const ans: number[] = [];
    for (let i = 0; i < equal; i++) {
        ans.push(less + i);
    }
    return ans;
};
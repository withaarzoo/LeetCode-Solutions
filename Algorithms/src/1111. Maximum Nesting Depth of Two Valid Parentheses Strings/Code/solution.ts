function maxDepthAfterSplit(seq: string): number[] {
    const ans: number[] = new Array(seq.length);
    let depth = 0;
    for (let i = 0; i < seq.length; ++i) {
        if (seq[i] === '(') {
            ++depth;
            ans[i] = depth % 2;
        } else {
            ans[i] = depth % 2;
            --depth;
        }
    }
    return ans;
};
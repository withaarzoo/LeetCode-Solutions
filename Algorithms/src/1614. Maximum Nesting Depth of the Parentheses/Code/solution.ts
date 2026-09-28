function maxDepth(s: string): number {
    let depth = 0, maxDepth = 0;
    for (let c of s) {
        if (c === '(') {
            depth++;
            if (depth > maxDepth) maxDepth = depth;
        } else if (c === ')') {
            depth--;
        }
    }
    return maxDepth;
};
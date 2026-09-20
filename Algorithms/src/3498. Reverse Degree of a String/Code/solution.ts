function reverseDegree(s: string): number {
    let ans = 0;
    for (let i = 0; i < s.length; ++i) {
        ans += (26 - (s.charCodeAt(i) - 97)) * (i + 1);
    }
    return ans;
};
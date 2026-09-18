function maxNumOfSubstrings(s: string): string[] {
    const n = s.length;
    // first and last occurrence of each letter
    const left: number[] = new Array(26).fill(-1);
    const right: number[] = new Array(26).fill(-1);
    for (let i = 0; i < n; i++) {
        const c = s.charCodeAt(i) - 97;
        if (left[c] === -1) left[c] = i;   // record first time we see it
        right[c] = i;                      // always update last time
    }

    // collect every valid closed interval
    const intervals: [number, number][] = [];
    for (let i = 0; i < 26; i++) {
        if (left[i] === -1) continue;      // letter never appeared
        let L = left[i], R = right[i];
        let valid = true;
        // expand right end while scanning the current range
        for (let j = L; j <= R; j++) {
            const c = s.charCodeAt(j) - 97;
            if (left[c] < L) {             // needs to start earlier -> not closed
                valid = false;
                break;
            }
            R = Math.max(R, right[c]);     // push right end if needed
        }
        if (valid) intervals.push([L, R]);
    }

    // sort by ending position so greedy can pick earliest-ending first
    intervals.sort((a, b) => a[1] - b[1]);

    // greedy selection of non-overlapping intervals
    const ans: string[] = [];
    let lastEnd = -1;
    for (const [L, R] of intervals) {
        if (L > lastEnd) {                 // completely after previous piece
            ans.push(s.substring(L, R + 1));
            lastEnd = R;
        }
    }
    return ans;
}
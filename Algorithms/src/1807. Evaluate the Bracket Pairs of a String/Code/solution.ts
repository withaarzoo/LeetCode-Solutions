function evaluate(s: string, knowledge: string[][]): string {
    const mp = new Map<string, string>();
    for (const [k, v] of knowledge) mp.set(k, v);
    let res = "";
    const n = s.length;
    for (let i = 0; i < n; ) {
        if (s[i] === '(') {
            let j = i + 1;
            while (s[j] !== ')') ++j;
            const key = s.substring(i + 1, j);
            res += mp.has(key) ? mp.get(key)! : "?";
            i = j + 1;
        } else {
            res += s[i++];
        }
    }
    return res;
};
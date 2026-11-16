function digitSum(s: string, k: number): string {
    while (s.length > k) {
        let next = "";
        for (let i = 0; i < s.length; i += k) {
            let sum = 0;
            for (let j = i; j < s.length && j < i + k; j++) {
                sum += Number(s[j]);
            }
            next += sum.toString();
        }
        s = next;
    }
    return s;
};
/**
 * @param {string} expression
 * @return {string[]}
 */
var braceExpansionII = function(expression) {
    let i = 0;
    const parse = () => {
        const res = new Set();
        let cur = new Set([""]);
        while (i < expression.length && expression[i] !== '}') {
            if (expression[i] === '{') {
                i++;
                const next = parse();
                i++;
                cur = product(cur, next);
            } else if (expression[i] === ',') {
                for (const s of cur) res.add(s);
                cur = new Set([""]);
                i++;
            } else {
                const next = new Set([expression[i]]);
                i++;
                cur = product(cur, next);
            }
        }
        for (const s of cur) res.add(s);
        return res;
    };
    const product = (a, b) => {
        const res = new Set();
        for (const x of a)
            for (const y of b)
                res.add(x + y);
        return res;
    };
    const result = parse();
    return Array.from(result).sort();
};
function braceExpansionII(expression: string): string[] {
    let pos = 0;

    function product(a: Set<string>, b: Set<string>): Set<string> {
        const res = new Set<string>();

        for (const x of a) {
            for (const y of b) {
                res.add(x + y);
            }
        }

        return res;
    }

    function parseFactor(): Set<string> {
        if (expression[pos] === "{") {
            pos++;
            const res = parseExpression();
            pos++;
            return res;
        }

        const res = new Set<string>();
        res.add(expression[pos++]);
        return res;
    }

    function parseTerm(): Set<string> {
        let res = new Set<string>([""]);

        while (
            pos < expression.length &&
            expression[pos] !== "}" &&
            expression[pos] !== ","
        ) {
            res = product(res, parseFactor());
        }

        return res;
    }

    function parseExpression(): Set<string> {
        const res = new Set<string>();

        while (
            pos < expression.length &&
            expression[pos] !== "}"
        ) {
            const term = parseTerm();

            for (const word of term) {
                res.add(word);
            }

            if (
                pos < expression.length &&
                expression[pos] === ","
            ) {
                pos++;
            }
        }

        return res;
    }

    return Array.from(parseExpression()).sort();
}
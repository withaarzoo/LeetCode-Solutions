/**

 * @param {string} s

 * @return {boolean}

 */

var isValid = function(s) {

    const st = [];

    for (const c of s) {

        if (c === '(' || c === '{' || c === '[') {

            st.push(c);

        } else {

            if (st.length === 0) return false;

            const top = st.pop();

            if ((c === ')' && top !== '(') || (c === '}' && top !== '{') || (c === ']' && top !== '[')) {

                return false;

            }

        }

    }

    return st.length === 0;

};
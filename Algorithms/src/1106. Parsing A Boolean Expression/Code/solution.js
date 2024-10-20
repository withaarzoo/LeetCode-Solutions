/**
 * @param {string} expression
 * @return {boolean}
 */
var parseBoolExpr = function (expression) {
  let stack = [];

  for (let c of expression) {
    if (c === ")") {
      let subExpr = [];
      while (stack[stack.length - 1] !== "(") {
        subExpr.push(stack.pop());
      }
      stack.pop(); // Remove '('

      let op = stack.pop(); // Get the operator

      if (op === "!") {
        stack.push(subExpr[0] === "t" ? "f" : "t");
      } else if (op === "&") {
        let result = "t";
        for (let e of subExpr) {
          if (e === "f") {
            result = "f";
            break;
          }
        }
        stack.push(result);
      } else if (op === "|") {
        let result = "f";
        for (let e of subExpr) {
          if (e === "t") {
            result = "t";
            break;
          }
        }
        stack.push(result);
      }
    } else if (c !== ",") {
      stack.push(c);
    }
  }

  return stack[0] === "t";
};

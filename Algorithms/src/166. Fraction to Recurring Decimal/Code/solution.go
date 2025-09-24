package main

import (
    "strconv"
    "strings"
)

func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }

    var sb strings.Builder
    // sign
    if (numerator < 0) != (denominator < 0) {
        sb.WriteByte('-')
    }

    // convert to int64 and take absolute values to be safe
    n := int64(numerator)
    d := int64(denominator)
    if n < 0 {
        n = -n
    }
    if d < 0 {
        d = -d
    }

    // integer part
    sb.WriteString(strconv.FormatInt(n/d, 10))
    rem := n % d
    if rem == 0 {
        return sb.String()
    }

    sb.WriteByte('.')
    posMap := make(map[int64]int) // remainder -> index in current sb string

    for rem != 0 {
        if p, ok := posMap[rem]; ok {
            s := sb.String()
            // insert '(' at p and ')' at end
            return s[:p] + "(" + s[p:] + ")"
        }
        posMap[rem] = sb.Len()
        rem *= 10
        sb.WriteString(strconv.FormatInt(rem/d, 10))
        rem = rem % d
    }
    return sb.String()
}

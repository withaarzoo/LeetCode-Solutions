func totalMoney(n int) int {
    w := n / 7 // full weeks
    r := n % 7 // remaining days
    // full weeks sum = w*28 + 7 * (0 + 1 + ... + (w-1))
    fullWeeksSum := w*28 + 7*(w*(w-1)/2)
    // remaining days sum = r*(1 + w) + r*(r-1)/2
    remSum := r*(1+w) + (r*(r-1)/2)
    return fullWeeksSum + remSum
}

func numWaterBottles(numBottles int, numExchange int) int {
    total := numBottles
    empties := numBottles

    for empties >= numExchange {
        newFull := empties / numExchange
        total += newFull
        empties = newFull + empties%numExchange
    }

    return total
}

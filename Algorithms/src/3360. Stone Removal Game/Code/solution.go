func canAliceWin(n int) bool {
    remove := 10
    aliceTurn := true
    for n >= remove {
        n -= remove
        remove--
        aliceTurn = !aliceTurn
    }
    return !aliceTurn
}
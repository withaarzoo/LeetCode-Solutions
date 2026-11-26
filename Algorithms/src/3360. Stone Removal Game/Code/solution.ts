function canAliceWin(n: number): boolean {
    let remove = 10;
    let aliceTurn = true;
    while (n >= remove) {
        n -= remove;
        remove--;
        aliceTurn = !aliceTurn;
    }
    return !aliceTurn;
};
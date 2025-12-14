func numberOfWays(corridor string) int {
    const mod = 1000000007
    seats := []int{}

    for i := 0; i < len(corridor); i++ {
        if corridor[i] == 'S' {
            seats = append(seats, i)
        }
    }

    if len(seats) == 0 || len(seats)%2 != 0 {
        return 0
    }

    ways := 1
    for i := 2; i < len(seats); i += 2 {
        ways = (ways * (seats[i] - seats[i-1])) % mod
    }

    return ways
}

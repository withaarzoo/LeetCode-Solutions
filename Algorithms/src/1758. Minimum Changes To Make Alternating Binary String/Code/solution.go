func minOperations(s string) int {

    startWith0 := 0
    startWith1 := 0

    for i := 0; i < len(s); i++ {

        var expected0 byte
        var expected1 byte

        if i%2 == 0 {
            expected0 = '0'
            expected1 = '1'
        } else {
            expected0 = '1'
            expected1 = '0'
        }

        if s[i] != expected0 {
            startWith0++
        }

        if s[i] != expected1 {
            startWith1++
        }
    }

    if startWith0 < startWith1 {
        return startWith0
    }
    return startWith1
}
func getHappyString(n int, k int) string {

    chars := []byte{'a','b','c'}
    count := 0
    result := ""

    var dfs func(curr string)

    dfs = func(curr string){

        if result != "" {
            return
        }

        if len(curr) == n {
            count++
            if count == k {
                result = curr
            }
            return
        }

        for _,c := range chars {

            if len(curr) > 0 && curr[len(curr)-1] == c {
                continue
            }

            dfs(curr + string(c))
        }
    }

    dfs("")
    return result
}
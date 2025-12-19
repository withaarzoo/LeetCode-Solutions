func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][2] < meetings[j][2]
    })

    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }

    knows := make([]bool, n)
    knows[0], knows[firstPerson] = true, true

    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }

    union := func(x, y int) {
        x, y = find(x), find(y)
        if x != y {
            parent[y] = x
        }
    }

    i := 0
    for i < len(meetings) {
        time := meetings[i][2]
        people := []int{}

        j := i
        for j < len(meetings) && meetings[j][2] == time {
            union(meetings[j][0], meetings[j][1])
            people = append(people, meetings[j][0], meetings[j][1])
            j++
        }

        good := map[int]bool{}
        for _, p := range people {
            if knows[p] {
                good[find(p)] = true
            }
        }

        for _, p := range people {
            if good[find(p)] {
                knows[p] = true
            } else {
                parent[p] = p
            }
        }
        i = j
    }

    res := []int{}
    for i := 0; i < n; i++ {
        if knows[i] {
            res = append(res, i)
        }
    }
    return res
}

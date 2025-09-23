package main

func compareVersion(version1 string, version2 string) int {
    i, j := 0, 0
    n, m := len(version1), len(version2)

    for i < n || j < m {
        var num1, num2 int
        // parse next revision from version1
        for i < n && version1[i] != '.' {
            num1 = num1*10 + int(version1[i]-'0')
            i++
        }
        if i < n && version1[i] == '.' {
            i++
        }

        // parse next revision from version2
        for j < m && version2[j] != '.' {
            num2 = num2*10 + int(version2[j]-'0')
            j++
        }
        if j < m && version2[j] == '.' {
            j++
        }

        if num1 < num2 {
            return -1
        }
        if num1 > num2 {
            return 1
        }
    }
    return 0
}

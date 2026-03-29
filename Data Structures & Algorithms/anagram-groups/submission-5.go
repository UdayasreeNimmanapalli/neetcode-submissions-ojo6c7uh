func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)

    for _, str := range strs {
        runes := []rune(str)

        sort.Slice(runes, func(i, j int) bool {
            return runes[i] < runes[j]
        })

        sortedStr := string(runes)

        array, ok := groups[sortedStr]

        if ok {
            groups[sortedStr] = append(array, str)
        } else {
            groups[sortedStr] = []string{str}
        }
    }

    values := make([][]string,0,len(groups))

    for _,val := range groups{
        values = append(values, val)
    }

    return values
}

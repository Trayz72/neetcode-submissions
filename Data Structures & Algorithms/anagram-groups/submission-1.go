func groupAnagrams(strs []string) [][]string {
    groups := map[[26]int][]string{}

    for _, s := range strs{
        var counts [26]int
        for _, c := range s {
            counts[c-'a']++
        }
        groups[counts] = append(groups[counts],s)
    }
    result := [][]string{}
    for _,group := range groups{
        result = append(result,group)
    }
    return result
}

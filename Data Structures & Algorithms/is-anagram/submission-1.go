func isAnagram(s string, t string) bool {

    // First Ap.

    // if len(s) != len(t){
    //     return false
    // }

    // counts := map[rune]int{}
    // for _,val := range s{
    //     counts[val]++
    // }
    // for _,val := range t{
    //     counts[val]--

    //     // for the marginal improvement we can use:

    //     // if counts[val] < 0{
    //     //     return false
    //     // }
    // }

    // // more readable 
    // // for _,val := range counts{
    // //     if val != 0 {
    // //         return false
    // //     }
    // // }
    // return true



    // Second Ap.

    if len(s) != len(t){
        return false
    }

    var counts [26]int

    for i := 0 ; i<len(s) ; i++ {
        counts[s[i]-'a']++
        counts[t[i]-'a']--
    }
    for _,v := range counts {
        if v != 0{
            return false
        }
    }
    return true
}

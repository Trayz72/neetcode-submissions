func topKFrequent(nums []int, k int) []int {
    count := map[int]int{}
    for _,n := range nums{
        count[n]++
    }

    bucket := make([][]int, len(nums)+1)
    for n, fq := range count{
        bucket[fq] = append(bucket[fq], n)
    }

    result := make([]int,0,k)
    for i:=len(bucket)-1 ; i>=0 && len(result)<k; i-- {
        for _,n := range bucket[i]{
            if len(result) == k {
                break
            }
            result = append(result,n)
        }
       
    }
    return result
}

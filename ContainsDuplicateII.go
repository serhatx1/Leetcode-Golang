func containsNearbyDuplicate(nums []int, k int) bool {
    mp := make(map[int]int)
    for i, num := range nums {
        a,b:=mp[num]
        if b && i-a <= k {
            return true
        }
        mp[num] = i
    }
    return false
}

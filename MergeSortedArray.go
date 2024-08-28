func merge(nums1 []int, m int, nums2 []int, n int){
    k:=0
   for i:=0;i<len(nums1);i++{
    if nums1[i]==0 && k<n{
        nums1[i]=nums2[k]
        k++
    }
   }
    slices.Sort(nums1)
}
// https://leetcode.com/problems/merge-sorted-array/solutions/5703995/topic/
// Time: Beat 100% - Space: Beat %76.7

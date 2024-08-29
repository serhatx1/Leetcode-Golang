func sortColors(nums []int)  {
    l:=0
    h:=len(nums)-1
    i:=0
    for i<=h{
        if nums[i]==0{
            nums[i],nums[l]=nums[l],nums[i]
            l+=1
            i++
        }else if nums[i]==2{
            nums[i],nums[h]=nums[h],nums[i]
            h-=1
        }else{
            i++
        }
       }
}
//Runtime Beats 100.00% - Memory Beats 88.23%
// https://leetcode.com/problems/sort-colors/solutions/5706650/topic/

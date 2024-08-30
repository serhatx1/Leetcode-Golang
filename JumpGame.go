func canJump(nums []int) bool {
    k:=1
    for i:=len(nums)-2;i>=0;i--{
        if nums[i]==0{
            k+=1
        }else if nums[i]>=k{
            k=1
        }else{
            k+=1
        }
    }
    if k==1{
        return true
    }
    return false
}
// https://leetcode.com/problems/jump-game/submissions/
// Time = Beats 98% Memory = Beats %97.89

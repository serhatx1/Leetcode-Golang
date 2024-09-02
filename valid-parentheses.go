func isValid(s string) bool {
    arr:=make([]string,0)
    hashmp:=map[string]string{
        ")": "(",
        "}": "{",
        "]": "[",
    }
    for i:=0;i<len(s);i++{
        val:=string(s[i])
         _,e:=hashmp[val]
        if !e{ 
            arr=append(arr,val)
        }else if len(arr)>0 && arr[len(arr)-1]==hashmp[val] {
            arr=arr[:len(arr)-1]
        }else{
            return false
        }
    }
    if len(arr)==0{
        return true
    }
    return false
}

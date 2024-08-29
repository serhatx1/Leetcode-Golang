func plusOne(digits []int) []int {
    k:=1;
    for i:=len(digits)-1;i>=0;i--{
       if digits[i]+k<10{
        digits[i]+=k
        k=0
       }else if i!=0{
        digits[i]=0
        k=1
       }else if i==0{
        digits[i]=k
        digits=append(digits,0)
       }
       
    }
    return digits
}






func plusOne(digits []int) []int {
    k:=1;
    for i:=len(digits)-1;i>=0;i--{
        if i==len(digits)-1 && digits[i]<9{
            digits[i]=digits[i]+k
            return digits
            }  else if digits[i]+k>=10 && i!=0{
                digits[i]=0
                k=1
            } else if digits[i]+k==10&&i==0{
            digits[i]=1
            digits=append(digits,0)
            fmt.Println(digits)
        }else{
            digits[i]+=k
            k=0
        }
      
    }
    return digits
}

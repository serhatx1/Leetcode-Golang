func countPrimes(n int) int {
    
    count:=0
	slce := make([]bool, n+1)
    
    for i:=2;i<n;i++{
        slce[i]=true
    }
    for i:=2;i*i<n;i++{
        if slce[i]==true{
        for k:=i*i;k<=n;k+=i{
        slce[k]=false
       }
        }
       
    }
    for i:=2;i<n;i++{
        if slce[i]==true{
            count++
        }
    }
    return count
}
// Use Sieve of Eratosthenes for this question..

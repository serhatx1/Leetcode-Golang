func maxProfit(prices []int) int {
    buy:=prices[0]
    profit:=0
    for i:=1;i<len(prices);i++{
        if buy>prices[i]{
            buy=prices[i]
        }else if profit<prices[i]-buy{
            profit=prices[i]-buy
        }
    }
    return profit                           
}
//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/5714956/topic/

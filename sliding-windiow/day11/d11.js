const buyStock = (prices) => {
    let left = 0, right = 1, maxProfit = 0;
    while (right < prices.length) {
        let currentDifference = prices[right] - prices[left]
        if (currentDifference > 0) {
            maxProfit = Math.max(maxProfit, currentDifference)
            right++
        }
        else {
            left = right
            right++
        }
    }
    return maxProfit
}


const prices = [7, 1, 5, 3, 6, 4]
console.log(buyStock(prices))
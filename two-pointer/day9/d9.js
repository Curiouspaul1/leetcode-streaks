const maxArea = (height) =>{
    // create a left and right pointer  to 0 and height.length
    // maxArea = 0
    // while left < right
    // get the minimum between height[left] and height[right], square it and set ut as the currentArea
    // compare the maximum betweeek maxArea and currentArea, set the maxArea to the result
    // if curArea < maxArea , move left by 1
    // else move right by 1
    let left = 0, right = height.length - 1
    let maxArea = 0

    while (left < right){
        let currenMin = Math.min(height[right], height[left])
        let length = right - left
        console.log(currenMin);
        let currentArea = currenMin*length
        

        maxArea = Math.max(currentArea, maxArea)
        if (currentArea < maxArea){
            left++
        } else{
            right++
        }
    }
    return maxArea
}



const height = [1,8,6,2,5,4,8,3,7]

console.log(maxArea(height));

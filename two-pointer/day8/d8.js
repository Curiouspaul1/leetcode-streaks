const threeSum = (nums) => {
    const result = [];
    nums.sort((a, b) => a - b);

    for (let i = 0; i < nums.length; i++) {
        let a = nums[i];
        if (a > 0) break
        if (i > 0 && a === nums[i - 1]) continue
        let left = i + 1; right = nums.length - 1;

        while (left < right) {
            let threeSum = a + nums[left] + nums[right];
            if (threeSum > 0) {
                right--;
            }
            else if (threeSum < 0) {
                left++;
            }
            else {
                result.push([a, nums[left], nums[right]]);
                left++;
                while (nums[left] === nums[left - 1] && left < right) {
                    left++;
                }
            }
        }
    }
    return result;
};


const nums = [-1, 0, 1, 2, -1, -4];
console.log(threeSum(nums));
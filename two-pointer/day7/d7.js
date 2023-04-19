/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (numbers, target) {
    let r = numbers.length - 1, l = 0
    let result = []

    while (l < r) {
        const element1 = numbers[l];
        const element2 = numbers[r]
        let sum = element1 + element2

        if (sum > target) {
            r -= 1
        }
        else if (sum < target) {
            l += 1
        }
        else {
            return [1 + l, 1 + r]
        }
    }
};

console.log(twoSum([2, 7, 11, 15], 9))
console.log(twoSum([2, 3, 4], 6))
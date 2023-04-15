// create an empty result object
// iterate through the nums array
// count the number of times each number appears in the array
// add the number and its count to the result map
// sort the result map by the count
// return the first k key in the result map

const topKFrequent = (nums, k) => {
    const result = {};
    nums.forEach(num => {
        result[num] = result[num] + 1 || 1;
    }
    );
    const final = Object.entries(result).sort((a, b) => b[1] - a[1]).slice(0, k);
    return final.map(num => num[0]);
}

console.log(topKFrequent([1, 1, 1, 2, 2, 2, 2, 3], 2))
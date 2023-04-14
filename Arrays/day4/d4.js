const groupAnagrams = (strs) => {
    const map = new Map();
    for (const str of strs) {
        const key = [...str].sort().join()
        if (map.has(key)) {
            map.get(key).push(str)
        }
        else {
            map.set(key, [str])
        }
    }
    return [...map.values()]
}

const strs = ["eat","tea","tan","ate","nat","bat"]

console.log(groupAnagrams(strs))
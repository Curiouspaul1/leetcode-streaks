const lengthOfLongestSubstring = (s) => {
    var charSet = new Set(), l = 0, res = 0

    for (let r = 0; r < s.length; r++) {
        while (charSet.has(s[r])) {
            charSet.delete(s[l])
            l++
        }

        charSet.add(s[r])
        console.log( l, r, s[l], s[r], charSet.values(), charSet.size)
        res = Math.max(res, charSet.size)
        console.log(res)
    }
    return res

}

const s = "abcabcbb"
console.log(lengthOfLongestSubstring(s))
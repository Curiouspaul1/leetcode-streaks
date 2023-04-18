const assert = require("assert")
const isPalindrome = (str) => {
    // convert all character in str to lowercase
    //remove all spaces and punctuation from str
    // check if str is a palindrom using two pointer algorithm
    const regex = new RegExp(/[^\w\s]|_/g)
    const newString = str.toLowerCase().replace(regex, "").replace(/ +/g, "").trim()
    let r = newString.length - 1
    for (let i = 0; i <= r; i++) {
        if (newString[i] === newString[r]) { }
        else { return false }
        r -= 1
        console.log(i, r)
    }
    return true

}

console.log(isPalindrome("a man"));
const validParenthesis = s => {
    const map = {
        "(": ")",
        "[": "]",
        "{": "}"
    }
    let stack = []
    for (const i of s) {
        if (i !== Object.values(map)) {
            stack.push()
        }
        if ( stack[-1] !== map[i]) {
            return false
        }
        stack.pop()
    }
    return !stack
}


console.log(validParenthesis("()"))
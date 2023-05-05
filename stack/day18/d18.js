const generatePar = (n)=>{
    // create a empty result array to store ]
    // Create a while loop to add parenthesis
    // To add any parenthesis
    // check if # of open parenthesis is less than n, then add opem par else n
    // check if # of close   <  open parenthesis, then adding close par
    let stack = [], result = []
    const Backtrack = (open, close)=>{
        console.log(open, close, stack)
        if ((open === n) && (close === n)){
            const formatedStack = stack.toString().replaceAll(",", "")
            result.push(formatedStack)
            console.log("result called", result)
            return
        }
        if (open < n){
            stack.push("(")
            Backtrack(open + 1, close)
            stack.pop()
        }
        if (close < open){
            stack.push(")")
            Backtrack(open, close + 1)
            stack.pop()

        }
    }
    Backtrack(0,0)
    return result
    
}

console.log(generatePar(3))
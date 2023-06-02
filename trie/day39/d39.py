from typing import List

def longestCommonPrefix( strs: List[str]) -> str:
    short = min(strs, key=len)
    n =  len(short)

    ans = ""
    
    while n > 0:
        





strs = ["flower","flow","flight"]
print(longestCommonPrefix(strs))

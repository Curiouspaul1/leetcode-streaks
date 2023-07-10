def reversal(str):
    return str[-1] + reversal(str[:-1]) if str else ""

str = "Hello World!"
print(reversal(str))
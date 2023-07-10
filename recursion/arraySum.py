def Sum(arr):
    return arr[0] + Sum(arr[1:]) if arr else 0


arr = [i for i in range(6)]
print(Sum(arr))
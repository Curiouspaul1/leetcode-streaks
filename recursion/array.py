def flatten(arr):
    # Flatten a nested array
    ans = []
    if arr == []:
        return ans
    for i in arr:
        if isinstance(i, list):
            ans.extend(flatten(i))
        else:
            ans.append(i)

    return ans
array = [
    1,
    2,
    3,
    [4, 5, 6],
    7,
    [8, [9, 10, 11, [12, 13, 14]]],
    [15, 16, 17, 18, 19, [20, 21, 22, [23, 24, 25, [26, 27, 29]], 30, 31], 32],
    33,
]


print(flatten(array))
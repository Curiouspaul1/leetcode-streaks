def count_x(str):
    if str[0] == "x":
        return 1 + count_x(str[1 : len(str) - 1])
    else:
        return count_x(str[1 : len(str) - 1])


str = "axbxdx"
print(count_x(str))

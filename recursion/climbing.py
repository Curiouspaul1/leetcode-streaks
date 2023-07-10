def climbing(n):
    if n <= 0:
        return 0
    if n == 1:
        return 1
    if n == 2:
        return 2
    if n == 3:
        return 4
    return climbing(n - 1) + climbing(n - 2) + climbing(n - 3)


print(climbing(4))
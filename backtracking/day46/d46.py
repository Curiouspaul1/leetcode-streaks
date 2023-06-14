from itertools import chain, combinations
def subsets(nums):
    s = list(nums)
    final =  chain.from_iterable(combinations(s, r) for r in range(len(s) + 1))
    return list(final)
nums = [i for i in range(1,4)]
print(subsets(nums))

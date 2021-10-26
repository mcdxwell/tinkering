# David McDowell
# Tinkering with Python: Factorials, Permuations, and Combinations 
# Unneccessary "One-liners"


np = 'not possible'

factorial = lambda n: np if n < 0 else n * factorial(n - 1) if n > 1 else 1
permutation = (lambda n, r: np if n < 0 or r < 0 else 0 if n < r else 1 if n == 0 and r == 0 else ((n * factorial(n - 1)) / ((n - (r - 1)) * factorial(n - (r - 1) - 1)) if n == r else ((n * factorial(n - 1)) / ((n - r) * factorial(n - r - 1)))))
combination = lambda n, r: permutation(n, r) * 1 / factorial(r) if n > 0 and r > 0 else 1 if n == 0 and r == 0 else np

# Name functions
permutation.__name__ = 'Permutation'
combination.__name__ = 'Combination'

# Test cases
n = [7, 6, 7, 0, 2, 1, -4, 6, -5]
r = [3, 7, 7, 0, 3, 1, -5, -6, 5]

# This does not work for the factorial function
def printer(func):
    for i in range(len(n)):
        print(f'{func.__name__} ({n[i]}, {r[i]}) = {func(n[i], r[i])}')


printer(permutation)
printer(combination)

# autoPEP8 format for the one-line lambda functions above
# This is a little bit more readable

factorial = lambda n: np if n < 0 else n * factorial(n - 1) if n > 1 else 1
permutation = (
    lambda n, r: np
    if n < 0 or r < 0
    else 0
    if n < r
    else 1
    if n == 0 and r == 0
    else (
        (n * factorial(n - 1)) / ((n - (r - 1)) * factorial(n - (r - 1) - 1))
        if n == r
        else (n * factorial(n - 1)) / ((n - r) * factorial(n - r - 1))
    )
)
combination = (
    lambda n, r: permutation(n, r) * 1 / factorial(r)
    if n > 0 and r > 0
    else 1
    if n == 0 and r == 0
    else np
)
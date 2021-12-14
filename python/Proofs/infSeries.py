#  Tweet: https://twitter.com/infseriesbot/status/1466309773155254273




# Sum (k=0 => inf) (binom(a,k)*(binom(b,k)*(binom(c,k)) / (binom(a+b+c, k))

# equals

# (a+b)!(b+c)!(c+a)! / a!b!c!(a+b+c)!

# such that a,b,c are elements of all natural numbers including 0


a = 20
b = 45
c = 2

x = a + c
y = b + c
z = c + a

d = a + b + c

k = 3



def factorial(number: int) -> int:

    product = 1
    while(number > 1):
        product *= number
        number -= 1
    return product


def binomial(n: int, k: int) -> int:


    return factorial(n) / (factorial(k) * factorial(n-k))



print(binomial(3, 1)*binomial(4,1)*binomial(5,1))



leftSide = (binomial(a, k) * binomial(b, k) * binomial(c, k)) / (binomial(d, k))


rightSide = (factorial(a+b) * factorial(b+c) * factorial(c+a)) / (factorial(a) * factorial(b) * factorial(c) * factorial(d))


print(leftSide)
print(rightSide)
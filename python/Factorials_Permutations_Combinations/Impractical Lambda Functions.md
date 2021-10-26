# Impractical Lambda Functions
This script was inspired by Joe Norman's [tweet/thread](https://twitter.com/normonics/status/1450474773411151878?s=20) and the idea of reusing a factorial 
function for combination and permutation functions.
Tweet/thread: 

Tweet 1:
"The thing about formalism is there's really no such thing as a layman's understanding
 
   A full layman's understanding would yield a formal understanding
 
   The trouble people mostly have is deciphering the notation and linking it with muscle memory"

Tweet 2:
  "Recently saw I think it was 
   @sebs_tweets
   have the aha moment that the summation symbol is just like aggregating via for loop

   That is the manner in which things start to become mentally "compressible" and therefore manageable"


Analogous to summation notation described as aggregating via for loop, we use factorials to understand
ideas like recursion (at a technical level). Given that one can remember their middle school math
introduction to factorials.

For the layman that may not remember: a factorial is an operation that represents the multiplication
of any non-negative integer (n) by every preceding integer of (n) to (1). Factorial notation is a 
number followed by an exclamation mark (!).

So, six factorial (6!) is the same as:

    6! = 6 * 5 * 4 * 3 * 2 * 1

This notation is intuitive, a time-saver, and hard to forget once learned.

In an introductory computer science class, the first approach to solving a factorial is iterative.
__________________________
----Iterative approach----
```python
def factorial(number):

    product = 1
    while(number > 1):
        product = product * number
        number = number - 1
    return product

print(factorial(6))
__________________________
```


The second approach to solving a factorial is also many students' introduction to recursion.
__________________________
----Recursive approach----
```python
def factorial(number):
    if number == 0:
        return 1
    elif number < 0:
        return f'not possible'
    else:
        return number * factorial(number - 1)
```
Note: most students would have no idea how to use an f-string.
__________________________


The script is impractical; it is an example of what is hilariously possible.
Live in the moment and build on your ideas. Even if they're silly.





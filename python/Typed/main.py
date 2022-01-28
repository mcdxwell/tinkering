
def add(a: int, b: int) -> int:
    return a + b

def sub(a: int, b: int) -> int:
    return a - b

def funky(func, a: int, b: int) -> int:
    return func(a, b)

c = funky(add, 10, 20)   
print(c) 

def whisper(txt: str) -> str:
    return txt.lower()

def yell(txt: str) -> str:
    return txt.upper()

def talk(func, txt: str) -> str:
    return func(txt)

speak = talk(yell, "hello")
print(speak)

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



foods = {'bread':2.00, 'eggs': 1.24, 'oranges': 5.47}

for food, price in foods.copy().items():
    if price > 5.0:
        del foods[food]

cheaper_food = {}
for food, price in foods.items():
    if price < 5.0:
        cheaper_food[food] = price

print(cheaper_food)

# interesting tidbit of code
for i in range(len(list(range(5)))):
    print(i)

# Match statements (example)
def tx_area_codes(code) -> str:
    match code:
        case 210:
            return "San Antonio"
        case 214 | 469 | 972:
            return "Dallas/Plano"
        case 254:
            return "Waco/Killeen" 
        case 281 | 346 | 713 | 832:
            return "Houston"    
        case 325:
            return "Abilene"
        case 361:
            return "Corpus Christi/Victoria"
        case 409:
            return "Beaumont/Galveston"
        case 430 | 903:
            return "Texarkana, Tyler, Paris"
        case 432:
            return "Midland/Odessa"
        case 512 | 737:
            return "Austin"
        case 682 | 817:
            return "Arlington/Fort Worth"
        case 806:
            return "Amarillo/Lubbock"
        case 830: 
            return "Seguin/Uvalde"
        case 915:
            return "El Paso"
        case 936:
            return "Huntsville/Nacogdoches"
        case 940:
            return "Denton/Wichita Falls"
        case 956:
            return "Brownsville/Laredo"
        case 979:
            return "Bryan/College Station"

print(tx_area_codes(281))
# Output: Houston

class Names(object):
    def __init__(self, f, m, l) :
        self.first = f
        self.middle = m
        self.last = l

class Person(Names):

    def __init__(self, f, m, l, dob, eye):
        self.dob = dob
        self.eye_color = eye
        Names.__init__(self, f, m, l)

    def full_name(self):
        print(self.first, self.middle, self.last)
        
p0 = Person("John", "Buck", "Doe", "1-1-1111", "blue")

p0.full_name()






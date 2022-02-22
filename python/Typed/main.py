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

    def __init__(self, f, m, l, age, eye):
        self.age = age
        self.eye_color = eye
        Names.__init__(self, f, m, l)
     
    def full_name(self):
        print(self.first, self.middle, self.last)

    def get_age(self):
        print(self.age)
        return self.age

    def age_limit(self):
        if self.age < 18:
            print("Under 18")
            return False
        else:
            print("Allowed")
            True




class Employee(object):
    def __init__(self, p, wage, job):
        #super().__init__(f, m, l, age, eye, wage, job)
        self.p = p
        self.wage = wage
        self.job = job

    def get_emp(self):
        print(self.p)
    def get_job(self):
        print(self.job)
    def get_wage(self):
        print(self.wage)

p0 = Person("John", "Buck", "Doe", 18, "blue")
e0 = Employee(p0, 1000, "Washer")
p0.full_name()
p0.get_age()
p0.age_limit()
e0.get_wage()
e0.get_job()


# Postfix parameter type (int)
# return type (int)
def funky(number: int) -> int:
    # if (assign square to number^2) is greater than 5
    #     print square
    if (square := number**2) > 5:
        print(f"The square of {number} is {square}")
        return square
x = funky(10)

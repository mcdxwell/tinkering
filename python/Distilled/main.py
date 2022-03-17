
# Read each line in a file
""" with open('stuff.txt') as f:
    for line in f:
        print(line, end='') """

# By using `with open`, the file is automatically closed.

# Read large files in chunks
with open('stuff.txt') as f:
    while (chunk := f.read(10000)):
        print(chunk, end='')


# Write to file - align numbers (look at the print function)
years = [2012, 112, 3000, 2022, 20, 0]
with open('out.txt', 'wt') as out:
    for year in years:
        print(f'{year:>4d}', file=out)

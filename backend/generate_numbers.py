# Handy and helpful python-based tool for generating a huge amount of numbers!

FILE_NAME = "numbers.txt"
STEP = 10
START = 0
SIZE = 100000

with open(FILE_NAME, "w") as f:
    for number in range(START, SIZE, STEP):
        f.write(f"{number}\n")

print("Generated `numbers.txt` file")

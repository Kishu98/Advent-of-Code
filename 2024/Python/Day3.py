import re

t = "hello"
memory = open("Input.txt").read()
total = 0
on = True

for match in re.findall(r"do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)", memory):
    if match == "do()":
        on = True
    elif match == "don't()":
        on = False
    elif on:
        x, y = map(int, match[4:-1].split(","))
        total += x * y

print(total)
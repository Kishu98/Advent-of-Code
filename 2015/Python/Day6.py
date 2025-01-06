import collections
import re


string = [line.strip() for line in open("Input.txt").readlines()]

pattern = r"on|off|toggle|\d{1,3},\d{1,3}"

lights = collections.defaultdict(int)

for str in string:
    match = re.findall(pattern, str)
    start = [int(num) for num in re.split(",", match[1])]
    end = [int(num) for num in re.split(",", match[2])]
    for i in range(start[0], end[0]+1):
        for j in range(start[1], end[1]+1):
            if match[0] == "on":
                lights[(i,j)] += 1
            elif match[0] == "off":
                if lights[(i,j)] == 0:
                    continue
                lights[(i,j)] -= 1
            elif match[0] == "toggle":
                lights[(i,j)] += 2

count = sum(value for value in lights.values())
print(count)

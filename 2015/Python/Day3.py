from collections import defaultdict


content = open("Input.txt").read().strip()

homes = defaultdict(int)

dir = {"^": [-1, 0], ">": [0, 1], "v": [1, 0], "<": [0, -1]}
homes[0, 0] = 1
currSantaX, currSantaY = [0, 0]
currRoboX, currRoboY = [0, 0]

for i, d in enumerate(content):
    if i % 2 == 0:
        currSantaX, currSantaY = currSantaX + dir[d][0], currSantaY + dir[d][1]
        homes[currSantaX, currSantaY] += 1
        continue

    currRoboX, currRoboY = currRoboX + dir[d][0], currRoboY + dir[d][1]
    homes[currRoboX, currRoboY] += 1

print(len(homes))

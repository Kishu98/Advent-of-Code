input = open("Input.txt").read()

floor = 0
for i,s in enumerate(input):
    if s == "(":
        floor += 1
    elif s == ")":
        floor -= 1

    if floor < 0:
        print(i)
        print(floor)
        break

# print(floor)

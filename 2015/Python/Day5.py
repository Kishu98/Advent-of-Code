nice = 0


def checkVowels(line):
    count = 0
    for char in line:
        if char in "aeiou":
            count += 1
    if count > 2:
        return True
    return False


def checkOccurance(line):
    for i in range(len(line)-1):
        if line[i] == line[i+1]:
            return True
    return False 

def checkSubString(line):
    strs = ("ab", "cd", "pq", "xy")
    for str in strs:
        if str in line:
            return False

    return True

def checkPair(line):
    l = len(line)-1
    for i in range(l):
        if f"{line[i]}{line[i+1]}" in line[i+2:]:
            return True
    return False

def checkRepeats(line):
    for i in range(len(line)-2):
        if line[i] == line[i+2]:
            return True
    return False
        

def checknice(line):
    if not checkPair(line):
        return False
    if not checkRepeats(line):
        return False
    # if not checkSubString(line):
    #     return False
    return True

input = [line.strip() for line in open("Input.txt").readlines()]

for word in input:
    if checknice(word):
        nice += 1

print(nice)

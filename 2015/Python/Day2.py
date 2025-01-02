sides = [sorted(map(int, line.strip().split("x"))) for line in open("Input.txt").readlines()]


totalArea = 0
totallenth = 0
for s1, s2, s3 in sides:
    totalArea += 2 * ((s1*s2) + (s1*s3) + (s2*s3)) + (s1*s2)
    totallenth += 2 * (s1 + s2) + (s1*s2*s3)

print(totallenth)
print(totalArea)

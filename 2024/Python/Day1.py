l, r = list(
    map(
        list,
        zip(
            *[
                list(map(int, line.split()))
                for line in open("Input.txt").read().splitlines()
            ]
        ),
    )
)


# for k in a: k.sort()
# print(sum(abs(x-y) for x, y in zip(*a)))

print(sum(x * r.count(x) for x in l))

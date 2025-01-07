lines = [
    list(instruction.strip().split(" -> "))
    for instruction in open("Input.txt").readlines()
]

instructions = {}

for line in lines:
    instructions[line[1]] = line[0]

memo = {}


def evaluate(wire, instructions, memo):
    if wire.isdigit():
        return int(wire)
    
    if wire in memo:
        return memo[wire]

    if wire not in instructions:
        raise KeyError(f"Wire '{wire}' not found in instructions")

    instruction = instructions[wire]
    parts = instruction.split()

    if "AND" in parts:
        val = evaluate(parts[0], instructions, memo) & evaluate(parts[2], instructions, memo)
    elif "OR" in parts:
        val = evaluate(parts[0], instructions, memo) | evaluate(parts[2], instructions, memo)
    elif "LSHIFT" in parts:
        val = evaluate(parts[0], instructions, memo) << int(parts[2])
    elif "RSHIFT" in parts:
        val = evaluate(parts[0], instructions, memo) >> int(parts[2])
    elif "NOT" in parts:
        val = ~evaluate(parts[1], instructions, memo) & 0xFFFF
    else:
        val = evaluate(parts[0], instructions, memo)

    memo[wire] = val
    return val

signal_a = evaluate("a", instructions, memo)
memo.clear()
memo["b"] = signal_a
print(evaluate("a", instructions, memo))


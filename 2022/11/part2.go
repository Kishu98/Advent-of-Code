package main

func part2(filename string) int {
	monkeys := getMonkeys(filename)

	modulus := 1
	for _, m := range monkeys {
		modulus *= m.divisibleBy
	}

	round := 10000
	reduceWorry := func(i int) int {
		return i % modulus
	}

	return calculateMonkeyBusiness(monkeys, round, reduceWorry)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.OpenFile("Input.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		targetVal := str_to_int(line[0])
		strNums := strings.Fields(line[1])

		var nums []int
		for _, num := range strNums {
			nums = append(nums, str_to_int(num))
		}
		if checkPossible(targetVal, nums) {
			total += targetVal
		}
	}

	fmt.Println(total)
}

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Middle *TreeNode
}

func checkPossible(targetVal int, nums []int) bool {
	root := &TreeNode{Val: nums[0]}

	buildTree(root, nums[1:])

	return evaluate(root, targetVal)
}

func buildTree(node *TreeNode, nums []int) {
	if len(nums) == 0 {
		return
	}

	node.Left = &TreeNode{Val: node.Val + nums[0]}
	node.Right = &TreeNode{Val: node.Val * nums[0]}
	node.Middle = &TreeNode{Val: str_to_int(strconv.Itoa(node.Val) + strconv.Itoa(nums[0]))}

	buildTree(node.Left, nums[1:])
	buildTree(node.Right, nums[1:])
	buildTree(node.Middle, nums[1:])

}

func evaluate(node *TreeNode, target int) bool {
	if node.Left == nil && node.Right == nil && node.Middle == nil {
		return node.Val == target
	}

	isValid := false
	if node.Left != nil {
		isValid = isValid || evaluate(node.Left, target)
	}
	if node.Right != nil {
		isValid = isValid || evaluate(node.Right, target)
	}
    if node.Middle != nil {
        isValid = isValid || evaluate(node.Middle, target)
    }

	return isValid
}

func str_to_int(s string) int {
	val, _ := strconv.Atoi(s)

	return val
}

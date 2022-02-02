package tasks

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var data = make([][]int, 0)
var trackers = make([][]bool, 0)

func Day9() {

	file, err := os.Open("C:\\Users\\all\\GolandProjects\\adventofcode\\tasks\\9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		n := scanner.Text()
		nn := strings.Split(n, "")
		m := make([]int, 0)
		for i := 0; i < len(nn); i++ {
			m = append(m, 0)
		}
		for i := 0; i < len(nn); i++ {

			m[i], _ = strconv.Atoi(nn[i])
		}
		data = append(data, m)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//part 1
	counter := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if checkRight(i, j) && checkLeft(i, j) && checkDown(i, j) && checkUp(i, j) {
				counter++
				counter += data[i][j]
				continue
			}
		}
	}
	fmt.Printf("Part 1: %d", counter)

	//part 2
	for i := 0; i < len(data); i++ {
		trackers = append(trackers, make([]bool, len(data[0])))
	}
	count := make([]int, 0)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			tmp := travel(i, j)
			if tmp != 0 {
				count = append(count, tmp)
			}

		}
	}
	sort.Ints(count)
	//fmt.Println(count)
	fmt.Printf("\nPart 2: %d", count[len(count)-1]*count[len(count)-2]*count[len(count)-3])
}

func travel(i int, j int) int {

	if i < 0 || i >= len(data) || j < 0 || j >= len(data[0]) {
		return 0
	}

	if trackers[i][j] {
		return 0
	}

	if data[i][j] == 9 {
		return 0
	}
	count := 1
	trackers[i][j] = true
	count += travel(i, j+1)
	count += travel(i, j-1)
	count += travel(i-1, j)
	count += travel(i+1, j)
	return count
}

func checkLeft(i int, j int) bool {
	if j == 0 {
		return true
	}
	if data[i][j] < data[i][j-1] {
		return true
	}
	return false
}

func checkRight(i int, j int) bool {
	if j == len(data[0])-1 {
		return true
	}

	if data[i][j] < data[i][j+1] {
		return true
	}
	return false
}

func checkDown(i int, j int) bool {
	if i == len(data)-1 {
		return true
	}
	if data[i][j] < data[i+1][j] {
		return true
	}
	return false
}

func checkUp(i int, j int) bool {
	if i == 0 {
		return true
	}
	if data[i][j] < data[i-1][j] {
		return true
	}
	return false
}

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Possibility struct {
	barn1_1 int
	barn2_1 int
	barn1_2 int
	barn2_2 int
	barn1_3 int
	barn2_3 int
	barn1_4 int
	barn2_4 int
}

func main() {
	file, err := os.ReadFile("./backforth.in")
	if err != nil {
		log.Fatal(err)
	}

	input := make([]string, len(file))
	for i, n := range file {
		input[i] = string(n)
	}

	for i, n := range input {
		if n == "\n" || n == " " {
			input = append(input[:i], input[i+1:]...)
		}
	}

	barn1Buc := input[:10]
	barn2Buc := input[10:]

	fmt.Println(input)
	fmt.Println(barn1Buc)
	fmt.Println(barn2Buc)

	barn1 := 1000
	barn2 := 1000

	pos := make([]Possibility, 0)
	posNum := 1

	for i := 0; i < 10; i++ {
		ogBarnBuc1 := barn1Buc
		ogBarnBuc2 := barn2Buc

		pos = append(pos, Possibility{barn1_1: barn1, barn2_1: barn2})
		//first day
		bucAmt1, err := strconv.ParseInt(barn1Buc[i], 10, 8)
		if err != nil {
			log.Fatal(err)
		}
		pos[posNum-1].barn1_1 -= int(100 - bucAmt1)
		fmt.Println("Barn1_1: ", pos[posNum-1].barn1_1)
		pos[posNum-1].barn2_1 += int(100)
		bucAmt1 = 0

		barn2Buc = append(barn2Buc, barn1Buc[i])
		barn1Buc = append(barn1Buc[:i], barn1Buc[i+1:]...)

		fmt.Println("Barn2_1: ", pos[posNum-1].barn2_1)
		pos[posNum-1].barn1_2 = pos[posNum-1].barn1_1
		pos[posNum-1].barn2_2 = pos[posNum-1].barn2_1

		for j := 0; j < 11; j++ {
			//second day
			bucAmt2, err := strconv.ParseInt(barn2Buc[j], 10, 8)
			if err != nil {
				log.Fatal(err)
			}
			pos[posNum-1].barn1_2 -= int(100 - bucAmt2)
			fmt.Println("Barn1_2: ", pos[posNum-1].barn1_2)
			pos[posNum-1].barn2_2 += int(100)
			bucAmt2 = 0
			fmt.Println("Barn2_2: ", pos[posNum-1].barn2_2)

			barn1Buc = append(barn1Buc, barn2Buc[i])
			barn2Buc = append(barn2Buc[:j], barn2Buc[j+1:]...)

			// for k := 0; k < 10; k++ {
			// 	//third day
			// 	bucAmt3, err := strconv.ParseInt(barn1Buc[i], 10, 8)
			// 	bucCap3, err := strconv.ParseInt(input[i+10], 10, 8)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	barn1 -= int(bucCap3 - bucAmt3)
			// 	fmt.Println("Barn1: ", barn1)
			// 	barn2 += int(bucCap3)
			// 	bucAmt3 = 0
			// 	fmt.Println("Barn2: ", barn2)

			// 	for l := 0; l < 11; l++ {
			// 		//fourth day
			// 		bucAmt4, err := strconv.ParseInt(input[i], 10, 8)
			// 		bucCap4, err := strconv.ParseInt(input[i+10], 10, 8)
			// 		if err != nil {
			// 			log.Fatal(err)
			// 		}
			// 		barn1 -= int(bucCap4 - bucAmt4)
			// 		fmt.Println("Barn1: ", barn1)
			// 		barn2 += int(bucCap4)
			// 		bucAmt4 = 0
			// 		fmt.Println("Barn2: ", barn2)
			// 	}
			// }
		}
	}
}

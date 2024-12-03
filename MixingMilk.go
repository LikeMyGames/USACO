package main

import(
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	data, err := os.ReadFile("mixmilk.in")
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)
	buckets := strings.Split(input, "\n")
	// var bucketVals [][]string
	bucketVals := make([][]int64, len(buckets))
	for i, n := range buckets {
		bucketVals[i], err = strconv.ParseInt(strings.Split(n, " "), 10, 64)
		if err!= nil{
			log.Fatal(err)
		}

	}
	fmt.Println(bucketVals)

	lastBucket := 0
	ops := 0
	for ops < 100 {
		if(bucketVals[lastBucket][1] + bucketVals[lastBucket+1][1] >= bucketVals[lastBucket+1][0]){
			bucketVals[lastBucket+1][0] = bucketVals[lastBucket][1] + bucketVals[lastBucket+1][1]
		} else if (bucketVals[lastBucket][1] + bucketVals[lastBucket+1][1] < bucketVals[lastBucket+1][0]){
			bucketVals[lastBucket+1][0] = bucketVals[lastBucket+1][0]
			bucketVals[lastBucket][1] = bucketVals[lastBucket+1][0] - bucketVals[lastBucket+1][0] - bucketVals[lastBucket+1][1]
		}



		switch lastBucket {
		case 0:
			lastBucket = 1
			break
		case 1:
			lastBucket = 2
			break
		case 2:
			lastBucket = 0
			break
		}
		ops++
	}


}
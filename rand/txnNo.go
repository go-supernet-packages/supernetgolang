package rand

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	MAX_DIGIT_TRANS_NUMBER  = 20
	MAX_PRECISION_TIME_NANO = 10000 //10^-5 second presicion of unix time
	MAX_RANDOM_NUMBER       = 9000
)

func GenerateRandomNum(maxnum int) (randnum int) {
	randnum = rand.Intn(maxnum)
	return
}

func Generate() (transnumber string) {

	localtime := (time.Now().UnixNano())            //local unix time till nanosecond
	localtime = localtime / MAX_PRECISION_TIME_NANO //local unix time till 10^-4 second
	epochTrans := strconv.FormatInt(localtime, 10)
	localNum := GenerateRandomNum(MAX_RANDOM_NUMBER) //Generate random number max 900 based on seed of local unix machine
	referanceInt := strconv.Itoa(localNum)
	effectiveTrans := referanceInt + epochTrans
	effTransInt, _ := strconv.Atoi(effectiveTrans)
	transnumber = fmt.Sprintf("%0*d", MAX_DIGIT_TRANS_NUMBER, effTransInt) // make 20 digit trans number
	//fmt.Printf("Transaction Number for %d is %s \n", i, transnumber)
	return
}

package main

import (
	"math/rand"
	"sync"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyz"

	bitsForRndInt  = 6
	rndIntMask     = 1<<bitsForRndInt - 1
	countRndIntUse = 63 / bitsForRndInt

	bitsForRndStr  = 5
	rndStrMask     = 1<<bitsForRndStr - 1
	countRndStrUse = 63 / bitsForRndStr
)

var rndSource = rand.NewSource(time.Now().UnixNano())
var muR sync.Mutex

var iRnd = rndSource.Int63()
var reUseInt = countRndIntUse

var sRnd = rndSource.Int63()
var reUseStr = countRndStrUse

func IntGenerator(outChan chan<- int) {
	for true {
		outChan <- randomInt()
	}
}

func StrGenerator(outChan chan<- [SizeStr]byte) {
	for true {
		outChan <- randomStr()
	}
}
func randomStr() [SizeStr]byte {

	var s [SizeStr]byte

	for i := SizeStr - 1; i >= 0; {
		if reUseStr == 0 {
			muR.Lock()
			sRnd = rndSource.Int63()
			muR.Unlock()
			reUseStr = countRndStrUse
		}
		idx := int(sRnd & rndStrMask)
		if idx < len(letters) {
			s[i] = letters[idx]
			i--
		}
		reUseStr--
		sRnd >>= bitsForRndStr
	}
	return s
}

func randomInt() int {
	res := -1
	for res == -1 {
		if reUseInt == 0 {
			muR.Lock()
			iRnd = rndSource.Int63()
			muR.Unlock()
			reUseInt = countRndIntUse
		}
		r := int(iRnd & rndIntMask)
		if r < len(letters) {
			res = r
		}
		reUseInt--
		iRnd >>= bitsForRndInt
	}
	return res
}

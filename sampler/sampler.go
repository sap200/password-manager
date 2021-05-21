package sampler

import (
	"math/rand"
	"strings"
	"time"
)

func Sample(str string, sampleLen int) string {
	randStr := ""
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < sampleLen; i++ {
		// generate a random index
		ri := r1.Intn(len(str))
		// Take str at that index and concatenate
		randStr += str[ri : ri+1]
	}

	return randStr

}

func Shuffle(str string, length int) string {
	randStr := []string{}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Initialize randstr
	for i := 0; i < len(str); i++ {
		randStr = append(randStr, str[i:i+1])
	}

	for i := 0; i < len(str); i++ {
		// generate a random index
		ri := r1.Intn(len(str))
		// Take str at that index and concatenate
		randStr[i] = str[ri : ri+1]
	}

	return strings.Join(randStr, "")[0:length]
}

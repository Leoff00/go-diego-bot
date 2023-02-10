package handlers

import (
	"fmt"
	"math/rand"
	"time"
)

func RandPhrase(user string) string {

	rand.Seed(time.Now().Unix())

	g1 := fmt.Sprintf("Ola %s!", user)
	g2 := fmt.Sprintf("Iaee %s!", user)
	g3 := fmt.Sprintf("Oiee %s!", user)

	arr = append(arr, g1, g2, g3)
	return arr[rand.Intn(len(arr))]

}

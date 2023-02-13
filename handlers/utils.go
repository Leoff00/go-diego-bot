package handlers

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Leoff00/go-diego-bot/envs"
)

func RandPhrase(user string) string {
	rand.Seed(time.Now().Unix())
	g1 := fmt.Sprintf("Ola %s!", user)
	g2 := fmt.Sprintf("Iaee %s!", user)
	g3 := fmt.Sprintf("Oiee %s!", user)

	arr = append(arr, g1, g2, g3)
	return arr[rand.Intn(len(arr))]
}

func ReadImg() *os.File {
	var err error

	file, err := os.Open("./img.jpg")

	if err != nil {
		fmt.Println(err)
	}

	return file

}

func PictureGenerator(param string) (*http.Response, error) {
	letters := []string{
		"nature", "city", "technology", "food", "still_life", "abstract", "wildlife",
	}

	for i, _ := range letters {
		if param != letters[i] {
			return nil, errors.New("Cannot choose a different param, please choose this params \n" + letters[i])
		}
	}

	c := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		"https://api.api-ninjas.com/v1/randomimage?category="+param,
		nil)

	key := envs.Getenv("API_KEY")
	req.Header.Add("x-api-key", key)
	req.Header.Add("Accept", "image/jpeg")

	if err != nil {
		log.Default().Fatalln(err)
	}

	res, err := c.Do(req)

	if err != nil {
		log.Default().Fatalln(err)
	}

	return res, nil
}

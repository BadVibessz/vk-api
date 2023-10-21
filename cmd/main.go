package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	http "net/http"
	"os"
	"strconv"
	vkapi "vk-api"
)

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file found")
	}
}

func main() {

	loadEnv()

	logger := log.New(os.Stderr, "", 3)

	h := http.Client{}
	client := vkapi.Client{
		Http:       &h,
		BaseURL:    "https://api.vk.com/method/",
		Retry:      false,
		RetryCount: 0}

	ctx := context.Background()

	//resp, err := client.Post(ctx, "", nil, nil)
	//if err != nil {
	//	logger.Println(err)
	//}

	token, exists := os.LookupEnv("VK_API_TOKEN")
	if !exists {
		logger.Panicln("VK_API_TOKEN not specified in env")
	}

	vk := vkapi.VkAPI{
		Token:   token,
		Version: "5.154",
		Client:  &client,
	}

	resp, err := vk.SendMessage(ctx, vkapi.Params{
		"message":   "мбэп)",
		"random_id": "0",
		"peer_id":   strconv.Itoa(2000000000 + 1),
	})
	if err != nil {
		logger.Println(err)
	}

	//resp, err := h.Get("http://www.yandex.ru")
	//if err != nil {
	//	logger.Println(err)
	//}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Println(err)
	}
	fmt.Println(PrettyString(string(buf)))
}

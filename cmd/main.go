package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	http "net/http"
	"os"
	"strconv"
	vkapi "vk-api"
)

//func ContextClient(ctx context.Context) *http.Client {
//	if ctx != nil {
//		if hc, ok := ctx.Value(HTTPClient).(*http.Client); ok {
//			return hc
//		}
//	}
//	if appengineClientHook != nil {
//		return appengineClientHook(ctx)
//	}
//	return http.DefaultClient
//}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func main() {

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

	vk := vkapi.VkAPI{
		Token:   "...",
		Version: "5.154",
		Client:  &client,
	}

	resp, err := vk.SendMessage(ctx, vkapi.Params{
		"message":   "=)",
		"random_id": strconv.Itoa(0),
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

## Wrap for VK API written on Go

### Implemented methods: 
- [x] generic call method
- [x] message.send 


### Example
```go

logger := log.New(os.Stderr, "", 3)

	h := http.Client{}
	client := vkapi.Client{
		Http:       &h,
		BaseURL:    "https://api.vk.com/method/",
		Retry:      false,
		RetryCount: 0}

	ctx := context.Background()

	token, exists := os.LookupEnv("VK_API_TOKEN")
	if !exists {
		logger.Println("VK_API_TOKEN not specified in env")
		os.Exit(1)
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

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Println(err)
	}

```
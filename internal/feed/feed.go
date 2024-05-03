package feed

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var NEYNAR_API_ROUTE = "https://api.neynar.com/v2/farcaster"

func GetGlobalTrendingFeed() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", NEYNAR_API_ROUTE+"/feed?feed_type=filter&filter_type=global_trending&with_recasts=false&limit=25", nil)

	if err != nil {
		fmt.Println("err: ", err)
	}

	neynarApiKey := os.Getenv("NEYNAR_API_KEY")
	req.Header.Add("api_key", neynarApiKey)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("err: ", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err: ", err)
	}
	return string(body)
}

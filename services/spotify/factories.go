package spotify

import (
	"fmt"
	"net/http"
	//"io/ioutil"
	"os"
)

func GetToken()  {
	endpoint := "/api/token"
	payload := strings.NewReader("grant_type=client_credentials&undefined=")
	PostData(SpotifyAccessUrl(endpoint), payload )
}
func GetTrack() {

	id := "3n3Ppam7vgaVa1iaRUc9Lp"
	endpoint := "/tracks/" + id

	GetData(endpoint)
	
}

func GetData(endpoint string)  {
	req, _ := http.NewRequest("GET", SpotifyUrl(endpoint), nil)
	
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer BQCmB26BC8Z_dEfVcPwlxnYva5mbR9VqrneWVX5K_sffxBMnqz15MEAnjMhtI2NZCB0mVjEiChmaYvTbyXk")
	
	res, _ := http.DefaultClient.Do(req)
	
	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	
	fmt.Println(res)
	//fmt.Println(string(body))
	
}
func PostData(endpoint string, payload string )  {

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", "Basic OGQ5YzYxOWUzZTkyNGMyMzk5Yjk1NTdhYjczNmQ4MDg6MGRlNzFmN2ZiYzEwNDg0MjliMDYzMTYzNzZkZDdlNWY=")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body)

}

func SpotifyAccessUrl(endpoint string) string {
	return os.Getenv("SPOTIFY_ACCOUNT_URL") + endpoint
}

func SpotifyUrl(endpoint string) string {
	return os.Getenv("SPOTIFY_URL") + endpoint
}
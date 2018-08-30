package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/redis.v5"
)

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return client
}

func createLogger(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("Could not create log file")
	}
	log.SetOutput(file)
	return file
}

type postback struct {
	Method string            `json:"method"`
	URL    string            `json:"url"`
	Data   map[string]string `json:"data"`
}

func getPostObj(client *redis.Client, postObj string) (*postback, error) {
	str, err := client.BRPop(0, postObj).Result()

	post := postback{}

	if err != nil {
		panic("Could not pull Postback Object")
	} else {
		json.Unmarshal([]byte(str[1]), &post)
	}
	return &post, nil
}

func postToURL(data postback) string {
	//loop though the data section of the postback object replace {xxx} with Date[xxx]
	for key, value := range data.Data {
		value = url.QueryEscape(value)
		re := regexp.MustCompile(regexp.QuoteMeta("{" + key + "}"))
		data.URL = re.ReplaceAllString(data.URL, value)
	}

	//if there are any unmatched {xxx} strings remove them from the final url
	re := regexp.MustCompile("{.*?}")
	data.URL = re.ReplaceAllString(data.URL, "")

	return data.URL
}

type responseData struct {
	deliveryTime string
	responseTime string
	responseCode string
	responseBody string
}

func getRequest(URL string, requestType string) (*responseData, error) {

	var getData responseData

	startTime := time.Now()
	response, err := http.Get(URL)
	endTime := time.Now()

	deliveryTime := endTime.String()
	getData.deliveryTime = deliveryTime
	responseTime := endTime.Sub(startTime).String()
	getData.responseTime = responseTime

	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		getData.responseCode = strconv.Itoa(response.StatusCode)

		body, _ := ioutil.ReadAll(response.Body)
		getData.responseBody = string(body[:])

		return &getData, nil
	}
}

func main() {
	fmt.Println("Hello world. Go away")
	client := newClient()

	logger := createLogger("log.txt")
	defer logger.Close()

	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
	}

	for {
		postObj, err := getPostObj(client, "Post_Object")
		if err != nil {
			log.Println("Could not get Post Object")
		} else if postObj != nil {
			postURL := postToURL(*postObj)
			if strings.ToLower(postObj.Method) != "get" {
				log.Println("Only accepting GET requests currently :(")
			} else {
				response, err := getRequest(postURL, (postObj.Method))
				if err != nil {
					log.Println(err)
				} else if response != nil {
					log.Println("Delivery Time " + response.deliveryTime)
					log.Println("Response Time " + response.responseTime)
					log.Println("Response Code " + response.responseCode)
					log.Println("Response Body " + response.responseBody)
				}
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

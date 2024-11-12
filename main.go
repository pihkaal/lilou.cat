package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

//go:embed public/index.html
var indexSrc string

const PORT = 3000
const CACHE_TTL = 5 * time.Minute
const STORAGE_URL = "https://storage.bunnycdn.com/lilou-cat/"
const CDN_URL = "https://lilou-cat.b-cdn.net/"

type Image struct {
	ObjectName string
}

var (
	cachedImages    []string
	cacheExpiration time.Time
)

func list_images() []string {
	if time.Now().Before(cacheExpiration) {
		return cachedImages
	}

	req, _ := http.NewRequest("GET", STORAGE_URL, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("AccessKey", os.Getenv("API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return []string{}
	}

	var images = []Image{}

	err = json.NewDecoder(res.Body).Decode(&images)
	defer res.Body.Close()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return []string{}
	}

	var imageUrls = []string{}
	for _, image := range images {
		imageUrls = append(imageUrls, fmt.Sprintf("%s/%s", CDN_URL, image.ObjectName))
	}

	cachedImages = imageUrls
	cacheExpiration = time.Now().Add(CACHE_TTL)

	return imageUrls
}

func getPage(w http.ResponseWriter, r *http.Request) {
	var images = list_images()
	var image = images[rand.Intn(len(images))]
	src := strings.Replace(indexSrc, "{{ IMAGE_URL }}", image, 1)

	io.WriteString(w, src)
}

func getFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/favicon.ico")
}

func main() {
	http.HandleFunc("/", getPage)
	http.HandleFunc("/favicon.ico", getFavicon)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	fmt.Printf("listening on port %d\n", PORT)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

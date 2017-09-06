package shorty

import (
	"log"
	"net/http"
	"os"
)

// Start starts the HTTP server
func Start() error {
	http.HandleFunc("/", handler)
	http.ListenAndServe(getEnv("BIND", ":8080"), nil)
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	url := config.Redirects[path]
	if url == "" {
		url = config.Redirects["_default"]
	}

	log.Println(path, "->", url)
	http.Redirect(w, r, url, 301)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

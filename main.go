package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/skip2/go-qrcode"
)

const defaultPort = 8080

const defaultImageSize = 256

func main() {
  port, err := strconv.Atoi(os.Getenv("PORT"))
  if err != nil || port < 1 || port > 65535 {
    port = defaultPort
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")
    if url == "" {
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte("missing query parameter `url`"))
      return
    }

    size, err := strconv.Atoi(r.URL.Query().Get("size"))
    if err != nil || size < 1 {
      size = defaultImageSize
    }

    img, err := qrcode.Encode(url, qrcode.Medium, size)
    if err != nil {
      log.Println("error encoding qr code:", err)
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte("internal error while encoding qr code"))
      return
    }

    w.Header().Add("Content-Type", "image/png")
    w.Write(img)
  })

  log.Printf("HTTP server listening on port %d\n", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

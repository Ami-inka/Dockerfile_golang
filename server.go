package main

import (
    "crypto/sha1"
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Fprintf(w, "<html><body>Error - unable to get hostname: %v</body></html>", err)
        return
    }

    // Генерация цвета на основе хеша имени хоста
    hash := sha1.New()
    hash.Write([]byte(hostname))
    hashBytes := hash.Sum(nil)
    color := fmt.Sprintf("#%02x%02x%02x", hashBytes[0], hashBytes[1], hashBytes[2]) // Используем первые три байта хеша для создания цвета

    // Форматирование ответа с фоновым цветом
    fmt.Fprintf(w, "<html><body style='background-color:%s'>Container ID: %s</body></html>", color, hostname)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

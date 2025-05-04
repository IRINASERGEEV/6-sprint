package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// Хендлер для эндпоинта /upload
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	//file, _, err := r.FormFile("file")
	//if err != nil {
	//	http.Error(w, r.Form.Get ("file"), http.StatusOK)
	//	return
	//}
	//defer file.Close()
	//w.Write(data)

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusOK)
		return
	}
	w.Write(data)

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "Conversion error", http.StatusBadRequest)
		return
	}

	filename := fmt.Sprintf("%s.txt", time.Now().UTC().Format("2006-01-02_15-04-05"))
	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		http.Error(w, "Unable to write result to file", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Conversion successful! Result saved to %s", filename)
}

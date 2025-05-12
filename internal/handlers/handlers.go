package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Парсим html-форму из файла index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// Получаем файл из формы и закрываем
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(w)
		return
	}
	defer file.Close()

	//Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}

	//Передаем данные в функцию конвертации и получаем результат конвертации
	result, err := service.Convert(string(data))
	if err != nil {
		fmt.Println(string(data))
		return
	}

	//Записываем в локальный файл результат конвертации строки
	filename := fmt.Sprintf("%s.txt", time.Now().UTC().Format("2006-01-02_15-04-05"))
	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		fmt.Println(string(data))
		return
	}
	fmt.Println(result)
	//fmt.Fprintf(w, "Conversion successful! Result saved to %s", filename)
}

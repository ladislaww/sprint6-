package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)


func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r, "index.html")
}


func uploadHandler(w http.ResponseWriter, r *http.Request) {
	
	err := r.ParseMultipartForm(10 << 20) 
	if err != nil {
		http.Error(w, "Ошибка парсинга", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("myFile") 
	if err != nil{
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file) 
	if err != nil{
		http.Error(w, "ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	converted := service.ConvertMorseOrText(string(data))

	outFile, err := os.Create("taskFile.txt") 
	if err != nil{
		http.Error(w, "ошибка создания системного файла", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()
	
	_, err = outFile.Write([]byte(converted))
	if err != nil{
		http.Error(w, "ошибка записи системного файла", http.StatusInternalServerError)
		return
	}



}
	/*	fileName := "index.html"

	file, err := os.Open(fileName) 
	if err != nil {
		http.Error(w, "не удалось открыть файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

 */


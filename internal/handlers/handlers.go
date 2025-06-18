package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)


func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r, "index.html")
}


func UploadHandler(w http.ResponseWriter, r *http.Request) {
	
	err := r.ParseMultipartForm(10 << 20) 
	if err != nil {
		http.Error(w, "Ошибка парсинга", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile") 
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

	timestamp :=time.Now().UTC().Format("20060102_150405")
	ext := filepath.Ext(handler.Filename)
	if ext == "" {
		ext = ".txt"
	}
	
	fileName := timestamp + ext


	outputFile, err := os.Create(fileName) 
	if err != nil{
		http.Error(w, "ошибка создания системного файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()
	
	_, err = outputFile.Write([]byte(converted))
	if err != nil{
		http.Error(w, "ошибка записи системного файла", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_,_ = w.Write([]byte(converted))
}


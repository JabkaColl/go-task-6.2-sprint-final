package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/JabkaColl/go-task-6.2-sprint-final/internal/service"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Читаем файл и отдаем содержимое
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Unable to read index.html: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.ConvertString(string(data))
	if err != nil {
		http.Error(w, "Conversion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	// ИСПРАВЛЕННАЯ СТРОКА - формат времени для Windows
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	newFileName := fmt.Sprintf("result_%s%s", timestamp, ext)
	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "Unable to create output file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(result)
	if err != nil {
		http.Error(w, "Unable to write to output file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "File successfully converted!\n\nOriginal: %s\nOutput: %s\n\nResult:\n%s",
		header.Filename, newFileName, result)
}

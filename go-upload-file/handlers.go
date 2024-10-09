package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func storeUpload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	defer r.MultipartForm.RemoveAll()

	uf, ufh, err := r.FormFile("media")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer uf.Close()

	storagePath := "uploads/storage/media"
	outfile := path.Join(storagePath, ufh.Filename)

	if err := os.MkdirAll(storagePath, os.ModePerm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := os.Create(outfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, uf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("uploaded"))
}

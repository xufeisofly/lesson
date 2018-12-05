package handler

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func ErrWrapper(f appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(
				w,
				http.StatusText(code),
				code)
		}
	}
}

func HandleList(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	w.Write(data)
	return nil
}

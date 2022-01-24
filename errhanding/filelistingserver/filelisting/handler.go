package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(rw http.ResponseWriter, r *http.Request) error {
	if strings.Index(r.URL.Path, prefix) != 0 {
		return userError(fmt.Sprintf("Path %s must start with %s", r.URL.Path, prefix))
	}
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	rw.Write(all)
	return nil
}

package handler

import (
	"encoding/json"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/usecase"
	"io/ioutil"
	"log"
	"net/http"
)

func uploadPicture(w http.ResponseWriter, r *http.Request)  {
	var picture = new(domain.Picture)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	err = json.Unmarshal(body, &picture)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = usecase.UploadPicture(picture)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

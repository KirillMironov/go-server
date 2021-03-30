package handler

import (
	"encoding/json"
	"github.com/KirillMironov/go-server/domain"
	"github.com/KirillMironov/go-server/pkg/usecase"
	"log"
	"net/http"
	"strconv"
)

func addItem(w http.ResponseWriter, r *http.Request)  {
	var item = new(domain.Item)

	item.Title = r.URL.Query().Get("title")
	item.Description = r.URL.Query().Get("description")
	item.Price, _ = strconv.ParseFloat(r.URL.Query().Get("price"), 64)
	item.Attributes = r.URL.Query().Get("attributes")
	item.StatusId, _ = strconv.ParseInt(r.URL.Query().Get("statusId"), 10, 64)

	id, err := usecase.CreateItem(item)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusRequestEntityTooLarge)
	}

	js, err := json.Marshal(id)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(js)
	if err != nil {
		log.Println(err)
	}

	log.Println("Item added")
}

func findItems(w http.ResponseWriter, r *http.Request)  {
	q := r.URL.Query().Get("q")

	items, err := usecase.GetItemsByTitleOrDescription(q)
	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(items)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(js)
	if err != nil {
		log.Println(err)
	}
}



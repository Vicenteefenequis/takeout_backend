package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"takeout-backend/application/database"
	"takeout-backend/domain"
	"takeout-backend/services"
)

type PostController struct {
	Connection *database.MongoDB
}

func NewPostController() *PostController {
	return &PostController{}
}

func (p *PostController) dbConnect() {
	client := database.NewConnection()
	p.Connection = client
}

func (p *PostController) postServiceInitializer() *services.PostService {
	coll := p.Connection.Client.Database("takeout_db").Collection("post")
	postService := services.NewPostService(coll)

	return postService
}

func (p *PostController) GetAll(w http.ResponseWriter, r *http.Request) {
	p.dbConnect()
	defer p.Connection.Client.Disconnect(p.Connection.Context)

	postService := p.postServiceInitializer()
	posts, err := postService.GetAll()

	if err != nil {
		sb, _ := json.Marshal(err.Error())
		w.Write(sb)
	}
	pb, _ := json.Marshal(posts)
	w.Write(pb)
}

func (p *PostController) GetById(w http.ResponseWriter, r *http.Request) {
	p.dbConnect()
	defer p.Connection.Client.Disconnect(p.Connection.Context)
	postService := p.postServiceInitializer()

	params := mux.Vars(r)
	post, err := postService.Get(params["id"])

	if err != nil {
		sb, _ := json.Marshal(err.Error())
		w.Write(sb)
	}
	pb, _ := json.Marshal(post)
	w.Write(pb)
}

func (p *PostController) Create(w http.ResponseWriter, r *http.Request) {
	p.dbConnect()
	defer p.Connection.Client.Disconnect(p.Connection.Context)
	postService := p.postServiceInitializer()

	body, err := ioutil.ReadAll(r.Body)

	var post domain.Post

	err = json.Unmarshal(body, &post)

	postResponse, _ := postService.Create(post)

	if err != nil {
		sb, _ := json.Marshal(err.Error())
		w.Write(sb)
	}

	pb, _ := json.Marshal(postResponse)
	w.Write(pb)
}

func (p *PostController) Delete(w http.ResponseWriter, r *http.Request) {
	p.dbConnect()
	defer p.Connection.Client.Disconnect(p.Connection.Context)
	postService := p.postServiceInitializer()

	params := mux.Vars(r)
	err := postService.Delete(params["id"])

	if err != nil {
		sb, _ := json.Marshal(err.Error())
		w.Write(sb)
	}

}

func (p *PostController) Update(w http.ResponseWriter, r *http.Request) {
	p.dbConnect()
	defer p.Connection.Client.Disconnect(p.Connection.Context)
	postService := p.postServiceInitializer()

	body, err := ioutil.ReadAll(r.Body)

	var post domain.Post

	err = json.Unmarshal(body, &post)

	params := mux.Vars(r)
	postResult, err := postService.Update(&post, params["id"])

	if err != nil {
		sb, _ := json.Marshal(err.Error())
		w.Write(sb)
	}

	pb, _ := json.Marshal(postResult)
	w.Write(pb)

}

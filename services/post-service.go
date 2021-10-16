package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"takeout-backend/domain"
)

type PostService struct {
	Collection *mongo.Collection
}

func NewPostService(coll *mongo.Collection) *PostService {
	return &PostService{
		Collection: coll,
	}
}

func (p *PostService) GetAll() ([]*domain.Post, error) {

	var posts []*domain.Post
	findOptions := options.Find()

	cur, err := p.Collection.Find(context.Background(), findOptions)

	defer cur.Close(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var post *domain.Post

		err := cur.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (p *PostService) Get(id string) (*domain.Post, error) {
	var post *domain.Post

	objectId, err := primitive.ObjectIDFromHex(id)

	r := p.Collection.FindOne(context.Background(), bson.D{{"_id", objectId}})
	err = r.Decode(&post)

	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostService) Create(post domain.Post) (*domain.Post, error) {

	doc := bson.D{
		{"name_user", post.NameUser},
		{"type_post", post.TypePost},
		{"visibility", post.Visibility},
		{"case_status", post.CaseStatus},
		{"image", post.Image},
		{"description", post.Description},
	}

	_, err := p.Collection.InsertOne(context.Background(), doc)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *PostService) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	_, err = p.Collection.DeleteOne(context.Background(), bson.D{{"_id", objectId}})

	if err != nil {
		return err
	}

	return nil

}

func (p *PostService) Update(post *domain.Post, id string) (*domain.Post, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	postResult := new(domain.Post)

	updatedPost := bson.D{}

	mapper(*post, &updatedPost)

	update := bson.D{{"$set", updatedPost}}

	sr := p.Collection.FindOneAndUpdate(context.Background(), bson.D{{"_id", objectId}}, update)

	err = sr.Decode(postResult)

	if err != nil {
		return nil, err
	}

	return postResult, nil
}

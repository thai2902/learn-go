package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/thai2902/mono-go/week1/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// make a Http GET request
	res, err := http.Get("http://localhost:9090/api/v1/students")
	if err != nil {
		panic(err)
	}

	// read the result
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// parse result in to struct
	var students []model.Student
	json.Unmarshal(body, &students)
	fmt.Printf("%+v", students)

	// the InsertMany method requires a slice of interface{}
	var studentIfList []interface{}
	for _, student := range students {
		studentIfList = append(studentIfList, student)
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // to avoid context leak

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://dbuser:12345678@localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database("testdb").Collection("students")
	insertResult, err := collection.InsertMany(ctx, studentIfList)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", insertResult)

}

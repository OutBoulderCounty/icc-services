package main

import (
	"context"
	"fmt"
	"os"

	database "github.com/OutBoulderCounty/icc-database"
	forms "github.com/OutBoulderCounty/icc-forms"
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Forms []forms.Form `json:"forms"`
}

func handler(ctx context.Context) (response, error) {
	var res response
	fmt.Println("Hello from the getForms Lambda!")
	db, err := database.Connect(os.Getenv("APP_ENV"))
	if err != nil {
		return res, err
	}
	defer db.Close()
	myForms, err := forms.GetForms(db.DB)
	if err != nil {
		return res, err
	}
	res.Forms = myForms
	return res, nil
}

func main() {
	lambda.Start(handler)
}

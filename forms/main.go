package forms

import (
	"context"
	"fmt"
	"os"

	// "github.com/OutBoulderCounty/icc-services/database"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type Form struct {
	ID       int    `json:"id"`
	MongoID  string `json:"mongoID"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Live     bool   `json:"live"`
}

type Element struct {
	ID       int    `json:"id"`
	MongoID  string `json:"mongoID"`
	FormID   int    `json:"formID"`
	Label    string `json:"label"`
	Type     string `json:"type"`
	Index    int    `json:"index"`
	Required bool   `json:"required"`
	Priority int    `json:"priority"`
	Search   bool   `json:"search"`
}

type Option struct {
	ID        int    `json:"id"`
	MongoID   string `json:"mongoID"`
	ElementID int    `json:"elementID"`
	Name      string `json:"name"`
	Index     int    `json:"index"`
}

type Category struct {
	ID      int    `json:"id"`
	MongoID string `json:"mongoID"`
	Name    string `json:"name"`
}

type FormCategory struct {
	ID         int `json:"id"`
	FormID     int `json:"formID"`
	CategoryID int `json:"categoryID"`
}

// TODO
func handler(ctx context.Context) error {
	sess := session.Must(session.NewSession())
	svc := ssm.New(sess)
	path := fmt.Sprintf("/icc/%s/database/", os.Getenv("APP_ENV"))
	input := ssm.GetParametersByPathInput{
		Path: &path,
	}
	out, err := svc.GetParametersByPath(&input)
	if err != nil {
		return err
	}
	params := out.Parameters
	for i := 0; i < len(params); i++ {
		fmt.Println(*params[i].Name + ": " + *params[i].Value)
	}
	// connection := database.SqlConnection{
	// 	Host: os.Getenv("DB_HOST"),
	// 	Port: os.Getenv("DB_PORT"),
	// 	User: os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// }
	// database.Connect()
	return nil
}

func main() {
	lambda.Start(handler)
}

package apiv2

import (
	"context"
	"eduid-navet/pkg/model"
	"fmt"
	"time"

	"github.com/masv3971/gonavet2"
)

type RequestNamnSokning struct {
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
	City      string `json:"city"`
}

type ReplyNamnSokning struct {
}

func (c *Client) NamnSokning(ctx context.Context, indata *RequestNamnSokning) (*ReplyNamnSokning, error) {
	result, req, err := c.navet.Search.Person(ctx, &gonavet2.RequestSearchPerson{
		GivenName: indata.GivenName,
		Surname:   indata.Surname,
		City:      indata.City,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("result", result, "req", req)
	return nil, nil
}

// Status return status for Skatteverket navet
func (c *Client) Status(ctx context.Context) (*model.Status, error) {
	manyStatus := model.ManyStatus{}

	//navetStatus, err := c.navet.Status()
	//if err != nil {
	//	return nil, err
	//}

	manyStatus = append(manyStatus, &model.Status{
		Name:      "",
		Healthy:   true,
		Status:    "STATUS_OK_navet",
		Message:   "",
		Timestamp: time.Now(),
	})

	status := manyStatus.Check()

	return status, nil
}

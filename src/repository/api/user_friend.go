package api

import "github.com/go-resty/resty"
import "encoding/json"
import "fmt"
import "sync"
import APIEntity "example_app/entity/api"

type FriendAPIRepository struct {
}

type FriendAPIRepositoryInterface interface {
	GetFriendID(id int, friendResponse *APIEntity.FriendResponse, wg *sync.WaitGroup) error
}

func ThirdPartyAPIHandler() *FriendAPIRepository{
	return &FriendAPIRepository{}
}

func (repository *FriendAPIRepository) GetFriendID(id int, friendResponse *APIEntity.FriendResponse, wg *sync.WaitGroup) error {
	url := fmt.Sprintf("https://reqres.in/api/users/%d", id)
	client := resty.New()
	resp, errClient := client.R().EnableTrace().Get(url)
	if nil != errClient {
		return errClient 
	}
	errParse := json.Unmarshal([]byte(resp.Body()), friendResponse)
	if nil != errParse {
		return errParse 
	}
	wg.Done()
	return nil
}
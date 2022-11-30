package adapter

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/DimaSSV/diadocclient/pkg/model"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strings"
)

const (
	authEndpoint = "/V3/Authenticate"
)

type Adapter struct {
	clientId string
	login    string
	password string
	token    string
	host     string
	client   http.Client
}

func New(login string, password string, clientID string, initialToken string) *Adapter {
	adapter := Adapter{
		clientId: clientID,
		login:    login,
		password: password,
		token:    initialToken,
	}
	return &adapter
}

func (a *Adapter) UpdateToken(ctx context.Context) error {
	a.token = ""
	params := make(map[string]string)
	params["type"] = "password"
	message, _ := proto.Marshal(&model.LoginPassword{
		Login:    &a.login,
		Password: &a.password,
	})
	response, err := a.CallMethod(ctx, http.MethodPost, authEndpoint, &params, message)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New("authorization error")
	}
	body, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			//log
		}
	}(response.Body)
	a.token = string(body)
	return nil
}

func (a *Adapter) CallMethod(ctx context.Context, method string, resource string, params *map[string]string, data []byte) (*http.Response, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	if len(a.token) == 0 && strings.Compare(resource, authEndpoint) != 0 {
		err = a.UpdateToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	request, err = http.NewRequestWithContext(ctx, method, "https://diadoc-api.kontur.ru", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.URL.Path = resource

	if params != nil {
		q := request.URL.Query()
		for key, value := range *params {
			q.Set(key, value)
		}
		request.URL.RawQuery = q.Encode()
	}

	if strings.Compare(resource, authEndpoint) == 0 {
		request.Header.Add("Authorization", fmt.Sprintf("DiadocAuth ddauth_api_client_id=%s", a.clientId))
	} else if len(a.token) > 0 {
		request.Header.Add("Authorization",
			fmt.Sprintf("DiadocAuth ddauth_api_client_id=%s,ddauth_token=%s", a.clientId, a.token))
	} else {
		// ??? вызвать получение токена?
	}

	response, err = a.client.Do(request)

	if response.StatusCode == 401 && strings.Compare(resource, authEndpoint) != 0 {
		err = a.UpdateToken(ctx)
		if err != nil {
			return nil, err
		}
		return a.CallMethod(ctx, method, resource, params, data)
	}
	return response, nil
}

package docflow

import (
	"context"
	"fmt"
	"github.com/DimaSSV/diadocclient/internal/adapter"
	"github.com/DimaSSV/diadocclient/pkg/model"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

const (
	getDocflowsEndpoint           = "/V3/GetDocflows"
	getDocflowsByPacketIdEndpoint = "/V3/GetDocflowsByPacketId"
	searchDocflowsEndpoint        = "/V3/SearchDocflows"
	getDocflowEventsEndpoint      = "/V3/GetDocflowEvents"
)

func GetDocflows(ctx context.Context, a *adapter.Adapter, boxID string, request *model.GetDocflowBatchRequest) (*model.GetDocflowBatchResponseV3, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(request)
	response, err := a.CallMethod(ctx, http.MethodPost, getDocflowsEndpoint, &params, message)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			//log
		}
	}(response.Body)
	switch response.StatusCode {
	case 400:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case 401:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case 402:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет прав для доступа ко всем документам организации:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} В указанном ящике нет документов с указанными идентификаторами:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.GetDocflowBatchResponseV3{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetDocflowsByPacketId(ctx context.Context, a *adapter.Adapter, boxID string, request *model.GetDocflowsByPacketIdRequest) (*model.GetDocflowsByPacketIdResponseV3, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(request)
	response, err := a.CallMethod(ctx, http.MethodPost, getDocflowsByPacketIdEndpoint, &params, message)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			//log
		}
	}(response.Body)
	switch response.StatusCode {
	case 400:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case 401:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case 402:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет прав для доступа ко всем документам организации:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.GetDocflowsByPacketIdResponseV3{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func SearchDocflows(ctx context.Context, a *adapter.Adapter, boxID string, request *model.SearchDocflowsRequest) (*model.SearchDocflowsResponseV3, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(request)
	response, err := a.CallMethod(ctx, http.MethodPost, searchDocflowsEndpoint, &params, message)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			//log
		}
	}(response.Body)
	switch response.StatusCode {
	case 400:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case 401:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case 402:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет прав для доступа ко всем документам организации:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.SearchDocflowsResponseV3{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetDocflowEvents(ctx context.Context, a *adapter.Adapter, boxID string, request *model.GetDocflowEventsRequest) (*model.GetDocflowEventsResponse, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(request)
	response, err := a.CallMethod(ctx, http.MethodPost, getDocflowEventsEndpoint, &params, message)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			//log
		}
	}(response.Body)
	switch response.StatusCode {
	case 400:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case 401:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case 402:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет прав для доступа ко всем документам организации:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.GetDocflowEventsResponse{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

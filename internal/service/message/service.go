package message

import (
	"context"
	"fmt"
	"github.com/DimaSSV/diadocclient/internal/adapter"
	"github.com/DimaSSV/diadocclient/pkg/model"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	getEntityContentEndpoint = "/V4/GetEntityContent"
	getMessageEndpoint       = "/V5/GetMessage"
	postMessageEndpoint      = "/V3/PostMessage"
	postMessagePatchEndpoint = "/V3/PostMessagePatch"
)

func GetEntityContent(ctx context.Context, a *adapter.Adapter, boxID string, messageID string, entityID string) ([]byte, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["messageId"] = messageID
	params["entityId"] = entityID
	response, err := a.CallMethod(ctx, http.MethodGet, getEntityContentEndpoint, &params, nil)
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
	case http.StatusBadRequest:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} в указанном ящике нет сообщения с идентификатором messageId, или в указанном сообщении нет сущности с идентификатором entityId, или у указанной сущности отсутствует содержимое:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return body, nil
}

func GetMessage(ctx context.Context, a *adapter.Adapter, boxID string, messageID string, entityID string, originalSignature bool, injectEntityContent bool) (*model.Message, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["messageId"] = messageID
	if entityID != "" {
		params["entityId"] = entityID
	}
	if originalSignature {
		params["originalSignature"] = ""
	}
	if injectEntityContent {
		params["injectEntityContent"] = "true"
	}
	response, err := a.CallMethod(ctx, http.MethodGet, getMessageEndpoint, &params, nil)
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
	case http.StatusBadRequest:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} в указанном ящике нет сообщений с данным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Message{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func PostMessage(ctx context.Context, a *adapter.Adapter, operationID string, post *model.MessageToPost) (*model.Message, error) {
	params := make(map[string]string)
	params["operationId"] = operationID
	message, _ := proto.Marshal(post)
	response, err := a.CallMethod(ctx, http.MethodPost, postMessageEndpoint, &params, message)
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
	case 204:
		sleepTime, err := strconv.Atoi(response.Header.Get("Retry-After"))
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(sleepTime) * time.Second)
		return PostMessage(ctx, a, operationID, post)
	case http.StatusBadRequest:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат, или отсутствуют обязательные параметры, или превышено максимально допустимое количество документов в сообщении:\n%s", string(body))
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusConflict:
		return nil, fmt.Errorf("{409} Осуществляется попытка отправить дубликат сообщения или запрещен приема документов от контрагентов согласно свойству Sociability из Organization:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Message{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func PostMessagePatch(ctx context.Context, a *adapter.Adapter, operationID string, post *model.MessagePatchToPost) (*model.MessagePatch, error) {
	params := make(map[string]string)
	params["operationId"] = operationID
	message, _ := proto.Marshal(post)
	response, err := a.CallMethod(ctx, http.MethodPost, postMessagePatchEndpoint, &params, message)
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
	case 204:
		sleepTime, err := strconv.Atoi(response.Header.Get("Retry-After"))
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(sleepTime) * time.Second)
		return PostMessagePatch(ctx, a, operationID, post)
	case http.StatusBadRequest:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат, или отсутствуют обязательные параметры, или превышено максимально допустимое количество документов в сообщении:\n%s", string(body))
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusConflict:
		return nil, fmt.Errorf("{409} осуществляется попытка отправить дубликат или запрещен прием документов от контрагентов согласно свойству Sociability в структуре …/proto/Organization:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.MessagePatch{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

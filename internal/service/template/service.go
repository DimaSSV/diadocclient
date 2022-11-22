package template

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
	getTemplateEndpoint                = "/GetTemplate"
	postTemplateEndpoint               = "/PostTemplate"
	postTemplatePatchEndpoint          = "/PostTemplatePatch"
	transformTemplateToMessageEndpoint = "/TransformTemplateToMessage"
)

func GetTemplate(ctx context.Context, a *adapter.Adapter, boxID string, templateID string, entityID string) (*model.Template, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["templateId"] = templateID
	params["entityId"] = entityID
	response, err := a.CallMethod(ctx, http.MethodGet, getTemplateEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором orgID закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return nil, fmt.Errorf("{403} Доступ к ресурсу с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Не найден ящик или шаблон с указанными идентификаторами:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Template{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func PostTemplate(ctx context.Context, a *adapter.Adapter, operationID string, post *model.TemplateToPost) (*model.Template, error) {
	params := make(map[string]string)
	params["operationId"] = operationID
	message, _ := proto.Marshal(post)
	response, err := a.CallMethod(ctx, http.MethodPost, postTemplateEndpoint, &params, message)
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
	case http.StatusNoContent:
		if value, ok := response.Header["Retry-After"]; ok {
			sleepTime, err := strconv.Atoi(value[0])
			if err != nil {
				return nil, err
			}
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
		return PostTemplate(ctx, a, operationID, post)
	case http.StatusBadRequest:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором orgID закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return nil, fmt.Errorf("{403} Доступ к ресурсу с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Не найден ящик или шаблон с указанными идентификаторами:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusConflict:
		return nil, fmt.Errorf("{409} Осуществляется попытка отправить дубликат сообщения:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Template{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func PostTemplatePatch(ctx context.Context, a *adapter.Adapter, boxID string, templateID string, operationID string, post *model.TemplatePatchToPost) (*model.MessagePatch, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["templateId"] = templateID
	params["operationId"] = operationID
	message, _ := proto.Marshal(post)
	response, err := a.CallMethod(ctx, http.MethodPost, postTemplatePatchEndpoint, &params, message)
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
	case http.StatusNoContent:
		if value, ok := response.Header["Retry-After"]; ok {
			sleepTime, err := strconv.Atoi(value[0])
			if err != nil {
				return nil, err
			}
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
		return PostTemplatePatch(ctx, a, boxID, templateID, operationID, post)
	case http.StatusBadRequest:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором orgID закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен, или нет доступа к шаблону, или отсутствуют права на создание/редактирование документов:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Не найден шаблон документа:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusConflict:
		return nil, fmt.Errorf("{409}  осуществляется попытка отклонить шаблон в неподходящем статусе:\n%s", string(body))
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

func TransformTemplateToMessage(ctx context.Context, a *adapter.Adapter, operationID string, post *model.TemplateTransformationToPost) (*model.Message, error) {
	params := make(map[string]string)
	params["operationId"] = operationID
	message, _ := proto.Marshal(post)
	response, err := a.CallMethod(ctx, http.MethodPost, transformTemplateToMessageEndpoint, &params, message)
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
	case http.StatusNoContent:
		if value, ok := response.Header["Retry-After"]; ok {
			sleepTime, err := strconv.Atoi(value[0])
			if err != nil {
				return nil, err
			}
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
		return TransformTemplateToMessage(ctx, a, operationID, post)
	case http.StatusBadRequest:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором orgID закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Не найден шаблон документа:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusConflict:
		return nil, fmt.Errorf("{409} Осуществляется попытка отправить дубликат сообщения:\n%s", string(body))
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

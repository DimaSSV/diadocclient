package event

import (
	"context"
	"fmt"
	"github.com/DimaSSV/diadocclient/internal/adapter"
	"github.com/DimaSSV/diadocclient/pkg/model"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	getEventEndpoint     = "/V2/GetEvent"
	getNewEventsEndpoint = "/V7/GetNewEvents"
	getLastEventEndpoint = "/GetLastEvent"
)

func GetEvent(ctx context.Context, a *adapter.Adapter, boxID string, eventID string) (*model.BoxEvent, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["eventId"] = eventID
	response, err := a.CallMethod(ctx, http.MethodGet, getEventEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{404} В указанном ящике нет событий с данным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.BoxEvent{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetNewEvents(
	ctx context.Context,
	a *adapter.Adapter,
	boxID string,
	afterIndexKey string,
	departmentID string,
	messageTypes []string,
	typeNamedIDs []string,
	documentDirections []string,
	timestampFromTicks int64,
	timestampToTicks int64,
	counteragentBoxID string,
	orderBy string,
	limit int,
) (*model.BoxEventList, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	if afterIndexKey != "" {
		params["afterIndexKey"] = afterIndexKey
	}
	if departmentID != "" {
		params["departmentId"] = departmentID
	}
	if messageTypes != nil {
		var buf strings.Builder
		for _, messageType := range messageTypes {
			if buf.Len() > 0 {
				buf.WriteByte(',')
			}
			buf.WriteString(messageType)
		}
		params["messageType"] = buf.String()
	}
	if typeNamedIDs != nil {
		var buf strings.Builder
		for _, typeNameDoc := range typeNamedIDs {
			if buf.Len() > 0 {
				buf.WriteByte(',')
			}
			buf.WriteString(typeNameDoc)
		}
		params["typeNamedId"] = buf.String()
	}
	if documentDirections != nil {
		var buf strings.Builder
		for _, direction := range documentDirections {
			if buf.Len() > 0 {
				buf.WriteByte(',')
			}
			buf.WriteString(direction)
		}
		params["documentDirection"] = buf.String()
	}
	if timestampFromTicks != 0 {
		params["timestampFromTicks"] = strconv.FormatInt(timestampFromTicks, 10)
	}
	if timestampToTicks != 0 {
		params["timestampToTicks"] = strconv.FormatInt(timestampToTicks, 10)
	}
	if counteragentBoxID != "" {
		params["counteragentBoxId"] = counteragentBoxID
	}
	if orderBy != "" {
		params["orderBy"] = orderBy
	}
	if limit != 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	response, err := a.CallMethod(ctx, http.MethodGet, getNewEventsEndpoint, &params, nil)
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
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.BoxEventList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetLastEvent(ctx context.Context, a *adapter.Adapter, boxID string) (*model.BoxEvent, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	response, err := a.CallMethod(ctx, http.MethodGet, getLastEventEndpoint, &params, nil)
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
		return &model.BoxEvent{}, nil
	case http.StatusBadRequest:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.BoxEvent{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

package counteragent

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
	acquireCounteragentEndpoint         = "/V2/AcquireCounteragent"
	acquireCounteragentResultEndpoint   = "/AcquireCounteragentResult"
	breakWithCounteragentEndpoint       = "/BreakWithCounteragent"
	getCounteragentV1Endpoint           = "/GetCounteragent"
	getCounteragentV2Endpoint           = "/V2/GetCounteragent"
	getCounteragentsV1Endpoint          = "/GetCounteragents"
	getCounteragentsV2Endpoint          = "/V2/GetCounteragents"
	getCounteragentCertificatesEndpoint = "/GetCounteragentCertificates"
)

func AcquireCounteragent(ctx context.Context, a *adapter.Adapter, myOrgID string, myDepartmentID string, request *model.AcquireCounteragentRequest) (*model.AsyncMethodResult, error) {
	params := make(map[string]string)
	params["myOrgId"] = myOrgID
	if myDepartmentID != "" {
		params["myDepartmentId"] = myDepartmentID
	}
	message, _ := proto.Marshal(request)
	response, err := a.CallMethod(ctx, http.MethodPost, acquireCounteragentEndpoint, &params, message)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен, или у пользователя недостаточно прав для доступа ко всем документам организации, или у пользователя нет права работать со списком контрагентов:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} в указанном ящике нет документов с указанными идентификаторами:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 409:
		return nil, fmt.Errorf("{409} Требуется заявка на роуминг для отправки приглашения роуминговому контрагенту:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.AsyncMethodResult{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func AcquireCounteragentResult(ctx context.Context, a *adapter.Adapter, taskId string) (*model.AcquireCounteragentResult, error) {
	params := make(map[string]string)
	params["taskId"] = taskId
	response, err := a.CallMethod(ctx, http.MethodGet, acquireCounteragentResultEndpoint, &params, nil)
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
		return AcquireCounteragentResult(ctx, a, taskId)
	case 400:
		return nil, fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case 401:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case 402:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя недостаточно прав для доступа ко всем документам организации:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} В указанном ящике нет документов с указанными идентификаторами:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 409:
		return nil, fmt.Errorf("{409} Не удалось выполнить запрос на приглашение контрагента:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.AcquireCounteragentResult{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func BreakWithCounteragent(ctx context.Context, a *adapter.Adapter, myOrgID string, counteragentOrgID string, comment string) error {
	params := make(map[string]string)
	params["myOrgId"] = myOrgID
	params["counteragentOrgId"] = counteragentOrgID
	if comment != "" {
		params["comment"] = comment
	}
	response, err := a.CallMethod(ctx, http.MethodPost, breakWithCounteragentEndpoint, &params, nil)
	if err != nil {
		return err
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
		return fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case 401:
		return fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case 402:
		return fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case 403:
		return fmt.Errorf("{403} Доступ к списку контрагентов организации myOrgId с предоставленным авторизационным токеном запрещен или у пользователя нет права работать со списками контрагентов:\n%s", string(body))
	case 405:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 409:
		return fmt.Errorf("{409} метод используется для отзыва приглашения с вложением:\n%s", string(body))
	case 500:
		return fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

func GetCounteragentV1(ctx context.Context, a *adapter.Adapter, myOrgID string, counteragentOrgID string) (*model.Counteragent, error) {
	params := make(map[string]string)
	params["myOrgId"] = myOrgID
	params["counteragentOrgId"] = counteragentOrgID
	response, err := a.CallMethod(ctx, http.MethodGet, getCounteragentV1Endpoint, &params, nil)
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором myOrgId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к списку контрагентов организации myOrgId с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} Партнерские отношения между организациями myOrgId и counteragentOrgId не установлены:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Counteragent{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetCounteragentV2(ctx context.Context, a *adapter.Adapter, myOrgID string, counteragentOrgID string) (*model.Counteragent, error) {
	params := make(map[string]string)
	params["myOrgId"] = myOrgID
	params["counteragentOrgId"] = counteragentOrgID
	response, err := a.CallMethod(ctx, http.MethodGet, getCounteragentV2Endpoint, &params, nil)
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором myOrgId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к списку контрагентов организации myOrgId с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} Партнерские отношения между организациями myOrgId и counteragentOrgId не установлены:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Counteragent{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetCounteragentsV1(ctx context.Context, a *adapter.Adapter, myOrgID string, counteragentStatus string, afterIndexKey string) (*model.CounteragentList, error) {
	params := make(map[string]string)
	params["myOrgId"] = myOrgID
	if counteragentStatus != "" {
		params["counteragentStatus"] = counteragentStatus
	}
	if afterIndexKey != "" {
		params["afterIndexKey"] = afterIndexKey
	}
	response, err := a.CallMethod(ctx, http.MethodGet, getCounteragentsV1Endpoint, &params, nil)
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором myOrgId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к списку контрагентов организации myOrgId с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} Партнерские отношения между организациями myOrgId и counteragentOrgId не установлены:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.CounteragentList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetCounteragentsV2(ctx context.Context, a *adapter.Adapter, myOrgID string, counteragentStatus string, afterIndexKey string) (*model.CounteragentList, error) {
	params := make(map[string]string)
	params["myOrgId"] = myOrgID
	if counteragentStatus != "" {
		params["counteragentStatus"] = counteragentStatus
	}
	if afterIndexKey != "" {
		params["afterIndexKey"] = afterIndexKey
	}
	response, err := a.CallMethod(ctx, http.MethodGet, getCounteragentsV2Endpoint, &params, nil)
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором myOrgId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к списку контрагентов организации myOrgId с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} Партнерские отношения между организациями myOrgId и counteragentOrgId не установлены:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.CounteragentList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetCounteragentCertificates(ctx context.Context, a *adapter.Adapter, myOrgID string, counteragentOrgID string) (*model.CounteragentCertificateList, error) {
	params := make(map[string]string)
	params["myOrgId"] = myOrgID
	params["counteragentOrgId"] = counteragentOrgID
	response, err := a.CallMethod(ctx, http.MethodGet, getCounteragentCertificatesEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором myOrgId закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к списку сертификатов организации counteragentOrgId от организации myOrgId с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} Партнерские отношения между организациями myOrgId и counteragentOrgId не установлены:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.CounteragentCertificateList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

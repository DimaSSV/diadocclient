package employee

import (
	"context"
	"fmt"
	"github.com/DimaSSV/diadocclient/internal/adapter"
	"github.com/DimaSSV/diadocclient/pkg/model"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strconv"
)

const (
	createEmployeeEndpoint       = "/CreateEmployee"
	deleteEmployeeEndpoint       = "/DeleteEmployee"
	getMyUser2Endpoint           = "/V2/GetMyUser"
	getMyUserEndpoint            = "/GetMyUser"
	getEmployeeEndpoint          = "/GetEmployee"
	getEmployeesEndpoint         = "/GetEmployees"
	getMyEmployeeEndpoint        = "/GetMyEmployee"
	getOrganizationUsersEndpoint = "/GetOrganizationUsers"
	getSubscriptionsEndpoint     = "/GetSubscriptions"
	updateEmployeeEndpoint       = "/UpdateEmployee"
	updateMyUserEndpoint         = "/UpdateMyUser"
	updateSubscriptions          = "/UpdateSubscriptions"
	getMyCertificatesEndpoint    = "/GetMyCertificates"
)

func CreateEmployee(ctx context.Context, a *adapter.Adapter, boxID string, create *model.EmployeeToCreate) (*model.Employee, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(create)
	response, err := a.CallMethod(ctx, http.MethodPost, createEmployeeEndpoint, &params, message)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Employee{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func DeleteEmployee(ctx context.Context, a *adapter.Adapter, boxID string, userID string) error {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["userId"] = userID
	response, err := a.CallMethod(ctx, http.MethodPost, deleteEmployeeEndpoint, &params, nil)
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
		return fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return fmt.Errorf("{404} Не найден сотрудник с указанным идентификатором:\n%s", string(body))
	case 405:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

func GetEmployee(ctx context.Context, a *adapter.Adapter, boxID string, userID string) (*model.Employee, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["userId"] = userID
	response, err := a.CallMethod(ctx, http.MethodGet, getEmployeeEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} В указанном ящике нет пользователя с указанным идентификатором:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Employee{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetEmployees(ctx context.Context, a *adapter.Adapter, boxID string, page int, count int) (*model.EmployeeList, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["page"] = strconv.Itoa(page)
	params["count"] = strconv.Itoa(count)
	response, err := a.CallMethod(ctx, http.MethodGet, getEmployeesEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} Указанного ящика не существует:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.EmployeeList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetMyEmployee(ctx context.Context, a *adapter.Adapter, boxID string) (*model.Employee, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	response, err := a.CallMethod(ctx, http.MethodGet, getMyEmployeeEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} В указанном ящике нет пользователя с указанным идентификатором:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Employee{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetMyUserV2(ctx context.Context, a *adapter.Adapter) (*model.UserV2, error) {
	response, err := a.CallMethod(ctx, http.MethodGet, getMyUser2Endpoint, nil, nil)
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
	case 401:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.UserV2{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetMyUser(ctx context.Context, a *adapter.Adapter) (*model.User, error) {
	response, err := a.CallMethod(ctx, http.MethodGet, getMyUserEndpoint, nil, nil)
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
	case 401:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.User{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetOrganizationUsers(ctx context.Context, a *adapter.Adapter, orgID string) (*model.OrganizationUsersList, error) {
	params := make(map[string]string)
	params["orgId"] = orgID
	response, err := a.CallMethod(ctx, http.MethodGet, getOrganizationUsersEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} Организация с указанным идентификатором не найдена в справочнике:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.OrganizationUsersList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetSubscriptions(ctx context.Context, a *adapter.Adapter, boxID string, userID string) (*model.EmployeeSubscriptions, error) {
	params := make(map[string]string)
	params["boxID"] = boxID
	params["userID"] = userID
	response, err := a.CallMethod(ctx, http.MethodGet, getSubscriptionsEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора и не от имени пользователя, подписки которого запрошены:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} В указанном ящике нет пользователя с указанным идентификатором:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.EmployeeSubscriptions{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateEmployee(ctx context.Context, a *adapter.Adapter, boxID string, userID string, update *model.EmployeeToUpdate) (*model.Employee, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["userId"] = userID
	message, _ := proto.Marshal(update)
	response, err := a.CallMethod(ctx, http.MethodPost, updateEmployeeEndpoint, &params, message)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Employee{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateMyUser(ctx context.Context, a *adapter.Adapter, update *model.UserToUpdate) (*model.UserV2, error) {
	message, _ := proto.Marshal(update)
	response, err := a.CallMethod(ctx, http.MethodPost, updateMyUserEndpoint, nil, message)
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
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.UserV2{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateSubscriptions(ctx context.Context, a *adapter.Adapter, boxID string, userID string, update *model.SubscriptionsToUpdate) (*model.EmployeeSubscriptions, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["userId"] = userID
	message, _ := proto.Marshal(update)
	response, err := a.CallMethod(ctx, http.MethodPost, updateSubscriptions, &params, message)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора и не от имени пользователя, подписки которого редактируются:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} в указанном ящике нет пользователя с указанным идентификатором:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.EmployeeSubscriptions{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetMyCertificates(ctx context.Context, a *adapter.Adapter, boxID string) (*model.CertificateList, error) {
	params := make(map[string]string)
	params["boxID"] = boxID
	response, err := a.CallMethod(ctx, http.MethodGet, getMyCertificatesEndpoint, &params, nil)
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
	case 401:
		return nil, fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case 402:
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.CertificateList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

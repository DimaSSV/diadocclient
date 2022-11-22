package department

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
	getDepartmentsEndpoint   = "/admin/GetDepartments"
	getDepartmentEndpoint    = "/admin/GetDepartment"
	createDepartmentEndpoint = "/admin/CreateDepartment"
	updateDepartment         = "/admin/UpdateDepartment"
	deleteDepartment         = "/admin/DeleteDepartment"
)

func GetDepartmentFull(ctx context.Context, a *adapter.Adapter, boxID string, departmentID string) (*model.DepartmentAdmin, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["departmentId"] = departmentID
	response, err := a.CallMethod(ctx, http.MethodGet, getDepartmentEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Пв указанном ящике нет подразделения с указанным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.DepartmentAdmin{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetDepartmentsFull(ctx context.Context, a *adapter.Adapter, boxID string, page int, count int) (*model.DepartmentList, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["page"] = strconv.Itoa(page)
	params["count"] = strconv.Itoa(count)
	response, err := a.CallMethod(ctx, http.MethodGet, getDepartmentsEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} В указанном ящике нет подразделения с указанным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.DepartmentList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateDepartment(ctx context.Context, a *adapter.Adapter, boxID string, create *model.DepartmentToCreate) error {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(create)
	response, err := a.CallMethod(ctx, http.MethodPost, createDepartmentEndpoint, &params, message)
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
	case http.StatusBadRequest:
		return fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case http.StatusUnauthorized:
		return fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case http.StatusNotFound:
		return fmt.Errorf("{404} В указанном ящике нет подразделения с указанным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

func UpdateDepartment(ctx context.Context, a *adapter.Adapter, boxID string, departmentID string, update *model.DepartmentToUpdate) error {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["departmentId"] = departmentID
	message, _ := proto.Marshal(update)
	response, err := a.CallMethod(ctx, http.MethodPost, updateDepartment, &params, message)
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
	case http.StatusBadRequest:
		return fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case http.StatusUnauthorized:
		return fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case http.StatusNotFound:
		return fmt.Errorf("{404} В указанном ящике нет подразделения с указанным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

func DeleteDepartment(ctx context.Context, a *adapter.Adapter, boxID string, departmentID string) error {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["departmentId"] = departmentID
	response, err := a.CallMethod(ctx, http.MethodPost, deleteDepartment, &params, nil)
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
	case http.StatusBadRequest:
		return fmt.Errorf("{400} Данные в запросе имеют неверный формат или отсутствуют обязательные параметры:\n%s", string(body))
	case http.StatusUnauthorized:
		return fmt.Errorf("{401} В запросе отсутствует HTTP-заголовок Authorization или в этом заголовке содержатся некорректные авторизационные данные:\n%s", string(body))
	case http.StatusPaymentRequired:
		return fmt.Errorf("{402} У организации с указанным идентификатором boxId закончилась подписка на API:\n%s", string(body))
	case http.StatusForbidden:
		return fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case http.StatusNotFound:
		return fmt.Errorf("{404} В указанном ящике нет подразделения с указанным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusConflict:
		return fmt.Errorf("{409} Запрещено удалить подразделение в переданным departmentId:\n%s", string(body))
	case http.StatusInternalServerError:
		return fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

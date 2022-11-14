package organization

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
	getMyOrganizationsEndpoint        = "/GetMyOrganizations"
	getOrganizationEndpoint           = "/GetOrganization"
	getOrganizationByKPPEndpoint      = "/GetOrganizationsByInnKpp"
	getBoxEndpointEndpoint            = "/GetBox"
	getDepartmentEndpoint             = "/GetDepartment"
	getOrganizationsByInnListEndpoint = "/GetOrganizationsByInnList"
	getOrganizationFeaturesEndpoint   = "/GetOrganizationFeatures"
)

func GetBox(ctx context.Context, a *adapter.Adapter, boxID string) (*model.Box, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	response, err := a.CallMethod(ctx, http.MethodGet, getBoxEndpointEndpoint, &params, nil)
	if err != nil {
		return nil, err
	}
	result := model.Box{}
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
		return nil, fmt.Errorf("{404} Ящик с указанным идентификатором не найден в справочнике:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetDepartment(ctx context.Context, a *adapter.Adapter, orgID string, departmentId string) (*model.Department, error) {
	params := make(map[string]string)
	params["orgId"] = orgID
	params["departmentId"] = departmentId
	response, err := a.CallMethod(ctx, http.MethodGet, getDepartmentEndpoint, &params, nil)
	if err != nil {
		return nil, err
	}
	result := model.Department{}
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором orgID закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к подразделению с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} не найдена организация или подразделение с указанным идентификатором:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetMyOrganizations(ctx context.Context, a *adapter.Adapter, autoRegister bool) (*model.OrganizationList, error) {

	params := make(map[string]string)
	if !autoRegister {
		params["autoRegister"] = "false"
	}
	response, err := a.CallMethod(ctx, http.MethodGet, getMyOrganizationsEndpoint, &params, nil)
	if err != nil {
		return nil, err
	}
	result := model.OrganizationList{}
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором orgID закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к подразделению с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} не найдена организация или подразделение с указанным идентификатором:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetOrganizationByOrgID(ctx context.Context, a *adapter.Adapter, orgID string) (*model.Organization, error) {
	params := make(map[string]string)
	params["orgId"] = orgID
	return getOrganization(ctx, a, params)
}

func GetOrganizationByBoxID(ctx context.Context, a *adapter.Adapter, boxID string) (*model.Organization, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	return getOrganization(ctx, a, params)
}

func GetOrganizationByFnsParticipantId(ctx context.Context, a *adapter.Adapter, fnsParticipantID string) (*model.Organization, error) {
	params := make(map[string]string)
	params["fnsParticipantId"] = fnsParticipantID
	return getOrganization(ctx, a, params)
}

func GetOrganizationByINN(ctx context.Context, a *adapter.Adapter, INN string, KPP string) (*model.Organization, error) {
	params := make(map[string]string)
	params["inn"] = INN
	if KPP != "" {
		params["kpp"] = KPP
	}
	return getOrganization(ctx, a, params)
}

func getOrganization(ctx context.Context, a *adapter.Adapter, params map[string]string) (*model.Organization, error) {
	response, err := a.CallMethod(ctx, http.MethodGet, getOrganizationEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором orgID закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к подразделению с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} организация с указанным идентификатором не найдена в справочнике:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Organization{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetOrganizationsByInnKpp(ctx context.Context, a *adapter.Adapter, INN string, KPP string, includeRelations bool) (*model.OrganizationList, error) {
	params := make(map[string]string)
	params["inn"] = INN
	if KPP != "" {
		params["kpp"] = KPP
	}
	if includeRelations {
		params["includeRelations"] = "true"
	}
	response, err := a.CallMethod(ctx, http.MethodGet, getOrganizationByKPPEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором orgID закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к подразделению с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} организация с указанным идентификатором не найдена в справочнике:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.OrganizationList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetOrganizationsByInnList(ctx context.Context, a *adapter.Adapter, myOrgId string, INNs []string) (*model.GetOrganizationsByInnListResponse, error) {
	params := make(map[string]string)
	params["inn"] = myOrgId
	requestBody, err := proto.Marshal(&model.GetOrganizationsByInnListRequest{
		InnList: INNs,
	})
	if err != nil {
		return nil, err
	}
	response, err := a.CallMethod(ctx, http.MethodPost, getOrganizationsByInnListEndpoint, &params, requestBody)
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
		return nil, fmt.Errorf("{402} У организации с указанным идентификатором orgID закончилась подписка на API:\n%s", string(body))
	case 403:
		return nil, fmt.Errorf("{403} Доступ к подразделению с предоставленным авторизационным токеном запрещен или запрос сделан не от имени администратора:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} Организация с указанным идентификатором не найдена в справочнике:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.GetOrganizationsByInnListResponse{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetOrganizationFeatures(ctx context.Context, a *adapter.Adapter, boxID string) (*model.OrganizationFeatures, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	response, err := a.CallMethod(ctx, http.MethodGet, getOrganizationFeaturesEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case 404:
		return nil, fmt.Errorf("{404} Организация с указанным идентификатором не найдена в справочнике:\n%s", string(body))
	case 405:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 500:
		return nil, fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.OrganizationFeatures{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

package document

import (
	"context"
	"fmt"
	"github.com/DimaSSV/diadocclient/internal/adapter"
	"github.com/DimaSSV/diadocclient/pkg/model"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	deleteEndpoint                             = "/Delete"
	detectCustomPrintFormsEndpoint             = "/DetectCustomPrintForms"
	forwardDocumentEndpoint                    = "/V2/ForwardDocument"
	getDocumentEndpoint                        = "/V3/GetDocument"
	getDocumentsEndpoint                       = "/V3/GetDocuments"
	getDocumentsByMessageIdEndpoint            = "/GetDocumentsByMessageId"
	getForwardedDocumentEventsEndpoint         = "/V2/GetForwardedDocumentEvents"
	getResolutionRoutesForOrganizationEndpoint = "/GetResolutionRoutesForOrganization"
	getForwardedEntityContentEndpoint          = "/V2/GetForwardedEntityContent"
	getForwardedDocumentsEndpoint              = "/V2/GetForwardedDocuments"
	getGeneratedPrintFormEndpoint              = "/GetGeneratedPrintForm"
	moveDocumentsEndpoint                      = "/MoveDocuments"
	recycleDraftEndpoint                       = "/RecycleDraft"
	restoreEndpoint                            = "/Restore"
	shelfDownloadEndpoint                      = "/ShelfDownload"
	shelfUploadEndpoint                        = "/ShelfUpload"
	sendDraftEndpoint                          = "/SendDraft"
	maxFilePartForShelf                        = 512 * 1024
)

func Delete(ctx context.Context, a *adapter.Adapter, boxID string, messageID string, documentID string) error {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["messageId"] = messageID
	params["documentId"] = documentID
	response, err := a.CallMethod(ctx, http.MethodPost, deleteEndpoint, &params, nil)
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
		return fmt.Errorf("{404} Не найдено сообщение или документ с указанными идентификаторами.:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 409:
		return fmt.Errorf("{409} Осуществляется попытка повторного удаления документа или сообщения:\n%s", string(body))
	case http.StatusInternalServerError:
		return fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

func DetectCustomPrintForms(ctx context.Context, a *adapter.Adapter, boxID string, request *model.CustomPrintFormDetectionRequest) (*model.CustomPrintFormDetectionResult, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(request)
	response, err := a.CallMethod(ctx, http.MethodPost, detectCustomPrintFormsEndpoint, &params, message)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.CustomPrintFormDetectionResult{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ForwardDocument(ctx context.Context, a *adapter.Adapter, boxID string, request *model.ForwardDocumentRequest) (*model.ForwardDocumentResponse, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(request)
	response, err := a.CallMethod(ctx, http.MethodPost, forwardDocumentEndpoint, &params, message)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.ForwardDocumentResponse{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//ToDo: Реализовать методы
//GenerateTitleXml
//GenerateDocumentProtocol
//GenerateDocumentZip
//GenerateForwardedDocumentPrintForm
//GenerateForwardedDocumentProtocol
//GenerateReceiptXml
//GeneratePrintForm
//GeneratePrintFormFromAttachment
//GenerateRevocationRequestXml
//GenerateSignatureRejectionXml

func GetDocument(ctx context.Context, a *adapter.Adapter, boxID string, messageID string, entityID string, injectEntityContent bool) (*model.Document, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["messageId"] = messageID
	params["entityId"] = entityID
	if injectEntityContent {
		params["injectEntityContent"] = "true"
	}
	response, err := a.CallMethod(ctx, http.MethodGet, getDocumentEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} В указанном ящике не найдено сообщение с идентификатором messageId или в указанном сообщении нет сущности типа LetterAttachment с идентификатором entityId:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.Document{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type Filter struct {
	BoxID                 string
	FilterCategory        string
	CounteragentBoxID     string
	FromDepartmentID      string
	ToDepartmentID        string
	DocumentNumber        string
	TimestampFromTicks    time.Time
	TimestampToTicks      time.Time
	FromDocumentDate      time.Time
	ToDocumentDate        time.Time
	DepartmentID          string
	ExcludeSubdepartments bool
	AfterIndexKey         string
	SortDirection         string
	Count                 int
}

func NewFilter(boxID string, AfterIndexKey string) Filter {
	return Filter{
		BoxID:                 boxID,
		FilterCategory:        "Any.Inbound",
		CounteragentBoxID:     "",
		FromDepartmentID:      "",
		ToDepartmentID:        "",
		DocumentNumber:        "",
		TimestampFromTicks:    time.Time{},
		TimestampToTicks:      time.Time{},
		FromDocumentDate:      time.Time{},
		ToDocumentDate:        time.Time{},
		DepartmentID:          "",
		ExcludeSubdepartments: false,
		AfterIndexKey:         AfterIndexKey,
		SortDirection:         "",
		Count:                 100,
	}
}

func GetDocuments(ctx context.Context, a *adapter.Adapter, filter Filter) (*model.DocumentList, error) {
	return getDocuments(
		ctx,
		a,
		filter.BoxID,
		filter.FilterCategory,
		filter.CounteragentBoxID,
		filter.DepartmentID,
		filter.ToDepartmentID,
		filter.DocumentNumber,
		filter.TimestampFromTicks,
		filter.TimestampToTicks,
		filter.FromDocumentDate,
		filter.ToDocumentDate,
		filter.DepartmentID,
		filter.ExcludeSubdepartments,
		filter.AfterIndexKey,
		filter.SortDirection,
		filter.Count,
	)
}

func getDocuments(
	ctx context.Context,
	a *adapter.Adapter,
	boxID string,
	filterCategory string,
	counteragentBoxID string,
	fromDepartmentID string,
	toDepartmentID string,
	documentNumber string,
	timestampFromTicks time.Time,
	timestampToTicks time.Time,
	fromDocumentDate time.Time,
	toDocumentDate time.Time,
	departmentID string,
	excludeSubdepartments bool,
	afterIndexKey string,
	sortDirection string,
	count int,
) (*model.DocumentList, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["filterCategory"] = filterCategory
	if counteragentBoxID != "" {
		params["counteragentBoxId"] = counteragentBoxID
	}
	if fromDepartmentID != "" {
		params["fromDepartmentId"] = fromDepartmentID
	}
	if toDepartmentID != "" {
		params["toDepartmentId"] = toDepartmentID
	}
	if documentNumber != "" {
		params["documentNumber"] = documentNumber
	}
	if !timestampFromTicks.IsZero() {
		params["timestampFromTicks"] = strconv.FormatInt(timestampFromTicks.UnixNano(), 10)
	}
	if !timestampToTicks.IsZero() {
		params["timestampToTicks"] = strconv.FormatInt(timestampToTicks.UnixNano(), 10)
	}
	if !fromDocumentDate.IsZero() {
		params["fromDocumentDate"] = fromDocumentDate.Format("02.01.2006")
	}
	if !toDocumentDate.IsZero() {
		params["toDocumentDate "] = toDocumentDate.Format("02.01.2006")
	}
	if departmentID != "" {
		params["departmentId"] = departmentID
	}
	if excludeSubdepartments {
		params["excludeSubdepartments"] = "true"
	}
	if afterIndexKey != "" {
		params["afterIndexKey"] = afterIndexKey
	}
	if sortDirection != "" {
		params["sortDirection"] = sortDirection
	}
	if count != 0 {
		params["count"] = strconv.Itoa(count)
	}

	response, err := a.CallMethod(ctx, http.MethodGet, getDocumentsEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{404} В указанном ящике не найдено сообщение с идентификатором messageId или в указанном сообщении нет сущности типа LetterAttachment с идентификатором entityId:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.DocumentList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetDocumentsByMessageId(ctx context.Context, a *adapter.Adapter, boxID string, messageID string) (*model.DocumentList, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["messageId"] = messageID
	response, err := a.CallMethod(ctx, http.MethodGet, getDocumentsByMessageIdEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} В указанном ящике не найдено сообщение с идентификатором messageId или в указанном сообщении нет сущности типа LetterAttachment с идентификатором entityId:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.DocumentList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetForwardedDocumentEvents(ctx context.Context, a *adapter.Adapter, boxID string, request *model.GetForwardedDocumentEventsRequest) (*model.GetForwardedDocumentEventsResponse, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(request)
	response, err := a.CallMethod(ctx, http.MethodPost, getForwardedDocumentEventsEndpoint, &params, message)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Не найдено сообщение с заданным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.GetForwardedDocumentEventsResponse{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetResolutionRoutesForOrganization(ctx context.Context, a *adapter.Adapter, orgID string) (*model.ResolutionRouteList, error) {
	params := make(map[string]string)
	params["orgId"] = orgID
	response, err := a.CallMethod(ctx, http.MethodGet, getResolutionRoutesForOrganizationEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Не найдено сообщение с заданным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.ResolutionRouteList{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetForwardedEntityContent(ctx context.Context, a *adapter.Adapter, boxID string, fromBoxID string, messageID string, documentID string, forwardEventID string, entityID string) ([]byte, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["fromBoxId"] = fromBoxID
	params["messageId"] = messageID
	params["documentId"] = documentID
	params["forwardEventId"] = forwardEventID
	params["entityId"] = entityID
	response, err := a.CallMethod(ctx, http.MethodGet, getForwardedEntityContentEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Не найдено сообщение/документ/сущность с заданным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return body, nil
}

func GetForwardedDocuments(ctx context.Context, a *adapter.Adapter, boxID string, request *model.GetForwardedDocumentsRequest) (*model.GetForwardedDocumentsResponse, error) {
	params := make(map[string]string)
	params["boxId"] = boxID
	message, _ := proto.Marshal(request)
	response, err := a.CallMethod(ctx, http.MethodPost, getForwardedDocumentsEndpoint, &params, message)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Не найдено сообщение/документ с заданным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	result := model.GetForwardedDocumentsResponse{}
	err = proto.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetGeneratedPrintForm(ctx context.Context, a *adapter.Adapter, printFormID string) ([]byte, error) {
	params := make(map[string]string)
	params["printFormId"] = printFormID
	response, err := a.CallMethod(ctx, http.MethodGet, getGeneratedPrintFormEndpoint, &params, nil)
	if err != nil {
		return nil, err
	}
	if value, ok := response.Header["Retry-After"]; ok {
		sleepTime, err := strconv.Atoi(value[0])
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(sleepTime) * time.Second)
		return GetGeneratedPrintForm(ctx, a, printFormID)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Не найдено сформированной печатной формы:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return body, nil
}

func MoveDocuments(ctx context.Context, a *adapter.Adapter, operation *model.DocumentsMoveOperation) error {
	message, _ := proto.Marshal(operation)
	response, err := a.CallMethod(ctx, http.MethodPost, moveDocumentsEndpoint, nil, message)
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
		return fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

//ToDO: Ещё методы
//ParseTitleXml
//ParseRevocationRequestXml
//ParseSignatureRejectionXml
//PrepareDocumentsToSign

func RecycleDraft(ctx context.Context, a *adapter.Adapter, boxID string, draftID string) error {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["draftId"] = draftID
	response, err := a.CallMethod(ctx, http.MethodPost, recycleDraftEndpoint, &params, nil)
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
		return fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusNotFound:
		return fmt.Errorf("{404} Не найден черновик с указанным идентификатором:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 409:
		return fmt.Errorf("{409} Осуществляется попытка удаления уже утилизированного черновика:\n%s", string(body))
	case http.StatusInternalServerError:
		return fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

func Restore(ctx context.Context, a *adapter.Adapter, boxID string, messageID string, documentID string) error {
	params := make(map[string]string)
	params["boxId"] = boxID
	params["messageId"] = messageID
	if documentID != "" {
		params["documentId"] = documentID
	}
	response, err := a.CallMethod(ctx, http.MethodPost, restoreEndpoint, &params, nil)
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
		return fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusNotFound:
		return fmt.Errorf("{404} Не найдено сообщение или документ с указанными идентификаторами:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 409:
		return fmt.Errorf("{409} Осуществляется попытка повторного восстановления документа или сообщения:\n%s", string(body))
	case http.StatusInternalServerError:
		return fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

func ShelfDownload(ctx context.Context, a *adapter.Adapter, nameOnShelf string) ([]byte, error) {
	params := make(map[string]string)
	params["nameOnShelf"] = nameOnShelf
	response, err := a.CallMethod(ctx, http.MethodGet, shelfDownloadEndpoint, &params, nil)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Файл на полке не найден:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("{500} при обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return body, nil
}

func ShelfUpload(ctx context.Context, a *adapter.Adapter, data []byte) (string, error) {
	var nameOnShelf = fmt.Sprintf("api-%s", uuid.NewString())
	params := make(map[string]string)
	params["nameOnShelf"] = nameOnShelf
	parts := splitDataArray(data)
	for i, part := range parts {
		if i == len(parts) {
			params["isLastPart"] = "1"
		}
		err := shelfUploadPart(ctx, a, &params, part)
		if err != nil {
			return "", err
		}
	}
	return nameOnShelf, nil
}

func shelfUploadPart(ctx context.Context, a *adapter.Adapter, params *map[string]string, dataPart []byte) error {
	response, err := a.CallMethod(ctx, http.MethodPost, shelfUploadEndpoint, params, dataPart)
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
		return fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case http.StatusInternalServerError:
		return fmt.Errorf("{500} При обработке запроса возникла непредвиденная ошибка:\n%s", string(body))
	}
	return nil
}

func splitDataArray(data []byte) [][]byte {
	var result [][]byte
	for len(data) > maxFilePartForShelf {
		result = append(result, data[:maxFilePartForShelf])
		data = data[maxFilePartForShelf:]
	}
	if len(data) > 0 {
		result = append(result, data)
	}
	return result
}

func SendDraft(ctx context.Context, a *adapter.Adapter, operationID string, send *model.DraftToSend) (*model.Message, error) {
	params := make(map[string]string)
	params["operationId"] = operationID
	message, _ := proto.Marshal(send)
	response, err := a.CallMethod(ctx, http.MethodPost, sendDraftEndpoint, &params, message)
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
		return nil, fmt.Errorf("{403} Доступ к ящику с предоставленным авторизационным токеном запрещен или у пользователя нет доступа к каким-то документам из запроса:\n%s", string(body))
	case http.StatusNotFound:
		return nil, fmt.Errorf("{404} Файл на полке не найден:\n%s", string(body))
	case http.StatusMethodNotAllowed:
		return nil, fmt.Errorf("{405} Используется неподходящий HTTP-метод:\n%s", string(body))
	case 409:
		return nil, fmt.Errorf("{409} Осуществляется попытка отправить дубликат сообщения, указан несуществующий идентификатор содержимого документа, подготовленного к отправке, или запрещен прием документов от контрагентов согласно свойству Sociability в структуре Organization:\n%s", string(body))
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

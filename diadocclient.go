package diadocсlient

import (
	"context"
	"github.com/DimaSSV/diadocclient/internal/adapter"
	"github.com/DimaSSV/diadocclient/internal/service/counteragent"
	"github.com/DimaSSV/diadocclient/internal/service/department"
	"github.com/DimaSSV/diadocclient/internal/service/docflow"
	"github.com/DimaSSV/diadocclient/internal/service/document"
	"github.com/DimaSSV/diadocclient/internal/service/employee"
	"github.com/DimaSSV/diadocclient/internal/service/event"
	"github.com/DimaSSV/diadocclient/internal/service/message"
	"github.com/DimaSSV/diadocclient/internal/service/organization"
	"github.com/DimaSSV/diadocclient/internal/service/template"
	"github.com/DimaSSV/diadocclient/pkg/model"
)

type DiadocClient struct {
	adapter *adapter.Adapter
}

func New(login string, password string, clientID string, initialToken string) (DiadocClient, error) {
	client := DiadocClient{
		adapter: adapter.New(login, password, clientID, initialToken),
	}
	if len(initialToken) == 0 {
		if err := client.adapter.UpdateToken(context.Background()); err != nil {
			return client, err
		}
	}
	return client, nil
}

/////////////////////////////////////////////////////////////////
////////////////Работа с организациями///////////////////////////
/////////////////////////////////////////////////////////////////

func (c DiadocClient) GetBox(ctx context.Context, boxID string) (*model.Box, error) {
	return organization.GetBox(ctx, c.adapter, boxID)
}

func (c DiadocClient) GetDepartment(ctx context.Context, orgID string, departmentId string) (*model.Department, error) {
	return organization.GetDepartment(ctx, c.adapter, orgID, departmentId)
}

func (c DiadocClient) GetMyOrganizations(ctx context.Context) (*model.OrganizationList, error) {
	return organization.GetMyOrganizations(ctx, c.adapter, true)
}

func (c DiadocClient) GetOrganizationByOrgID(ctx context.Context, orgID string) (*model.Organization, error) {
	return organization.GetOrganizationByOrgID(ctx, c.adapter, orgID)
}

func (c DiadocClient) GetOrganizationByBoxID(ctx context.Context, boxID string) (*model.Organization, error) {
	return organization.GetOrganizationByBoxID(ctx, c.adapter, boxID)
}

func (c DiadocClient) GetOrganizationByFnsParticipantId(ctx context.Context, fnsParticipantID string) (*model.Organization, error) {
	return organization.GetOrganizationByFnsParticipantId(ctx, c.adapter, fnsParticipantID)
}

func (c DiadocClient) GetOrganizationByINN(ctx context.Context, INN string, KPP string) (*model.Organization, error) {
	return organization.GetOrganizationByINN(ctx, c.adapter, INN, KPP)
}

func (c DiadocClient) GetOrganizationsByInnKpp(ctx context.Context, INN string, KPP string, includeRelations bool) (*model.OrganizationList, error) {
	return organization.GetOrganizationsByInnKpp(ctx, c.adapter, INN, KPP, includeRelations)
}

func (c DiadocClient) GetOrganizationsByInnList(ctx context.Context, myOrgId string, INNs []string) (*model.GetOrganizationsByInnListResponse, error) {
	return organization.GetOrganizationsByInnList(ctx, c.adapter, myOrgId, INNs)
}

func (c DiadocClient) GetOrganizationFeatures(ctx context.Context, boxID string) (*model.OrganizationFeatures, error) {
	return organization.GetOrganizationFeatures(ctx, c.adapter, boxID)
}

/////////////////////////////////////////////////////////////////
////////////////Работа с сотрудниками////////////////////////////
/////////////////////////////////////////////////////////////////

func (c DiadocClient) CreateEmployee(ctx context.Context, boxID string, create *model.EmployeeToCreate) (*model.Employee, error) {
	return employee.CreateEmployee(ctx, c.adapter, boxID, create)
}

func (c DiadocClient) DeleteEmployee(ctx context.Context, boxID string, userID string) error {
	return employee.DeleteEmployee(ctx, c.adapter, boxID, userID)
}

func (c DiadocClient) GetEmployee(ctx context.Context, boxID string, userID string) (*model.Employee, error) {
	return employee.GetEmployee(ctx, c.adapter, boxID, userID)
}

func (c DiadocClient) GetEmployees(ctx context.Context, boxID string, page int, count int) (*model.EmployeeList, error) {
	return employee.GetEmployees(ctx, c.adapter, boxID, page, count)
}

func (c DiadocClient) GetMyEmployee(ctx context.Context, boxID string) (*model.Employee, error) {
	return employee.GetMyEmployee(ctx, c.adapter, boxID)
}

func (c DiadocClient) GetMyUserV2(ctx context.Context) (*model.UserV2, error) {
	return employee.GetMyUserV2(ctx, c.adapter)
}

func (c DiadocClient) GetMyUserV1(ctx context.Context) (*model.User, error) {
	return employee.GetMyUser(ctx, c.adapter)
}

func (c DiadocClient) GetOrganizationUsers(ctx context.Context, orgID string) (*model.OrganizationUsersList, error) {
	return employee.GetOrganizationUsers(ctx, c.adapter, orgID)
}

func (c DiadocClient) GetSubscriptions(ctx context.Context, boxID string, userID string) (*model.EmployeeSubscriptions, error) {
	return employee.GetSubscriptions(ctx, c.adapter, boxID, userID)
}

func (c DiadocClient) UpdateEmployee(ctx context.Context, boxID string, userID string, update *model.EmployeeToUpdate) (*model.Employee, error) {
	return employee.UpdateEmployee(ctx, c.adapter, boxID, userID, update)
}

func (c DiadocClient) UpdateMyUser(ctx context.Context, update *model.UserToUpdate) (*model.UserV2, error) {
	return employee.UpdateMyUser(ctx, c.adapter, update)
}

func (c DiadocClient) UpdateSubscriptions(ctx context.Context, boxID string, userID string, update *model.SubscriptionsToUpdate) (*model.EmployeeSubscriptions, error) {
	return employee.UpdateSubscriptions(ctx, c.adapter, boxID, userID, update)
}

func (c DiadocClient) GetMyCertificates(ctx context.Context, boxID string) (*model.CertificateList, error) {
	return employee.GetMyCertificates(ctx, c.adapter, boxID)
}

///////////////////////////////////////////////////////////////////
/////////////////////Работа с подразделениями//////////////////////
///////////////////////////////////////////////////////////////////

func (c DiadocClient) GetDepartmentFull(ctx context.Context, boxID string, departmentID string) (*model.DepartmentAdmin, error) {
	return department.GetDepartmentFull(ctx, c.adapter, boxID, departmentID)
}

func (c DiadocClient) GetDepartmentsFull(ctx context.Context, boxID string, page int, count int) (*model.DepartmentList, error) {
	return department.GetDepartmentsFull(ctx, c.adapter, boxID, page, count)
}

func (c DiadocClient) CreateDepartment(ctx context.Context, boxID string, create *model.DepartmentToCreate) error {
	return department.CreateDepartment(ctx, c.adapter, boxID, create)
}

func (c DiadocClient) UpdateDepartment(ctx context.Context, boxID string, departmentID string, update *model.DepartmentToUpdate) error {
	return department.UpdateDepartment(ctx, c.adapter, boxID, departmentID, update)
}

func (c DiadocClient) DeleteDepartment(ctx context.Context, boxID string, departmentID string) error {
	return department.DeleteDepartment(ctx, c.adapter, boxID, departmentID)
}

///////////////////////////////////////////////////////////////////
/////////////////////Работа с контрагентами////////////////////////
///////////////////////////////////////////////////////////////////

func (c DiadocClient) AcquireCounteragent(ctx context.Context, myOrgID string, myDepartmentID string, request *model.AcquireCounteragentRequest) (*model.AsyncMethodResult, error) {
	return counteragent.AcquireCounteragent(ctx, c.adapter, myOrgID, myDepartmentID, request)
}

func (c DiadocClient) AcquireCounteragentResult(ctx context.Context, taskID string) (*model.AcquireCounteragentResult, error) {
	return counteragent.AcquireCounteragentResult(ctx, c.adapter, taskID)
}

func (c DiadocClient) BreakWithCounteragent(ctx context.Context, myOrgID string, counteragentOrgID string, comment string) error {
	return counteragent.BreakWithCounteragent(ctx, c.adapter, myOrgID, counteragentOrgID, comment)
}

func (c DiadocClient) GetCounteragentV1(ctx context.Context, myOrgID string, counteragentOrgID string) (*model.Counteragent, error) {
	return counteragent.GetCounteragentV1(ctx, c.adapter, myOrgID, counteragentOrgID)
}

func (c DiadocClient) GetCounteragentV2(ctx context.Context, myOrgID string, counteragentOrgID string) (*model.Counteragent, error) {
	return counteragent.GetCounteragentV2(ctx, c.adapter, myOrgID, counteragentOrgID)
}

func (c DiadocClient) GetCounteragentsV1(ctx context.Context, myOrgID string, counteragentStatus string, afterIndexKey string) (*model.CounteragentList, error) {
	return counteragent.GetCounteragentsV1(ctx, c.adapter, myOrgID, counteragentStatus, afterIndexKey)
}

func (c DiadocClient) GetCounteragentsV2(ctx context.Context, myOrgID string, counteragentStatus string, afterIndexKey string) (*model.CounteragentList, error) {
	return counteragent.GetCounteragentsV2(ctx, c.adapter, myOrgID, counteragentStatus, afterIndexKey)
}

func (c DiadocClient) GetCounteragentCertificates(ctx context.Context, myOrgID string, counteragentOrgID string) (*model.CounteragentCertificateList, error) {
	return counteragent.GetCounteragentCertificates(ctx, c.adapter, myOrgID, counteragentOrgID)
}

///////////////////////////////////////////////////////////////////
/////////////////////Работа с сообщениями//////////////////////////
///////////////////////////////////////////////////////////////////

func (c DiadocClient) GetEntityContent(ctx context.Context, boxID string, messageID string, entityID string) ([]byte, error) {
	return message.GetEntityContent(ctx, c.adapter, boxID, messageID, entityID)
}

func (c DiadocClient) GetMessage(ctx context.Context, boxID string, messageID string, entityID string, originalSignature bool, injectEntityContent bool) (*model.Message, error) {
	return message.GetMessage(ctx, c.adapter, boxID, messageID, entityID, originalSignature, injectEntityContent)
}

func (c DiadocClient) PostMessage(ctx context.Context, operationID string, post *model.MessageToPost) (*model.Message, error) {
	return message.PostMessage(ctx, c.adapter, operationID, post)
}

func (c DiadocClient) PostMessagePatch(ctx context.Context, operationID string, post *model.MessagePatchToPost) (*model.MessagePatch, error) {
	return message.PostMessagePatch(ctx, c.adapter, operationID, post)
}

///////////////////////////////////////////////////////////////////
/////////////////////Работа с событиями////////////////////////////
///////////////////////////////////////////////////////////////////

func (c DiadocClient) GetEvent(ctx context.Context, boxID string, eventID string) (*model.BoxEvent, error) {
	return event.GetEvent(ctx, c.adapter, boxID, eventID)
}

func (c DiadocClient) GetNewEvents(
	ctx context.Context,
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
	return event.GetNewEvents(
		ctx,
		c.adapter,
		boxID,
		afterIndexKey,
		departmentID,
		messageTypes,
		typeNamedIDs,
		documentDirections,
		timestampFromTicks,
		timestampToTicks,
		counteragentBoxID,
		orderBy,
		limit,
	)
}

func (c DiadocClient) GetLastEvent(ctx context.Context, boxID string) (*model.BoxEvent, error) {
	return event.GetLastEvent(ctx, c.adapter, boxID)
}

///////////////////////////////////////////////////////////////////
//////////////////Работа с документообором/////////////////////////
///////////////////////////////////////////////////////////////////

func (c DiadocClient) GetDocflows(ctx context.Context, boxID string, request *model.GetDocflowBatchRequest) (*model.GetDocflowBatchResponseV3, error) {
	return docflow.GetDocflows(ctx, c.adapter, boxID, request)
}

func (c DiadocClient) GetDocflowsByPacketId(ctx context.Context, boxID string, request *model.GetDocflowsByPacketIdRequest) (*model.GetDocflowsByPacketIdResponseV3, error) {
	return docflow.GetDocflowsByPacketId(ctx, c.adapter, boxID, request)
}

func (c DiadocClient) SearchDocflows(ctx context.Context, boxID string, request *model.SearchDocflowsRequest) (*model.SearchDocflowsResponseV3, error) {
	return docflow.SearchDocflows(ctx, c.adapter, boxID, request)
}

func (c DiadocClient) GetDocflowEvents(ctx context.Context, boxID string, request *model.GetDocflowEventsRequest) (*model.GetDocflowEventsResponse, error) {
	return docflow.GetDocflowEvents(ctx, c.adapter, boxID, request)
}

///////////////////////////////////////////////////////////////////
//////////////////////Работа с документами/////////////////////////
///////////////////////////////////////////////////////////////////

func (c DiadocClient) Delete(ctx context.Context, boxID string, messageID string, documentID string) error {
	return document.Delete(ctx, c.adapter, boxID, messageID, documentID)
}

func (c DiadocClient) DetectCustomPrintForms(ctx context.Context, boxID string, request *model.CustomPrintFormDetectionRequest) (*model.CustomPrintFormDetectionResult, error) {
	return document.DetectCustomPrintForms(ctx, c.adapter, boxID, request)
}

func (c DiadocClient) ForwardDocument(ctx context.Context, boxID string, request *model.ForwardDocumentRequest) (*model.ForwardDocumentResponse, error) {
	return document.ForwardDocument(ctx, c.adapter, boxID, request)
}

func (c DiadocClient) GetDocument(ctx context.Context, boxID string, messageID string, entityID string, injectEntityContent bool) (*model.Document, error) {
	return document.GetDocument(ctx, c.adapter, boxID, messageID, entityID, injectEntityContent)
}

func (c DiadocClient) GetDocuments(ctx context.Context, filter document.Filter) (*model.DocumentList, error) {
	return document.GetDocuments(ctx, c.adapter, filter)
}

func (c DiadocClient) GetDocumentsByMessageId(ctx context.Context, boxID string, messageID string) (*model.DocumentList, error) {
	return document.GetDocumentsByMessageId(ctx, c.adapter, boxID, messageID)
}

func (c DiadocClient) GetForwardedDocumentEvents(ctx context.Context, boxID string, request *model.GetForwardedDocumentEventsRequest) (*model.GetForwardedDocumentEventsResponse, error) {
	return document.GetForwardedDocumentEvents(ctx, c.adapter, boxID, request)
}

func (c DiadocClient) GetResolutionRoutesForOrganization(ctx context.Context, orgID string) (*model.ResolutionRouteList, error) {
	return document.GetResolutionRoutesForOrganization(ctx, c.adapter, orgID)
}

func (c DiadocClient) GetForwardedEntityContent(ctx context.Context, boxID string, fromBoxID string, messageID string, documentID string, forwardEventID string, entityID string) ([]byte, error) {
	return document.GetForwardedEntityContent(ctx, c.adapter, boxID, fromBoxID, messageID, documentID, forwardEventID, entityID)
}

func (c DiadocClient) GetForwardedDocuments(ctx context.Context, boxID string, request *model.GetForwardedDocumentsRequest) (*model.GetForwardedDocumentsResponse, error) {
	return document.GetForwardedDocuments(ctx, c.adapter, boxID, request)
}

func (c DiadocClient) GetGeneratedPrintForm(ctx context.Context, printFormID string) ([]byte, error) {
	return document.GetGeneratedPrintForm(ctx, c.adapter, printFormID)
}

func (c DiadocClient) MoveDocuments(ctx context.Context, operation *model.DocumentsMoveOperation) error {
	return document.MoveDocuments(ctx, c.adapter, operation)
}

func (c DiadocClient) RecycleDraft(ctx context.Context, boxID string, draftID string) error {
	return document.RecycleDraft(ctx, c.adapter, boxID, draftID)
}

func (c DiadocClient) Restore(ctx context.Context, boxID string, messageID string, documentID string) error {
	return document.Restore(ctx, c.adapter, boxID, messageID, documentID)
}

func (c DiadocClient) ShelfDownload(ctx context.Context, nameOnShelf string) ([]byte, error) {
	return document.ShelfDownload(ctx, c.adapter, nameOnShelf)
}

func (c DiadocClient) ShelfUpload(ctx context.Context, data []byte) (string, error) {
	return document.ShelfUpload(ctx, c.adapter, data)
}

func (c DiadocClient) SendDraft(ctx context.Context, operationID string, send *model.DraftToSend) (*model.Message, error) {
	return document.SendDraft(ctx, c.adapter, operationID, send)
}

///////////////////////////////////////////////////////////////////
//////////////////////Работа с шаблонами///////////////////////////
///////////////////////////////////////////////////////////////////

func (c DiadocClient) GetTemplate(ctx context.Context, boxID string, templateID string, entityID string) (*model.Template, error) {
	return template.GetTemplate(ctx, c.adapter, boxID, templateID, entityID)
}

func (c DiadocClient) PostTemplate(ctx context.Context, operationID string, post *model.TemplateToPost) (*model.Template, error) {
	return template.PostTemplate(ctx, c.adapter, operationID, post)
}

func (c DiadocClient) PostTemplatePatch(ctx context.Context, boxID string, templateID string, operationID string, post *model.TemplatePatchToPost) (*model.MessagePatch, error) {
	return template.PostTemplatePatch(ctx, c.adapter, boxID, templateID, operationID, post)
}

func (c DiadocClient) TransformTemplateToMessage(ctx context.Context, operationID string, post *model.TemplateTransformationToPost) (*model.Message, error) {
	return template.TransformTemplateToMessage(ctx, c.adapter, operationID, post)
}

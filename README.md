# diadocer
diadocer - реализация клиента для взаимодействия с [API Диадока](http://api-docs.diadoc.ru/).

Для использования библиотеки в вашем проекте необходимо выполнить команду
```shell
go get github.com/DimaSSV/diadocer
```

Для работы клиента необходимо задать переменные окружения
- DIADOC_LOGIN
- DIADOC_PASSWORD
- DIADOC_CLIENT_ID - указать ключ разработчика, полученный по запросу у компании Диадок

Все реализованные функции доступны через экземпляр структуры "DiadocClient"

Пример получения информации по текущему пользователю
```go
package main

import (
	"context"
	"diadocer"
	"diadocer/internal/config"
)

func main() {
	client := diadocer.New()
	u, _ := client.GetMyUserV2(context.Background())
	println(u.String())
}
```

На текущий момент реализовано:
1. [Работа с документами](https://developer.kontur.ru/docs/diadoc-api/API_Documents.html) пакет document - Реализовано частично
2. [Работа с сообщениями](https://developer.kontur.ru/docs/diadoc-api/API_Messages.html) пакет message
3. [Работа с событиями](https://developer.kontur.ru/docs/diadoc-api/API_Events.html) пакет event
4. [Работа с организациями](https://developer.kontur.ru/docs/diadoc-api/API_Organizations.html) пакет organization
5. [Работа с подразделениями](https://developer.kontur.ru/docs/diadoc-api/API_Departments.html) пакет department
6. [Работа с сотрудниками](https://developer.kontur.ru/docs/diadoc-api/API_Employees.html) пакет employee
7. [Работа с контрагентами](https://developer.kontur.ru/docs/diadoc-api/API_Counteragents.html) пакет counteragent
8. [Работа с шаблонами](https://developer.kontur.ru/docs/diadoc-api/API_Templates.html) пакет template
9. [Docflow API](https://developer.kontur.ru/docs/diadoc-api/Docflow%20API.html) пакет docflow

Не реализовано: 
1. [Работа со счетами-фактурами](https://developer.kontur.ru/docs/diadoc-api/API_Invoices.html)
2. [Работа с УПД](https://developer.kontur.ru/docs/diadoc-api/API_UniversalTransferDocument.html)
3. [Регистрация организации и сотрудника по сертификату](https://developer.kontur.ru/docs/diadoc-api/API_Registration.html)
4. [Подпись Контур.Сертификатом](https://developer.kontur.ru/docs/diadoc-api/CloudSignApi.html)
5. [Подпись сертификатом без носителя](https://developer.kontur.ru/docs/diadoc-api/API_Dss.html)

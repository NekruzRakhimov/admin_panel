package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"log"
	"net/http"
	"strconv"
)

//CreateContract contract godoc
// @Summary Creating contract
// @Description Creating contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  contract  body models.Contract true "creating contract"
// @Param  type  query string true "type of contract"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/{type} [post]
func CreateContract(c *gin.Context) {
	var contract models.Contract

	if err := c.BindJSON(&contract); err != nil {
		log.Println("[controller.CreateContract]|[c.BindJSO]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	contract.Type = c.Param("type")

	if err := service.CreateContract(contract); err != nil {
		log.Println("[controller.CreateContract]|[service.CreateContract]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новый договор был успешно создан!"})
}

//AddAdditionalAgreement contract godoc
// @Summary Creating additional agreement
// @Description Creating additional agreement
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  contract  body models.Contract true "creating contract"
// @Param  id  path string true "id договора на основе которого создаётся ДС"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/additional_agreement/{id} [post]
func AddAdditionalAgreement(c *gin.Context) {
	var contract models.Contract
	prevContractIdStr := c.Param("id")
	prevContractId, err := strconv.Atoi(prevContractIdStr)
	if err != nil {
		log.Println("[controller.EditContract]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	if err := c.BindJSON(&contract); err != nil {
		log.Println("[controller.EditContract]|[c.BindJSO]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	contract.PrevContractId = prevContractId

	if err := service.AddAdditionalAgreement(contract); err != nil {
		log.Println("[controller.EditContract]|[service.EditContract]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "дополнительное соглашение успешно создано"})
}

//EditContract contract godoc
// @Summary Editing contract
// @Description Editing contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  contract  body models.Contract true "editing contract"
// @Param  id  path string true "id of contract"
// @Param  type  path string true "type of contract"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/{type}/{id} [put]
func EditContract(c *gin.Context) {
	var contract models.Contract
	contract.Type = c.Param("type")

	contractIdStr := c.Param("id")
	contractId, err := strconv.Atoi(contractIdStr)
	if err != nil {
		log.Println("[controller.EditContract]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	if err := c.BindJSON(&contract); err != nil {
		log.Println("[controller.EditContract]|[c.BindJSO]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	contract.ID = contractId

	if err := service.EditContract(contract); err != nil {
		log.Println("[controller.EditContract]|[service.EditContract]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "договор отправлен на согласование!"})
}

//GetAllContracts contract godoc
// @Summary Get All Contracts
// @Description Gel All Contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  status  query string false "status of contract"
// @Success 200 {array}  models.ContractMiniInfo
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/ [get]
func GetAllContracts(c *gin.Context) {

	contractType := c.Query("status")
	contractsMiniInfo, err := service.GetAllContracts(contractType)
	if err != nil {
		log.Println("[controller.GetAllContracts]|[service.GetAllContracts]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contractsMiniInfo)
}

//GetContractDetails contract godoc
// @Summary Get Contract Details
// @Description Gel Contract Details
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of contract"
// @Success 200 {object}  models.Contract
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/{id}/details [get]
func GetContractDetails(c *gin.Context) {
	contractIdStr := c.Param("id")
	contractId, err := strconv.Atoi(contractIdStr)
	if err != nil {
		log.Println("[controller.GetContractDetails]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	contract, err := service.GetContractDetails(contractId)
	if err != nil {
		log.Println("[controller.GetContractDetails]|[service.GetContractDetails]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contract)
}

//ConformContract contract godoc
// @Summary Conform contract
// @Description Conform contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of contract"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/conform/{id} [put]
func ConformContract(c *gin.Context) {
	contractIdStr := c.Param("id")
	contractId, err := strconv.Atoi(contractIdStr)
	if err != nil {
		log.Println("[controller.ConformContract]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.ConformContract(contractId, "в работе"); err != nil {
		log.Println("[controller.ConformContract]|[service.ConformContract]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "контракт успешно согласован"})
}

//CancelContract contract godoc
// @Summary Cancel contract
// @Description Cancel contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of contract"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/cancel/{id} [put]
func CancelContract(c *gin.Context) {
	contractIdStr := c.Param("id")
	contractId, err := strconv.Atoi(contractIdStr)
	if err != nil {
		log.Println("[controller.CancelContract]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.CancelContract(contractId); err != nil {
		log.Println("[controller.CancelContract]|[service.CancelContract]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "договор был успешно отменён"})
}

//GetContractHistory contract godoc
// @Summary Get Contract History
// @Description Gel Contract History
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of contract"
// @Success 200 {array}  models.Contract
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/history/{id} [get]
func GetContractHistory(c *gin.Context) {
	contractIdStr := c.Param("id")
	contractId, err := strconv.Atoi(contractIdStr)
	if err != nil {
		log.Println("[controller.GetContractHistory]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	contracts, err := service.GetContractHistory(contractId)
	if err != nil {
		log.Println("[controller.GetContractHistory]|[service.GetContractHistory]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contracts)
}

//GetContractStatusChangesHistory contract godoc
// @Summary Get Contract Status Changes History
// @Description Gel Contract Status Changes History
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of contract"
// @Success 200 {array}  models.ContractStatusHistory
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/status_history/{id} [get]
func GetContractStatusChangesHistory(c *gin.Context) {
	contractIdStr := c.Param("id")
	contractId, err := strconv.Atoi(contractIdStr)
	if err != nil {
		log.Println("[controller.GetContractStatusChangesHistory]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	contractStatusHistory, err := service.GetContractStatusChangesHistory(contractId)
	if err != nil {
		log.Println("[controller.GetContractStatusChangesHistory]|[service.GetContractStatusChangesHistory]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contractStatusHistory)
}

//FinishContract contract godoc
// @Summary Finish contract
// @Description Finish contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of contract"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/finish/{id} [put]
func FinishContract(c *gin.Context) {
	contractIdStr := c.Param("id")
	contractId, err := strconv.Atoi(contractIdStr)
	if err != nil {
		log.Println("[controller.FinishContract]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := service.FinishContract(contractId); err != nil {
		log.Println("[controller.FinishContract]|[service.FinishContract]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "договор был успешно завершён"})
}

//RevisionContract contract godoc
// @Summary Revision contract
// @Description Revision contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  id  path string true "id of contract"
// @Param  comment  query string true "comment of contract"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/revision/{id} [put]
func RevisionContract(c *gin.Context) {
	contractIdStr := c.Param("id")
	contractId, err := strconv.Atoi(contractIdStr)
	if err != nil {
		log.Println("[controller.RevisionContract]|[strconv.Atoi]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	comment := c.Query("comment")
	if err := service.RevisionContract(contractId, comment); err != nil {
		log.Println("[controller.RevisionContract]|[service.RevisionContract]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "добавьте комментарий перед отправкой договора на доработки"})
		return
	}

	fmt.Println("id: ", contractId)
	fmt.Println("comment: ", comment)
	c.JSON(http.StatusOK, gin.H{"reason": "договор был на отправлен доработку!"})
}

func GetProductsTemplate(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("files/applications/products_template.xlsx")
}

func ConvertExcelToStruct(c *gin.Context) {
	img, err := c.FormFile("file")
	if err != nil {
		log.Println("[controller.ConvertExcelToStruct]|[c.FormFile(\"file\")]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	//file, err := os.Create("files/applications/products_template.xlsx")
	//if err != nil {
	//	log.Println("[controller.ConvertExcelToStruct]|[os.Create]| error is: ", err.Error())
	//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
	//	return
	//}

	if err := c.SaveUploadedFile(img, "files/applications/edited_products_template.xlsx"); err != nil {
		log.Println("[controller.ConvertExcelToStruct]|[c.SaveUploadedFile]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	f, err := excelize.OpenFile("files/applications/edited_products_template.xlsx")
	//c.JSON(http.StatusOK, gin.H{"reason": "ok"})
	if err != nil {
		log.Println("[controller.ConvertExcelToStruct]|[excelize.OpenFile]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	var products []models.Product
	counter := 2
	for {
		var product models.Product
		product.Sku, err = f.GetCellValue("page1", fmt.Sprintf("A%d", counter))
		if err != nil {
			log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		product.ProductName, err = f.GetCellValue("page1", fmt.Sprintf("B%d", counter))
		if err != nil {
			log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		priceStr, err := f.GetCellValue("page1", fmt.Sprintf("C%d", counter))
		if err != nil {
			log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		product.Currency, err = f.GetCellValue("page1", fmt.Sprintf("D%d", counter))
		if err != nil {
			log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		//product.Substance, err = f.GetCellValue("page1", fmt.Sprintf("E%d", counter))
		//if err != nil {
		//	log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
		//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		//	return
		//}

		//product.StorageCondition, err = f.GetCellValue("page1", fmt.Sprintf("F%d", counter))
		//if err != nil {
		//	log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
		//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		//	return
		//}

		//product.Producer, err = f.GetCellValue("page1", fmt.Sprintf("G%d", counter))
		//if err != nil {
		//	log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
		//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		//	return
		//}

		if product.ProductNumber == "" && product.ProductName == "" && priceStr == "" && product.Currency == "" {
			break
		}

		if product.Sku == "" {
			c.JSON(http.StatusBadRequest, gin.H{"reason": "не все номера товаров были заполнены, проверьте заполненность полей"})
			return
		}

		if product.ProductName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"reason": "не все торговые названия были заполнены, проверьте заполненность полей"})
			return
		}

		if priceStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"reason": "не все цены заполнены, проверьте заполненность полей"})
			return
		}

		product.Price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("test: ", priceStr)
			log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if product.Currency == "" {
			c.JSON(http.StatusBadRequest, gin.H{"reason": "не все валюты заполнены, проверьте заполненность полей"})
			return
		}

		products = append(products, product)
		counter++
	}

	c.JSON(http.StatusOK, products)
}

// CounterpartyContract GetCounterpartyContract contract godoc
// @Summary Get CounterpartyContract
// @Description Берет данные контрагента
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param client path string true "BINClient"
// @Success 200 {array}  models.Counterparty
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /counterparty/{client} [get]
func CounterpartyContract(c *gin.Context) {
	binClient := c.Param("client")

	contract, err := service.CounterpartyContract(binClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contract)

}

// SearchBinClient godoc
// @Summary      Search Client
// @Description  add by json account
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        bin  body      models.ClientBin  true  "bin"
// @Success      200      {object}  models.Client
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /client_search [post]
func SearchBinClient(c *gin.Context) {
	var clientBin models.ClientBin

	err := c.ShouldBindJSON(&clientBin)
	if err != nil {
		c.JSON(http.StatusBadRequest, "что-то пошло не так")
		return
	}

	client, err := service.SearchByBinClient(clientBin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, client)

}

func AddIndividualContract(c *gin.Context) {
	img, err := c.FormFile("file")
	if err != nil {
		log.Println("[controller.AddIndividualContract]|[c.FormFile(\"file\")]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	//_, err = os.Contextreate("files/applications/products_template.xlsx")
	//if err != nil {
	//	log.Println("[controller.ConvertExcelToStruct]|[os.Create]| error is: ", err.Error())
	//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
	//	return
	//}

	if err := c.SaveUploadedFile(img, "files/individ/individ.pdf"); err != nil {
		log.Println("[controller.ConvertExcelToStruct]|[c.SaveUploadedFile]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "файл успешно загружен"})
}

func GetIndividContract(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/pdf")
	c.File("files/individ/individ.pdf")
}

func Notification(c *gin.Context) {
	//TODO: ты должен сделать select -> типа выбрать дату договора, где будет время: и сделать сравнение мол если
	// 100 - изначальный договор
	// 60 осталось
	// Вопрос как сделать проверку?
	// Варианты:
	// 1. проверка по дню
	// 2. + 60 дней добавить нынешнему дню
	// 3. после чего все договора возьмем и вернем слайс
	// 4. после чего пройдемся по массиву и там же будет проверка даты
	// 5. и если он нашел какую-то дату то все отправляем уведомляем

	//TODO: Шаги
	// 2.
	// 1. вызов логики

}

// SearchByNumber godoc
// @Summary      Search Contract by Number
// @Description  add by json account
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        contract_number	 query string  true  "contract_number"
// @Param        status	 query string  true  "status"
// @Success      200      {object}  models.SearchContract
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /search_contract/ [get]
func SearchContractByNumber(c *gin.Context) {
	//TODO: цель взять не только договора по статусам
	// но и все договора:
	// можно сделать так, если мы получили  ACTIVE_AND_EXPIRED -  можем пустой статус отправить
	//

	status := c.Query("status")
	contractNumber := c.Query("contract_number")
	switch status {
	case "DRAFT":
		status = "черновик"
	case "ON_APPROVAL":
		status = "на согласовании"
	case "ACTIVE":
		status = "в работе"
	case "EXPIRED":
		status = "заверщённый"
	case "CANCELED":
		status = "отменен"
	case "ACTIVE_AND_EXPIRED":
		status = ""
	}
	log.Println(status, "checking status")

	result, err := service.SearchContractByNumber(contractNumber, status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}

// SearchContractDC godoc
// @Summary      Search Contract by contract_number, author, beneficiary
// @Description  поиск либо по одним из параметров - contract_number, author, beneficiary:
// @Description Примеры:
// @Description 1. target=contract_number&param=00001
// @Description 2. target=author&param=Иван
// @Description 3. target=beneficiary&param=ТОО «AK NIET GROUP
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        id   path     string  true  "id"
// @Param        target   query     string  true  "target"
// @Param        param    query     string  true  "target"
// @Success      200      {object}  models.SearchContract
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /search_history/{id} [get]
func SearchContractDC(c *gin.Context) {
	target := c.Query("target")
	param := c.Query("param")
	id := c.Param("id")

	log.Println(id, "добавить потом ID  в аргументах")

	result, err := service.SearchContractHistory(target, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}

// ChangeDataContract godoc
// @Summary     Продлить дату окончание договора
// @Description  Продлить дату окончание договора по ID и у которого статус в работе
// @Description Пример:
// @Description  change_date_contract/?extend_contract=true&id=163
// @Tags         contracts
// @Accept       json
// @Produce      json
// @Param        id    query     string  true  "id"
// @Success      200      {object}  interface{}
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /change_date_contract/ [get]
func ChangeDataContract(c *gin.Context) {
	id := c.Query("id")
	//extendContract := c.Query("is_extend_contract")
	//extendContractBool, err := strconv.ParseBool(extendContract)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//fmt.Println(extendContractBool, "extendContractBool")

	convertID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = service.ChangeDataContract(convertID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(200, gin.H{"reason": "договор успешно продлён!"})

}

// GetCountries godoc
// @Summary     Получаем список стран
// @Description  Получаем список стран
// @Tags         country
// @Accept       json
// @Produce      json
// @Success      200      {object}  models.Country
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /country/ [get]
func GetCountries(c *gin.Context) {

	countries, err := service.GetCountries()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
	}

	c.JSON(http.StatusOK, countries)

}

// SearchHistoryExecution godoc
// @Summary      Search History Execution by contract_number, author, beneficiary
// @Description  поиск либо по одним из параметров - contract_number, author, beneficiary:
// @Description Примеры:
// @Description 1. target=contract_number&param=00001
// @Description 2. target=author&param=Иван
// @Description 3. target=beneficiary&param=ТОО «AK NIET GROUP
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        id   path     string  true  "id"
// @Param        target   query     string  true  "target"
// @Param        param    query     string  true  "target"
// @Success      200      {object}  models.SearchContract
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /search_history_ex/{id} [get]
func SearchHistoryExecution(c *gin.Context) {
	target := c.Query("target")
	param := c.Query("param")
	result, err := service.SearchHistoryExecution(target, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}

// GetPriceType godoc
// @Summary     get price type by BIN
// @Description  get price type by BIN
// @Tags         price_type
// @Accept       json
// @Produce      json
// @Param        client_bin  body     models.BinPriceType  true  "client_bin"
// @Success      200      {object}  models.RespPriceType
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /price_type/ [post]
func GetPriceType(c *gin.Context) {
	var payload models.CodePriceType

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	priceType, err := service.GetPriceType(payload.ClientCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	c.JSON(http.StatusOK, priceType)
}

// CreatePriceType godoc
// @Summary     get price type by BIN
// @Description  get price type by BIN
// @Tags         price_type
// @Accept       json
// @Produce      json
// @Param        payload  body     models.PriceTypeCreate  true  "payload"
// @Success      200      {object}  models.PriceTypeResponse
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /create_price_type/ [post]
func CreatePriceType(c *gin.Context) {
	var payload models.PriceTypeCreate

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	payload.PricetypeCurrency = "398"
	payload.ClientBin = payload.ClientCode
	priceTypeResponse, err := service.CreatePriceType(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	c.JSON(http.StatusOK, priceTypeResponse)

}

// GetCurrencies godoc
// @Summary     get currencies
// @Description  get currencies
// @Tags         price_type
// @Accept       json
// @Produce      json
// @Success      200      {object}  []models.ConvertCurrency
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /currencies [get]
func GetCurrencies(c *gin.Context) {
	currencies, err := service.GetCurrencies()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, currencies)

}

func CheckContractIn1C(c *gin.Context) {

	var contracts models.ResponseContractFrom1C

	//var payload models.BinPriceType
	bin := c.Query("bin")
	contractType := c.Query("type")

	//err := c.ShouldBind(&payload)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err)
	//	return
	//}
	respContract1C, err := service.CheckContractIn1C(bin)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	for _, value := range respContract1C.ContractArr {
		if value.ContractType == contractType {
			contracts.ContractArr = append(contracts.ContractArr, value)
		}
	}

	c.JSON(200, contracts)

}

func GetSuppliers(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	suppliers, err := repository.GetSuppliersByParameter(field, value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	//suppliers, err := service.GetListSuppliersFrom1C()

	c.JSON(http.StatusOK, suppliers)
}

// GetProducts godoc
// @Summary     get products by key
// @Description  get products by key
// @Description   пример (typeValue=selectByPartName, typeParameters=аспири)
// @Description   пример (typeValue=selectByProductArrCode, typeParameters=00000026167)
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        payload  body     models.PayloadProduct  true  "payload"
// @Success      200      {object}  []models.ProductsData
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /products [post]
func GetProducts(c *gin.Context) {
	//TODO: add two key
	var payload models.PayloadProduct
	err := c.ShouldBind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("REQUEST", payload)

	products, err := service.GetListProductsFrom1C(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}

	c.JSON(http.StatusOK, products)

}

func SaveSuppliers(c *gin.Context) {
	suppliers, err := service.GetSuppliers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}

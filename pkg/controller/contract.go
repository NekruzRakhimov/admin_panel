package controller

import (
	"admin_panel/model"
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
// @Param  contract  body model.Contract true "creating contract"
// @Param  type  query string true "type of contract"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/{type} [post]
func CreateContract(c *gin.Context) {
	var contract model.Contract

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
// @Param  contract  body model.Contract true "creating contract"
// @Param  id  path string true "id договора на основе которого создаётся ДС"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/additional_agreement/{id} [post]
func AddAdditionalAgreement(c *gin.Context) {
	var contract model.Contract
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
// @Param  contract  body model.Contract true "editing contract"
// @Param  id  path string true "id of contract"
// @Param  type  path string true "type of contract"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/{type}/{id} [put]
func EditContract(c *gin.Context) {
	var contract model.Contract
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
// @Success 200 {array}  model.ContractMiniInfo
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
// @Success 200 {object}  model.Contract
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

	if err := c.SaveUploadedFile(img, "files/applications/products_template.xlsx"); err != nil {
		log.Println("[controller.ConvertExcelToStruct]|[c.SaveUploadedFile]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	f, err := excelize.OpenFile("files/applications/products_template.xlsx")
	//c.JSON(http.StatusOK, gin.H{"reason": "ok"})
	if err != nil {
		log.Println("[controller.ConvertExcelToStruct]|[excelize.OpenFile]| error is: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	var products []model.Product
	counter := 2
	for {
		var product model.Product
		product.ProductNumber, err = f.GetCellValue("page1", fmt.Sprintf("A%d", counter))
		if err != nil {
			log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		if product.ProductNumber == "" {
			break
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

		product.Price, err = strconv.ParseFloat(priceStr, 2)
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

		product.Substance, err = f.GetCellValue("page1", fmt.Sprintf("E%d", counter))
		if err != nil {
			log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		product.StorageCondition, err = f.GetCellValue("page1", fmt.Sprintf("F%d", counter))
		if err != nil {
			log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}

		product.Producer, err = f.GetCellValue("page1", fmt.Sprintf("G%d", counter))
		if err != nil {
			log.Println("[controller.ConvertExcelToStruct]|[f.GetCellValue]| error is: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
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
// @Success 200 {array}  model.Counterparty
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /counterparty/{client} [get]
func CounterpartyContract(c *gin.Context) {
	binClient := c.Param("client")

	//organization := c.Query("organization")


	contract, err := service.CounterpartyContract(binClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data:": contract})

}

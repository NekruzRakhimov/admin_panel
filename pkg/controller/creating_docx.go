package controller

import (
	"admin_panel/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
	"log"
	"net/http"
	"strconv"
	"time"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(apiKey)
	if err != nil {
		panic(err)
	}
}

//FormContract contract godoc
// @Summary Forming contract
// @Description Forming contract
// @Accept  json
// @Produce  json
// @Tags contracts
// @Param  contract  body models.Contract true "forming contract"
// @Param  with_temp_conditions  param string true "with temperature conditions"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/{type} [post]

func FormContract(c *gin.Context) {
	var contract models.Contract

	if err := c.BindJSON(&contract); err != nil {
		log.Println("[controller.FormContract]|[c.BindJSO]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	contract.Type = c.Param("contract_type")

	withTempConditions, err := strconv.ParseBool(c.Param("with_temp_conditions"))
	if err != nil {
		log.Println("[controller.FormContract]|[strconv.ParseBool]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	switch contract.Type {
	case "supply":
		SupplyContract(c, contract, withTempConditions)
	case "marketing_services":
		MarketingServiceContract(c, contract)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"reason": "некорректный тип договора"})
		return
	}
}

func SupplyContract(c *gin.Context, contract models.Contract, withTempConditions bool) {
	if withTempConditions {
		SupplyContractWithTempConditions(c, contract)
	} else {
		SupplyContractWithoutTempConditions(c, contract)
	}
}

func SupplyContractWithTempConditions(c *gin.Context, contract models.Contract) {
	doc, err := document.Open("files/contracts/supply/without_temp_cond.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
	}

	// This sample document uses structured document tags, which are not common
	// except for in document templates.  Normally you can just iterate over the
	// document's paragraphs.
	for _, sdt := range doc.StructuredDocumentTags() {
		for _, p := range sdt.Paragraphs() {
			paragraphs = append(paragraphs, p)
		}
	}

	for _, p := range paragraphs {
		for _, r := range p.Runs() {
			switch r.Text() {
			case "№ CONTRACT_NUMBER":
				// ClearContent clears both text and line breaks within a run,
				// so we need to add the line break back
				r.ClearContent()
				r.AddText(contract.ContractParameters.ContractNumber)
				r.AddBreak()

			//para := doc.InsertParagraphBefore(p)
			//para.AddRun().AddText("Mr.")
			//para.SetStyle("Name") // Name is a default style in this template file

			//para = doc.InsertParagraphAfter(p)
			//para.AddRun().AddText("III")
			//para.SetStyle("Name")

			case "NEKRUZ": //CONTRACT_NUMBER
				r.ClearContent()
				r.AddText(contract.ContractParameters.ContractNumber)
			//case "MANAGER":
			//	r.ClearContent()
			//	r.AddText(contract.Manager)
			case "BENEFICIARY":
				r.ClearContent()
				r.AddText(contract.Requisites.Beneficiary)
			case "ADDRESS":
				r.ClearContent()
				r.AddText(contract.ContractParameters.ContractDate)
			case "BENEFICIARY_BANK_ADDRESS":
				r.ClearContent()
				r.AddText("CUSTOM_BENEFICIARY_BANK_ADDRESS")
			case "BENEFICIARY_BANK":
				r.ClearContent()
				r.AddText(contract.Requisites.BankOfBeneficiary)
			case "SWIFT_CODE":
				r.ClearContent()
				r.AddText(contract.Requisites.SwiftCode)
			case "ACCOUNT":
				r.ClearContent()
				r.AddText(contract.Requisites.AccountNumber)
			case "AZIZ": // END-DATE
				r.ClearContent()
				r.AddText(contract.ContractParameters.EndDate)
			case "AMOUNT":
				r.ClearContent()
				r.AddText(fmt.Sprintf("%f", contract.ContractParameters.ContractAmount))
			case "INTERVAL": // DeliveryTimeInterval
				r.ClearContent()
				r.AddText(fmt.Sprintf("%d", contract.ContractParameters.DeliveryTimeInterval))
			case "DELIVERY_DATE": // DeliveryTimeInterval
				r.ClearContent()
				r.AddText(contract.ContractParameters.DateOfDelivery)
			//case "PRE_PAYMENT":
			//r.ClearContent()
			//r.AddText(fmt.Sprintf("%f", contract.ContractParameters.Prepayment))
			case "RETURNTIME": // ReturnTimeDelivery
				r.ClearContent()
				r.AddText(fmt.Sprintf("%d", contract.ContractParameters.ReturnTimeDelivery))
			case "DELIVERIES": // DATE_OF_DELIVERY
				r.ClearContent()
				r.AddText(contract.ContractParameters.DateOfDelivery)
			case "TABLE_PLACE":
				// First Table
				r.ClearContent()

				//paragraph := doc.InsertParagraphAfter(p)
				//paragraph.AddRun().AddText("")

				table := doc.InsertTableAfter(p)
				// width of the page
				table.Properties().SetWidthPercent(100)
				// with thick borers
				borders := table.Properties().Borders()
				borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

				row := table.AddRow()
				run := row.AddCell().AddParagraph().AddRun()
				run.AddText("№")
				row.AddCell().AddParagraph().AddRun().AddText("Торговое название / Trade Name ")
				row.AddCell().AddParagraph().AddRun().AddText("ТЦена, / Price, CIP Almaty")
				row.AddCell().AddParagraph().AddRun().AddText("Состав. Характеристика ")
				row.AddCell().AddParagraph().AddRun().AddText("Условия хранения ")
				row.AddCell().AddParagraph().AddRun().AddText("Производитель ")
				run.Properties().SetHighlight(wml.ST_HighlightColorYellow)

				for _, product := range contract.Products {
					row = table.AddRow()
					row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%s", product.ProductNumber))
					row.AddCell().AddParagraph().AddRun().AddText(product.ProductName)
					row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%f", product.Price))
					row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%s", product.Substance))
					row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%s", product.StorageCondition))
					row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%s", product.Producer))
				}

			default:
				fmt.Println("not modifying", r.Text())
			}
		}
	}

	doc.SaveToFile("files/contracts/edit-document.docx")

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	c.File("files/contracts/edit-document.docx")
}

func SupplyContractWithoutTempConditions(c *gin.Context, contract models.Contract) {
	doc, err := document.Open("files/contracts/supply/without_temp_cond.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
	}

	// This sample document uses structured document tags, which are not common
	// except for in document templates.  Normally you can just iterate over the
	// document's paragraphs.
	for _, sdt := range doc.StructuredDocumentTags() {
		for _, p := range sdt.Paragraphs() {
			paragraphs = append(paragraphs, p)
		}
	}

	for _, p := range paragraphs {
		for _, r := range p.Runs() {
			switch r.Text() {
			case "№ CONTRACT_NUMBER":
				// ClearContent clears both text and line breaks within a run,
				// so we need to add the line break back
				r.ClearContent()
				r.AddText(contract.ContractParameters.ContractNumber)
				r.AddBreak()

			//para := doc.InsertParagraphBefore(p)
			//para.AddRun().AddText("Mr.")
			//para.SetStyle("Name") // Name is a default style in this template file

			//para = doc.InsertParagraphAfter(p)
			//para.AddRun().AddText("III")
			//para.SetStyle("Name")

			case "NEKRUZ": //CONTRACT_NUMBER
				r.ClearContent()
				r.AddText(contract.ContractParameters.ContractNumber)
			//case "MANAGER":
			//	r.ClearContent()
			//	r.AddText(contract.Manager)
			case "BENEFICIARY":
				r.ClearContent()
				r.AddText(contract.Requisites.Beneficiary)
			case "ADDRESS":
				r.ClearContent()
				r.AddText(contract.ContractParameters.ContractDate)
			case "BENEFICIARY_BANK_ADDRESS":
				r.ClearContent()
				r.AddText("CUSTOM_BENEFICIARY_BANK_ADDRESS")
			case "BENEFICIARY_BANK":
				r.ClearContent()
				r.AddText(contract.Requisites.BankOfBeneficiary)
			case "SWIFT_CODE":
				r.ClearContent()
				r.AddText(contract.Requisites.SwiftCode)
			case "ACCOUNT":
				r.ClearContent()
				r.AddText(contract.Requisites.AccountNumber)
			case "AZIZ": // END-DATE
				r.ClearContent()
				r.AddText(contract.ContractParameters.EndDate)
			case "AMOUNT":
				r.ClearContent()
				r.AddText(fmt.Sprintf("%f", contract.ContractParameters.ContractAmount))
			case "INTERVAL": // DeliveryTimeInterval
				r.ClearContent()
				r.AddText(fmt.Sprintf("%d", contract.ContractParameters.DeliveryTimeInterval))
			case "DELIVERY_DATE": // DeliveryTimeInterval
				r.ClearContent()
				r.AddText(contract.ContractParameters.DateOfDelivery)
			//case "PRE_PAYMENT":
			//r.ClearContent()
			//r.AddText(fmt.Sprintf("%f", contract.ContractParameters.Prepayment))
			case "RETURNTIME": // ReturnTimeDelivery
				r.ClearContent()
				r.AddText(fmt.Sprintf("%d", contract.ContractParameters.ReturnTimeDelivery))
			case "DELIVERIES": // DATE_OF_DELIVERY
				r.ClearContent()
				r.AddText(contract.ContractParameters.DateOfDelivery)
			case "TABLE_PLACE":
				// First Table
				r.ClearContent()

				//paragraph := doc.InsertParagraphAfter(p)
				//paragraph.AddRun().AddText("")

				table := doc.InsertTableAfter(p)
				// width of the page
				table.Properties().SetWidthPercent(100)
				// with thick borers
				borders := table.Properties().Borders()
				borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

				row := table.AddRow()
				run := row.AddCell().AddParagraph().AddRun()
				run.AddText("№")
				row.AddCell().AddParagraph().AddRun().AddText("Торговое название / Trade Name ")
				row.AddCell().AddParagraph().AddRun().AddText("ТЦена / Price, CIP Almaty")
				run.Properties().SetHighlight(wml.ST_HighlightColorYellow)

				for _, product := range contract.Products {
					row = table.AddRow()
					row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%s", product.ProductNumber))
					row.AddCell().AddParagraph().AddRun().AddText(product.ProductName)
					row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%f", product.Price))
				}

			default:
				fmt.Println("not modifying", r.Text())
			}
		}
	}

	doc.SaveToFile("files/contracts/edit-document.docx")

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	c.File("files/contracts/edit-document.docx")
}

func MarketingServiceContract(c *gin.Context, contract models.Contract) {
	doc, err := document.Open("files/contracts/marketing_service/marketing_service.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	defer doc.Close()

	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
	}

	// This sample document uses structured document tags, which are not common
	// except for in document templates.  Normally you can just iterate over the
	// document's paragraphs.
	for _, sdt := range doc.StructuredDocumentTags() {
		for _, p := range sdt.Paragraphs() {
			paragraphs = append(paragraphs, p)
		}
	}

	for _, p := range paragraphs {
		for _, r := range p.Runs() {
			//fmt.Println("")
			//fmt.Println(">>>>>>>")
			//fmt.Println(r.Text())
			//fmt.Println(">>>>>>>")
			//fmt.Println("")
			switch r.Text() {
			case "NEKRUZ": //CONTRACT_NUMBER
				r.ClearContent()
				r.AddText(contract.ContractParameters.ContractNumber)
			case "BENEFICIARY":
				r.ClearContent()
				r.AddText(contract.Requisites.Beneficiary)
			case "Date":
				r.ClearContent()
				r.AddText(time.Now().Format("Jan 2, 2006"))
			case "ADDRESS":
				r.ClearContent()
				if len(contract.ContractParameters.DeliveryAddress) > 0 {
					r.AddText(contract.ContractParameters.DeliveryAddress[0])
				}
			case "BANK_BENEFICIARY":
				r.ClearContent()
				r.AddText(contract.Requisites.BankOfBeneficiary)
			case "BIC":
				r.ClearContent()
				r.AddText(contract.Requisites.SwiftCode)
			case "IIC":
				r.ClearContent()
				r.AddText(contract.Requisites.IIC)
			case "BIN":
				r.ClearContent()
				r.AddText(contract.Requisites.BIN)
			case "RNN":
				r.ClearContent()
				r.AddText(contract.Requisites.AccountNumber)
			case "PHONE":
				r.ClearContent()
				r.AddText(contract.SupplierCompanyManager.Phone)
			case "MANAGER":
				r.ClearContent()
				r.AddText(contract.SupplierCompanyManager.Position + " " + contract.SupplierCompanyManager.FullName)
			case "BASE":
				r.ClearContent()
				r.AddText(contract.SupplierCompanyManager.Base)
			default:
				fmt.Println("not modifying", r.Text())
			}
		}
	}

	doc.SaveToFile("files/contracts/edit-document.docx")

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	c.File("files/contracts/edit-document.docx")
}

const (
	apiKey = `4819ce4158e078898d2209c9cb83f40e894dcdc68c0b8a5eb792ec2008534334`
)

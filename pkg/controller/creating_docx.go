package controller

import (
	"admin_panel/model"
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
// @Param  contract  body model.Contract true "forming contract"
// @Param  with_temp_conditions  param string true "with temperature conditions"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contract/{type} [post]

func FormContract(c *gin.Context) {
	var contract model.Contract

	if err := c.BindJSON(&contract); err != nil {
		log.Println("[controller.FormContract]|[c.BindJSO]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	contract.Type = c.Param("type")

	withTempConditions, err := strconv.ParseBool(c.Param("with_temp_conditions"))
	if err != nil {
		log.Println("[controller.FormContract]|[strconv.ParseBool]| error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if withTempConditions {
		c.File("files/contracts/with_temp_conditions")
	} else {
		WithoutTempConditions(c, contract)
	}
}

func WithoutTempConditions(c *gin.Context, contract model.Contract) {
	doc, err := document.Open("files/contracts/without_temp_cond.docx")
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

			case "MANAGER":
				r.ClearContent()
				r.AddText(contract.Manager)
			case "BENEFICIARY":
				r.ClearContent()
				r.AddText(contract.Requisites.Beneficiary)
			case "ADDRESS":
				r.ClearContent()
				r.AddText("CIP- Красногвардейский Тракт (ул. Суюнбая) 258, г. Алматы, Республика Казахстан, Инкотермс 2010")
			case "ACCOUNT":
				r.ClearContent()
				r.AddText(contract.Requisites.AccountNumber)
			case "CONTRACT_AMOUNT":
				r.ClearContent()
				r.AddText(fmt.Sprintf("%f", contract.ContractParameters.ContractAmount))
			case "PRE_PAYMENT":
				r.ClearContent()
				r.AddText(fmt.Sprintf("%f", contract.ContractParameters.Prepayment))
			case "BANK_BENEFICIARY":
				r.ClearContent()
				r.AddText(contract.Requisites.BankOfBeneficiary)
			case "TABLE_PLACE":
				// First Table
				r.ClearContent()

				paragraph := doc.AddParagraph()
				paragraph.AddRun().AddText("")

				table := doc.InsertTableAfter(paragraph)
				// width of the page
				table.Properties().SetWidthPercent(100)
				// with thick borers
				borders := table.Properties().Borders()
				borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

				row := table.AddRow()
				run := row.AddCell().AddParagraph().AddRun()
				run.AddText("№")
				row.AddCell().AddParagraph().AddRun().AddText("Торговое название / Trade Name ")
				row.AddCell().AddParagraph().AddRun().AddText("ТЦена, CIP Алматы, в долларах США / Price, CIP Almaty, USD ")
				run.Properties().SetHighlight(wml.ST_HighlightColorYellow)

				for _, product := range contract.Products {
					row = table.AddRow()
					row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%d", product.ProductNumber))
					row.AddCell().AddParagraph().AddRun().AddText(product.ProductName)
					row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%f", product.Price))
				}

			//case "BENEFICIARY":
			//	r.ClearContent()
			//	r.AddText("TOO TEST")
			//	r.AddBreak()
			//case "BENEFICIARY_BOTTOM":
			//	r.ClearContent()
			//	r.AddText("TOO TEST")
			//	r.AddBreak()
			//case "Title":
			//	// we remove the title content entirely
			//	p.RemoveRun(r)
			//case "Company":
			//	r.ClearContent()
			//	r.AddText("Smith Enterprises")
			//	r.AddBreak()
			//case "Address":
			//	r.ClearContent()
			//	r.AddText("112 Rustic Rd")
			//	r.AddBreak()
			//case "City, ST ZIP Code":
			//	r.ClearContent()
			//	r.AddText("San Francisco, CA 94016")
			//	r.AddBreak()
			//case "Dear Recipient:":
			//	r.ClearContent()
			//	r.AddText("Dear Mrs. Smith:")
			//	r.AddBreak()
			//case "Your Name":
			//	r.ClearContent()
			//	r.AddText("John Smith")
			//	r.AddBreak()
			//
			//	run := p.InsertRunBefore(r)
			//	run.AddText("---Before----")
			//	run.AddBreak()
			//	run = p.InsertRunAfter(r)
			//	run.AddText("---After----")

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

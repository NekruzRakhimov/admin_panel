package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gopkg.in/gomail.v2"
	"log"
)

func CreateSegment(segment models.Segment) error {

	product, err := json.Marshal(segment.Products)
	if err != nil {
		return err
	}
	segment.ProductStr = string(product)

	region, err := json.Marshal(segment.Region)
	if err != nil {
		return err
	}
	segment.RegionStr = string(region)

	return repository.CreateSegment(segment)

}

//func GetSegmentByID(id int) error {
//
//	product, err := json.Marshal(segment.Products)
//	if err != nil {
//		return err
//	}
//	segment.ProductStr = string(product)
//
//	region, err := json.Marshal(segment.Region)
//	if err != nil {
//		return err
//	}
//	segment.RegionStr = string(region)
//
//	return repository.CreateSegment(segment)
//
//}

//func GetGraphicByID(c *gin.Context) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"reason": "ERROR"})
//		return
//	}
//
//	graphic, err := service.GetGraphicByID(id)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, graphic)
//}

func GetSegmentByID(id int) (models.Segment, error) {
	segment, err := repository.GetSegmentByID(id)
	if err != nil {
		return segment, err
	}

	err = json.Unmarshal([]byte(segment.ProductStr), &segment.Products)
	if err != nil {
		return segment, err
	}

	err = json.Unmarshal([]byte(segment.RegionStr), &segment.Region)
	if err != nil {
		return segment, err
	}

	//SendNotificationSegment("НАДО УКАЗАТЬ ПУТЬ или сперва заполнить Эксель")
	//
	return segment, nil

}

//func GetSegmentByID(id int) (models.Segment, error) {
//	segment, err := repository.GetSegmentByID(id)
//	if err != nil {
//		return segment, err
//	}
//
//	err = json.Unmarshal([]byte(segment.ProductStr), &segment.Products)
//	if err != nil {
//		return segment, err
//	}
//
//	err = json.Unmarshal([]byte(segment.RegionStr), &segment.Region)
//	if err != nil {
//		return segment, err
//	}
//
//	//SendNotificationSegment("НАДО УКАЗАТЬ ПУТЬ или сперва заполнить Эксель")
//	//
//	return segment, nil
//
//}

func GetSegments() ([]models.Segment, error) {
	segments, err := repository.GetSegments()
	if err != nil {
		return nil, err
	}
	var segmentsSL []models.Segment

	for _, segment := range segments {
		var seg models.Segment
		seg.ID = segment.ID
		seg.SegmentCode = segment.SegmentCode
		seg.NameSegment = segment.NameSegment
		seg.Beneficiary = segment.Beneficiary
		seg.Bin = segment.Bin
		seg.ClientCode = segment.ClientCode
		seg.Email = segment.Email
		seg.ForMarket = segment.ForMarket

		err = json.Unmarshal([]byte(segment.ProductStr), &seg.Products)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(segment.RegionStr), &seg.Region)
		if err != nil {
			return nil, err
		}
		segmentsSL = append(segmentsSL, seg)

	}

	//SendNotificationSegment("НАДО УКАЗАТЬ ПУТЬ или сперва заполнить Эксель")
	//
	return segmentsSL, nil

}

func GetSegment(supplier string) (models.Segment, error) {
	segment, err := repository.GetSegment(supplier)
	if err != nil {
		return segment, err
	}

	err = json.Unmarshal([]byte(segment.ProductStr), &segment.Products)
	if err != nil {
		return segment, err
	}

	err = json.Unmarshal([]byte(segment.RegionStr), &segment.Region)
	if err != nil {
		return segment, err
	}

	//SendNotificationSegment("НАДО УКАЗАТЬ ПУТЬ или сперва заполнить Эксель")
	//
	return segment, nil

}

func DeleteSegmentByID(id int) error {
	return repository.DeleteSegmentByID(id)
}

const segment = "сегменты"

func FillSegment(graphic models.FormedGraphic, products []models.FormedGraphicProduct, graphicAnother models.Graphic) {

	f := excelize.NewFile()
	f.NewSheet(segment)
	//ineration := 1
	f.SetCellValue(segment, "A3", "Поставщик:")
	f.SetCellValue(segment, "A5", "Склад:")
	f.SetCellValue(segment, "A7", "Регион:")
	//fmt.Println("graphicAnother.RegionName", graphicAnother.RegionName)

	// часть товаров
	f.SetCellValue(segment, "A10", "№:")
	f.SetCellValue(segment, "B10", "Код:")
	f.SetCellValue(segment, "C10", "Товар:")
	f.SetCellValue(segment, "D10", "Штрихкод:")
	f.SetCellValue(segment, "E10", "Кол-во:")
	f.SetCellValue(segment, "F10", "Ед.:")
	f.SetCellValue(segment, "B3", graphic.Supplier)
	f.SetCellValue(segment, "B5", graphic.Store)
	f.SetCellValue(segment, "B7", graphicAnother.RegionName)
	//f.SetColWidth(segment, "A", "A", 42)

	f.SetColWidth(segment, "A", "F", 22)
	f.SetColWidth(segment, "C", "C", 58)
	f.SetColWidth(segment, "D", "D", 17)

	i := 10
	var id = 1
	for _, product := range products {
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "A", i+1), id)
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "B", i+1), product.ProductCode)
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "C", i+1), product.ProductName)
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "D", i+1), product.StoreCode)
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "E", i+1), product.SalesCount)
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "F", i+1), "шт")
		i++
		id++

	}

	f.DeleteSheet("Sheet1")
	err := f.SaveAs("files/segments/segment.xlsx")
	if err != nil {
		log.Println(err)
	}
}

func SendNotificationSegment(path string, email string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "aziz.rahimov0001@gmail.com")

	// Set E-Mail receivers
	//m.SetHeader("To", "aziz.rahimov0001@gmail.com")
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Segments")

	// Set E-Mail body. You can set plain text or html with text/html
	//m.SetBody("text/plain", "This is Gomail test body")
	m.Attach(path)

	// Settings for SMTP server
	//d := gomail.NewDialer("smtp.gmail.com", 587, "thief65mk@gmail.com", "Aziz65mk")
	d := gomail.NewDialer("smtp.gmail.com", 587, "aziz.rahimov0001@gmail.com", "yknixmmoyledqfxn")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		//panic(err)
	}
	fmt.Println("successfully sent email!")
	return

}

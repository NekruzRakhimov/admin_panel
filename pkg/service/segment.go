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

func ChangeSegment(segment models.Segment) error {

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

	return repository.ChangeSegment(segment)

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

func ChangeLetter(id int) error {
	return repository.ChangeLetter(id)

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

	styleLargeText, _ := f.NewStyle(&excelize.Style{

		Font: &excelize.Font{
			Bold:      true,
			Italic:    false,
			Underline: "single",
			Family:    "Arial",
			Size:      14,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "left",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})

	styleTopBorders, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 5,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      true,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      9,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})

	styleBottomBorder, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 5,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      8,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})

	styleBorderLeft, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 5,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      8,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})

	styleBorderRight, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 5,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      8,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})

	styleButtonBorderUpLeft, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 5,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 5,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      8,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})

	styleButtonBorderUpRight, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 5,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 5,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      8,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})

	styleBorderDownLeft, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 5,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 5,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      8,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})

	styleBBorderDownRight, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 5,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 5,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      8,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})

	styleBorderCenter, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      8,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})
	styleBoldCenter, _ := f.NewStyle(&excelize.Style{

		Font: &excelize.Font{
			Bold:      true,
			Italic:    false,
			Underline: "",
			Family:    "Arial",
			Size:      9,
			Strike:    false,
			Color:     "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "left",
			Indent:          1,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})
	log.Println("GRAPHIC", graphic)

	f.NewSheet(segment)
	//ineration := 1
	f.MergeCell(segment, "B1", "AI1")
	f.SetCellValue(segment, "B1", "Заказ № (Тестовый заказ) от 01.01.2021")
	f.SetCellStyle(segment, "B1", "AI1", styleLargeText)

	f.SetCellValue(segment, "B2", "Поставщик:")
	f.MergeCell(segment, "B2", "E2")
	f.SetCellValue(segment, "F2", graphic.Supplier)
	f.MergeCell(segment, "F2", "AI2")
	f.SetCellStyle(segment, "F2", "AI2", styleBoldCenter)

	f.SetCellValue(segment, "B4", "Покупатель:")
	f.MergeCell(segment, "B4", "E4")

	f.SetCellValue(segment, "F4", "Тестовый покупатель")
	f.MergeCell(segment, "F4", "AI4")
	f.SetCellStyle(segment, "F4", "AI4", styleBoldCenter)

	f.SetCellValue(segment, "B6", "Договор:")
	f.MergeCell(segment, "B6", "F6")
	f.SetCellValue(segment, "B8", "Дата поставки:")
	f.MergeCell(segment, "B8", "F8")
	f.SetCellValue(segment, "B10", "Склад:")
	f.MergeCell(segment, "B10", "F10")
	f.SetCellValue(segment, "B12", "Менеджер:")
	f.MergeCell(segment, "B12", "F12")

	f.SetCellValue(segment, "G6", "Тестовый договор:")
	f.MergeCell(segment, "G6", "U6")
	f.SetCellValue(segment, "G8", "Тестовая дата поставки:")
	f.MergeCell(segment, "G8", "U8")
	f.SetCellValue(segment, "G10", graphic.Store)
	f.MergeCell(segment, "G10", "U10")
	f.SetCellValue(segment, "G12", "Тестовый менеджер")
	f.MergeCell(segment, "G12", "U12")
	//fmt.Println("graphicAnother.RegionName", graphicAnother.RegionName)
	f.SetCellValue(segment, "W6", "Валюта заказа:")
	f.MergeCell(segment, "W6", "AA6")
	f.SetCellValue(segment, "W8", "Вид транспорта:")
	f.MergeCell(segment, "W8", "AA8")

	f.SetCellValue(segment, "AB6", "KZT (тествоая)")
	f.MergeCell(segment, "AB6", "AH6")
	f.SetCellValue(segment, "AB8", "тестовый транспорт")
	f.MergeCell(segment, "AB8", "AH8")

	// часть товаров
	f.SetCellValue(segment, "B14", "№")
	f.MergeCell(segment, "B14", "C15")

	f.SetCellValue(segment, "D14", "Код")
	f.MergeCell(segment, "D14", "G15")

	f.SetCellValue(segment, "H14", "Товар")
	f.MergeCell(segment, "H14", "T15")

	f.SetCellValue(segment, "U14", "Штрихкод")
	f.MergeCell(segment, "U14", "U15")

	f.SetCellValue(segment, "V14", "Производитель")
	f.MergeCell(segment, "V14", "V15")

	f.SetCellValue(segment, "W14", "Кол-во")
	f.MergeCell(segment, "W14", "Y15")

	f.SetCellValue(segment, "Z14", "Ед.")
	f.MergeCell(segment, "Z14", "AA15")

	f.SetCellValue(segment, "AB14", "Закуп. Цена")
	f.MergeCell(segment, "AB14", "AE15")

	f.SetCellValue(segment, "AF14", "Сумма")
	f.MergeCell(segment, "AF14", "AI15")

	f.SetCellValue(segment, "AJ14", "Лот")
	f.MergeCell(segment, "AJ14", "AM15")

	//f.SetColWidth(segment, "A", "A", 42)
	f.SetColWidth(segment, "A", "T", 3)
	f.SetColWidth(segment, "U", "V", 21)
	f.SetColWidth(segment, "W", "AI", 3.8)
	f.SetColWidth(segment, "AJ", "BP", 2.7)

	f.SetRowHeight(segment, 1, 31)
	f.SetRowHeight(segment, 2, 25)
	f.SetRowHeight(segment, 3, 6.8)
	f.SetRowHeight(segment, 4, 24.8)
	f.SetRowHeight(segment, 5, 6.8)
	f.SetRowHeight(segment, 6, 12.8)
	f.SetRowHeight(segment, 7, 6.8)
	f.SetRowHeight(segment, 8, 12.8)
	f.SetRowHeight(segment, 9, 6.8)
	f.SetRowHeight(segment, 10, 12.8)
	f.SetRowHeight(segment, 11, 6.8)
	f.SetRowHeight(segment, 12, 12.8)
	f.SetRowHeight(segment, 13, 6.8)
	f.SetRowHeight(segment, 8, 12.8)

	f.SetCellStyle(segment, "B14", "AM14", styleTopBorders)
	f.SetCellStyle(segment, "B14", "C15", styleButtonBorderUpLeft)
	f.SetCellStyle(segment, "AJ14", "AM15", styleButtonBorderUpRight)

	i := 16
	var id = 1
	var total float64
	for _, product := range products {
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "B", i), id)
		f.MergeCell(segment, "B"+fmt.Sprint(i), "C"+fmt.Sprint(i))
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "D", i), product.ProductCode)
		f.MergeCell(segment, "D"+fmt.Sprint(i), "G"+fmt.Sprint(i))
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "H", i), product.ProductName)
		f.MergeCell(segment, "H"+fmt.Sprint(i), "T"+fmt.Sprint(i))
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "U", i), product.StoreCode)
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "V", i), "Тестовый производитель")
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "W", i), product.SalesCount)
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "Z", i), "шт")
		f.MergeCell(segment, "Z"+fmt.Sprint(i), "AA"+fmt.Sprint(i))

		f.SetCellValue(segment, fmt.Sprintf("%s%d", "AB", i), 999) //цена
		f.MergeCell(segment, "AB"+fmt.Sprint(i), "AE"+fmt.Sprint(i))

		var sum = product.SalesCount * 999                        //находим сумму
		f.SetCellValue(segment, fmt.Sprintf("%s%d", "AF", i), sum) //сумма
		f.MergeCell(segment, "AF"+fmt.Sprint(i), "AI"+fmt.Sprint(i))

		f.SetCellValue(segment, fmt.Sprintf("%s%d", "AJ", i), "тестовый лот") //лот
		f.MergeCell(segment, "AJ"+fmt.Sprint(i), "AM"+fmt.Sprint(i))

		f.SetRowHeight(segment, i, 11.3)
		i++
		id++
		total+=sum
	}

	f.SetCellStyle(segment, fmt.Sprintf("%s%d", "D", i-1), fmt.Sprintf("%s%d", "AI", i-1), styleBottomBorder)
	f.SetCellStyle(segment, fmt.Sprintf("%s%d", "B", i-1), fmt.Sprintf("%s%d", "C", i-1), styleBorderDownLeft)
	f.SetCellStyle(segment, fmt.Sprintf("%s%d", "AJ", i-1), fmt.Sprintf("%s%d", "AM", i-1), styleBBorderDownRight)

	f.SetCellStyle(segment, fmt.Sprintf("%s%d", "B", 16), fmt.Sprintf("%s%d", "C", i-2), styleBorderLeft)
	f.SetCellStyle(segment, fmt.Sprintf("%s%d", "AJ", 16), fmt.Sprintf("%s%d", "AM", i-2), styleBorderRight)
	f.SetCellStyle(segment, fmt.Sprintf("%s%d", "D", 16), fmt.Sprintf("%s%d", "AI", i-2), styleBorderCenter)

	f.SetCellValue(segment, "AE" + fmt.Sprint(i+1), "Итого")
	f.SetCellValue(segment, "AF" + fmt.Sprint(i+1), total)

	f.SetCellValue(segment, "AE" + fmt.Sprint(i+2), "В том числе НДС:")
	f.SetCellValue(segment, "AF" + fmt.Sprint(i+1), total)
	f.SetCellValue(segment, "AF" + fmt.Sprint(i+2), total*0.12)

	f.MergeCell(segment, "AB" + fmt.Sprint(i+1), "AE" + fmt.Sprint(i+1))
	f.MergeCell(segment, "AB" + fmt.Sprint(i+2), "AE" + fmt.Sprint(i+2))
	f.MergeCell(segment, "AF" + fmt.Sprint(i+1), "AI" + fmt.Sprint(i+1))
	f.MergeCell(segment, "AF" + fmt.Sprint(i+2), "AI" + fmt.Sprint(i+2))
	f.SetCellStyle(segment, "AB" + fmt.Sprint(i+1), "AI" + fmt.Sprint(i+2), styleBoldCenter)



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

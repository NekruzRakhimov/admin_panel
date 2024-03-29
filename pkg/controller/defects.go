package controller

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"admin_panel/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
func GetDefectsByPharmacyLS(c *gin.Context) {
	date := struct {
		Date string `json:"date"`
	}{}

	if err := c.BindJSON(&date); err != nil {
		c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
		return
	}

	req := models.DefectsRequest{
		Startdate: date.Date,
		Enddate:   date.Date,
	}

	_, err := service.GetDefectsLS(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, res)

	//c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//c.File("./files/defects/defects_pharmacy.xlsx")

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("./files/defects/res_ls.xlsx")
}
*/

const TempData = " 00:00:00"

func GetSalesCount(c *gin.Context) {
	var req models.SalesCountRequest

	fmt.Println("Started Main")
	fmt.Println("Started unmarshalling req_body")
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	fmt.Println("Finished unmarshalling req_body")
	req.Startdate += TempData
	fmt.Println(req.Startdate)
	req.Enddate += TempData

	salesCount, err := service.GetSalesCountExt(req)
	if err != nil {
		fmt.Println(salesCount)
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	fmt.Println("Finished Main")
	c.JSON(http.StatusOK, salesCount)
}

func SaveMatrix(c *gin.Context) {
	//var matrix []models.MatrixInfoFrom1C
	//if err := c.BindJSON(&matrix); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
	//	return
	//}

	//if err := service.SaveAllMatrixFrom1C(matrix); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
	//	return
	//}

	stores, _ := repository.GetAllStores()

	c.JSON(http.StatusOK, stores)
}

func GetDefectsByPharmacyPF(c *gin.Context) {
	date := struct {
		Date string `json:"date"`
	}{}

	if err := c.BindJSON(&date); err != nil {
		c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
		return
	}

	req := models.DefectsRequest{
		Startdate: date.Date,
		Enddate:   date.Date,
		IsPF:      true,
	}
	log.Println(time.Now(), " Started Defects - Main")
	fmt.Println(time.Now(), " Started Defects - Main")
	mainTime := time.Now()
	err := service.GetNewDefectsPf(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, res)

	//c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//c.File("./files/defects/defects_pharmacy.xlsx")

	log.Println(time.Now(), " Finished Defects - Main: durance[", time.Now().Sub(mainTime), "]")
	fmt.Println(time.Now(), " Finished Defects - Main: durance[", time.Now().Sub(mainTime), "]")
	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File("./files/defects/res.xlsx")
}

func OrderDefectsPfReport(c *gin.Context) {
	date := struct {
		Date string `json:"date"`
	}{}

	if err := c.BindJSON(&date); err != nil {
		c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
		return
	}

	req := models.DefectsRequest{
		Startdate: date.Date,
		Enddate:   date.Date,
		IsPF:      true,
	}
	log.Println(time.Now(), " Started Defects - Main")
	fmt.Println(time.Now(), " Started Defects - Main")
	//mainTime := time.Now()

	go func() {
		err := service.OrderDefectsPF(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{"reason": "запрос на сформирование отчета принят. Статус: 'в процессе'"})

	//log.Println(time.Now(), " Finished Defects - Main: durance[", time.Now().Sub(mainTime), "]")
	//fmt.Println(time.Now(), " Finished Defects - Main: durance[", time.Now().Sub(mainTime), "]")
	//c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//c.File("./files/defects/res.xlsx")
}

func OrderDefectsLsReport(c *gin.Context) {
	date := struct {
		Date string `json:"date"`
	}{}

	if err := c.BindJSON(&date); err != nil {
		c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
		return
	}

	req := models.DefectsRequest{
		Startdate: date.Date,
		Enddate:   date.Date,
		IsPF:      true,
	}
	log.Println(time.Now(), " Started Defects - Main")
	fmt.Println(time.Now(), " Started Defects - Main")
	//mainTime := time.Now()

	go func() {
		err := service.OrderDefectsLs(req)
		if err != nil {
			//c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
			log.Println(err.Error())
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{"reason": "запрос на сформирование отчета принят. Статус: 'в процессе'"})

	//log.Println(time.Now(), " Finished Defects - Main: durance[", time.Now().Sub(mainTime), "]")
	//fmt.Println(time.Now(), " Finished Defects - Main: durance[", time.Now().Sub(mainTime), "]")
	//c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//c.File("./files/defects/res.xlsx")
}

func GetDefectsPfList(c *gin.Context) {
	orders, err := service.GetFormedDefectsList(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetDefectsLsList(c *gin.Context) {
	orders, err := service.GetFormedDefectsList(false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetDefectExcel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "некорректный параметр id"})
		return
	}

	order, err := service.GetFormedDefectByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	switch order.Status {
	case "сформирован":
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		//c.File("./files/defects/res.xlsx")
		c.File(fmt.Sprintf("./%s", order.FileName))
	case "ошибка при формировании":
		c.JSON(http.StatusBadRequest, gin.H{"reason": "отчет не возможно скачать. Возникла ошибка при формировании"})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"reason": "отчет формируется. Пожалуйста подождите..."})
	}
}

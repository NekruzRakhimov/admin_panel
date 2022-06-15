package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
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

func SendNotificationSegment(path string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "aziz.rahimov0001@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "aziz.rahimov0001@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")
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
		panic(err)
	}
	fmt.Println("successfully sent email!")
	return

}

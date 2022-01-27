package controller

//
//doc, err := document.Open("document.docx")
//if err != nil {
//log.Fatalf("error opening document: %s", err)
//}
//defer doc.Close()
//
//paragraphs := []document.Paragraph{}
//for _, p := range doc.Paragraphs() {
//paragraphs = append(paragraphs, p)
//}
//
//// This sample document uses structured document tags, which are not common
//// except for in document templates.  Normally you can just iterate over the
//// document's paragraphs.
//for _, sdt := range doc.StructuredDocumentTags() {
//for _, p := range sdt.Paragraphs() {
//paragraphs = append(paragraphs, p)
//}
//}
//
//for _, p := range paragraphs {
//for _, r := range p.Runs() {
////fmt.Println("")
////fmt.Println(">>>>>>>")
////fmt.Println(r.Text())
////fmt.Println(">>>>>>>")
////fmt.Println("")
//switch r.Text() {
//
//case "NEKRUZ": //CONTRACT_NUMBER
//r.ClearContent()
//r.AddText("11111111")
//case "BENEFICIARY":
//r.ClearContent()
//r.AddText("CUSTOM_BENEFICIARY")
//
//case "ADDRESS":
//r.ClearContent()
//r.AddText("CUSTOM_ADDRESS")
//case "BENEFICIARY_BANK_ADDRESS":
//r.ClearContent()
//r.AddText("CUSTOM_BENEFICIARY_BANK_ADDRESS")
//case "BENEFICIARY_BANK":
//r.ClearContent()
//r.AddText("CUSTOM_BENEFICIARY_BANK")
//case "SWIFT_CODE":
//r.ClearContent()
//r.AddText("CUSTOM_SWIFT_CODE")
//case "ACCOUNT":
//r.ClearContent()
//r.AddText("CUSTOM_ACCOUNT")
//case "AZIZ": // END-DATE
//r.ClearContent()
//r.AddText("CUSTOM_END_DATE")
//case "AMOUNT":
//r.ClearContent()
//r.AddText("CUSTOM_CONTRACT_AMOUNT")
//case "INTERVAL": // DeliveryTimeInterval
//r.ClearContent()
//r.AddText("CUSTOM_DELIVERY_INTERVAL")
//case "DELIVERY_DATE": // DeliveryTimeInterval
//r.ClearContent()
//r.AddText("CUSTOM_DELIVERY_DATE")
//case "RETURNTIME": // ReturnTimeDelivery
//r.ClearContent()
//r.AddText("CUSTOM_RETURN_TIME_DELIVERY")
//case "DELIVERIES": // DATE_OF_DELIVERY
//r.ClearContent()
//r.AddText("CUSTOM_DELIVERY_DATE")
////case "PREPAYMENT":
//// r.ClearContent()
//// r.AddText("CUSTOM_PREPAYMENT")
//case "TABLE_PLACE":
//// First Table
//r.ClearContent()
//
////paragraph := doc.InsertParagraphAfter(p)
////paragraph.AddRun().AddText("")
//
//table := doc.InsertTableAfter(p)
//// width of the page
//table.Properties().SetWidthPercent(100)
//// with thick borers
//borders := table.Properties().Borders()
//borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)
//
//row := table.AddRow()
//run := row.AddCell().AddParagraph().AddRun()
//run.AddText("№")
//row.AddCell().AddParagraph().AddRun().AddText("Торговое название / Trade Name ")
//row.AddCell().AddParagraph().AddRun().AddText("ТЦена, CIP Алматы, в долларах США / Price, CIP Almaty, USD ")
//run.Properties().SetHighlight(wml.ST_HighlightColorYellow)
//
//for i := 1; i <= 5; i++ {
//row = table.AddRow()
//row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%d", i))
//row.AddCell().AddParagraph().AddRun().AddText("NAME")
//row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%s", "PRICE"))
//}
//default:
//fmt.Println("not modifying", r.Text())
//}
//}
//}
//
//doc.SaveToFile("edit-document.docx")

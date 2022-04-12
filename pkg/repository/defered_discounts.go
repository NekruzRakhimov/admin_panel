package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"github.com/jinzhu/gorm"
)

func AddOrUpdateDdReport(report models.StoredReport) error {
	var checker models.StoredReport
	sqlQuery := `SELECT *
				FROM stored_reports
				WHERE bin = ?
				  AND contract_id = ?
				  AND start_date = ?
				  AND end_date = ?`
	err := db.GetDBConn().Raw(sqlQuery, report.Bin, report.ContractID, report.StartDate, report.EndDate).Scan(&checker).Error

	if err == gorm.ErrRecordNotFound {
		return AddDdReport(report)
	}

	if err != nil {
		return err
	}

	return UpdateDdReport(report)
}

func AddDdReport(report models.StoredReport) error {
	sqlQuery := `INSERT INTO stored_reports (
				bin, 
				contract_amount, 
				start_date, 
				end_date, 
				discount_amount,
				contract_id,
			    beneficiary,
				contract_number,
				content,
                contract_amount_with_discount) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if err := db.GetDBConn().Exec(sqlQuery,
		report.Bin,
		report.ContractAmount,
		report.StartDate,
		report.EndDate,
		report.DiscountAmount,
		report.ContractID,
		report.Beneficiary,
		report.ContractNumber,
		report.Content,
		report.ContractAmountWithDiscount).Error; err != nil {
		return err
	}

	return nil
}

func UpdateDdReport(report models.StoredReport) error {
	sqlQuery := `UPDATE stored_reports
				set contract_amount               = ?,
					discount_amount               = ?,
					contract_amount_with_discount = ?,
					beneficiary = ?,
					contract_number = ?,
					content = ?
				WHERE bin = ?
				  AND contract_id = ?
				  AND start_date = ?
				  AND end_date = ?`
	if err := db.GetDBConn().Exec(sqlQuery,
		report.ContractAmount,
		report.DiscountAmount,
		report.ContractAmountWithDiscount,
		report.Beneficiary,
		report.ContractNumber,
		report.Content,
		report.Bin,
		report.ContractID,
		report.StartDate,
		report.EndDate).Error; err != nil {
		return err
	}

	return nil
}

func GetAllDdStoredReports() (reports []models.StoredReport, err error) {
	sqlQuery := "SELECT * FROM stored_reports"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&reports).Error; err != nil {
		return nil, err
	}

	return reports, err
}

func GetDdStoredReportDetails(storedReportID int) (storedReport models.StoredReport, err error) {
	sqlQuery := "SELECT * FROM stored_reports WHERE id = ?"
	if err := db.GetDBConn().Raw(sqlQuery, storedReportID).Scan(&storedReport).Error; err != nil {
		return models.StoredReport{}, err
	}

	return
}

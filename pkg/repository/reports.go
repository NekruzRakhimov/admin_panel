package repository

import (
	"admin_panel/db"
	"admin_panel/models"
	"github.com/jinzhu/gorm"
)

func AddOrUpdateReport(report models.StoredReport) error {
	var checker models.StoredReport
	sqlQuery := `SELECT *
				FROM stored_reports
				WHERE bin = ?
				  AND contract_id = ?
				  AND start_date = ?
				  AND end_date = ?`
	err := db.GetDBConn().Raw(sqlQuery, report.Bin, report.ContractID, report.StartDate, report.EndDate).Scan(&checker).Error

	if err == gorm.ErrRecordNotFound {
		return AddReport(report)
	}

	if err != nil {
		return err
	}

	return UpdateReport(report)
}

func AddReport(report models.StoredReport) error {
	sqlQuery := `INSERT INTO stored_reports (
				bin, 
				contract_amount, 
				start_date, 
				end_date, 
				discount_amount,
				contract_id, 
				contract_number,
                contract_amount_with_discount) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	if err := db.GetDBConn().Exec(sqlQuery,
		report.Bin,
		report.ContractAmount,
		report.StartDate,
		report.EndDate,
		report.DiscountAmount,
		report.ContractID,
		report.ContractNumber,
		report.ContractAmountWithDiscount).Error; err != nil {
		return err
	}

	return nil
}

func UpdateReport(report models.StoredReport) error {
	sqlQuery := `UPDATE stored_reports
				set contract_amount               = ?,
					discount_amount               = ?,
					contract_amount_with_discount = ?,
					contract_number = ?,
				WHERE bin = ?
				  AND contract_id = ?
				  AND start_date = ?
				  AND end_date = ?`
	if err := db.GetDBConn().Exec(sqlQuery,
		report.ContractAmount,
		report.DiscountAmount,
		report.ContractAmountWithDiscount,
		report.ContractNumber,
		report.Bin,
		report.ContractID,
		report.StartDate,
		report.EndDate).Error; err != nil {
		return err
	}

	return nil
}

func GetAllStoredReports() (reports []models.StoredReport, err error) {
	sqlQuery := "SELECT * FROM stored_reports"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&reports).Error; err != nil {
		return nil, err
	}

	return reports, err
}

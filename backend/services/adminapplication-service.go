package services

import (
	"kfs-backend/database"
)

func GetAllRoleApplications() ([]map[string]interface{}, error) {
	var applications []map[string]interface{}
	query := "SELECT * FROM roleapplicationforms"
	err := database.DB.Raw(query).Scan(&applications).Error
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func UpdateRoleApplicationStatus(applicationId int, status string) error {
	tx := database.DB.Begin()
	updateQuery := `UPDATE roleapplicationforms SET status = $1 WHERE applicationid = $2`
	if err := tx.Exec(updateQuery, status, applicationId).Error; err != nil {
		tx.Rollback()
		return err
	}
	if status == "accepted" {
		var userId int
		var applicationType string
		query := `SELECT userId, applicationType FROM roleapplicationforms WHERE applicationid = $1`
		if err := tx.Raw(query, applicationId).Row().Scan(&userId, &applicationType); err != nil {
			tx.Rollback()
			return err
		}
		roleUpdateQuery := `UPDATE roles SET role = array_append(role, $1) WHERE userId = $2`
		if err := tx.Exec(roleUpdateQuery, applicationType, userId).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
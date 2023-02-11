package models

import (
	"time"

	"gorm.io/gorm"
)

type VisitVrach struct {
	ID               uint `gorm:"primaryKey"`
	PODRAZD_ID       int
	TEAM_ID          int
	VISIT_DATE       *time.Time
	LOCATION_DATE    *time.Time
	LAT              string
	LONG             string
	USER_ID          int
	TG_CHAT_ID       int
	REGION_ID        int
	LPU_ID           int
	VRACH_ID         int
	COMMENT          string
	SOVEMESTNO       int
	STATUS           string
	OTSENKA_VISIT_ID int
}

func (VisitVrach) TableName() string {
	return "md_visit_vrach"
}
func GetVisitsVrach(db *gorm.DB, VisitVrach *[]VisitVrach) (err error) {
	err = db.Limit(100).Find(VisitVrach).Error
	if err != nil {
		return err
	}
	return nil
}
func GetVisitsVrachAll(db *gorm.DB, VisitVrach *[]VisitVrach) (err error) {
	err = db.Find(VisitVrach).Error
	if err != nil {
		return err
	}
	return nil
}

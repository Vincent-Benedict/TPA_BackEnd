package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type Report struct {
	ID           uint `gorm:"primary_key"`
	ReportedId   int
	ReporterId   int
	ReportReason string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&Report{})
	db.AutoMigrate(&Report{})
}

func init(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&Report{
		ReportedId:   1,
		ReporterId:   2,
		ReportReason: "He is Noob",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&Report{
		ReportedId:   1,
		ReporterId:   3,
		ReportReason: "He is Noob too",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&Report{
		ReportedId:   1,
		ReporterId:   4,
		ReportReason: "I Hate this Guy ! He is Trully Unbelievable",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&Report{
		ReportedId:   1,
		ReporterId:   3,
		ReportReason: "Fuck This Guy",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	//db.Create(&Report{
	//	ReportedId:   1,
	//	ReporterId:   4,
	//	ReportReason: "Nobble this guy for the noob !",
	//	CreatedAt:    time.Time{},
	//	UpdatedAt:    time.Time{},
	//	DeletedAt:    nil,
	//})



	db.Create(&Report{
		ReportedId:   2,
		ReporterId:   1,
		ReportReason: "He is From Me",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
	db.Create(&Report{
		ReportedId:   3,
		ReporterId:   1,
		ReportReason: "He is Unknown",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
}

func DeleteReportByReportedId(reportedId int)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Where("reported_id = ?", reportedId).Delete(&Report{})

	return true, nil
}

func InsertReport(token string, reportedId int, reportReason string)(bool, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	reporterId := int(getIdFromJWTToken(token))

	db.Create(&Report{
		ReportedId:   reportedId,
		ReporterId:   reporterId,
		ReportReason: reportReason,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	return true, nil
}

func GetReportedByJWTToken(token string)([]Report, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	reporterId := getIdFromJWTToken(token)

	var report []Report
	db.Where("reported_id = ? and DATE_PART( 'day', NOW()) - DATE_PART( 'day', created_at) <= 7", reporterId).Find(&report)

	return report, nil
}

func GetReportByReportedId(id int)([]Report, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var report []Report
	db.Where("reported_id = ?", id).Find(&report)

	return report, nil
}
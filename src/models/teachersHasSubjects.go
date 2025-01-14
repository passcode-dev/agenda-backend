package models

import "time"

type TeacherSubject struct {
	TeacherID uint `gorm:"primaryKey" json:"teacher_id"`
	SubjectID uint `gorm:"primaryKey" json:"subject_id"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
}

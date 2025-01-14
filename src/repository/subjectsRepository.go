// src/repository/subjects_repository.go
package repository

import (
	"agenda-backend/src/database"
	"agenda-backend/src/models"
)

func GetAllSubjects(page int) ([]models.Subjects, error) {
	var subjects []models.Subjects
	offset := (page - 1) * 10
	if err := database.DB.Limit(10).Offset(offset).Find(&subjects).Error; err != nil {
		return nil, err
	}
	return subjects, nil
}

func CreateSubjectWithAssociation(subject *models.Subjects, teacherID *uint) error {
	tx := database.DB.Begin()

	if err := tx.Create(subject).Error; err != nil {
		tx.Rollback()
		return err
	}

	if teacherID != nil {
		association := models.TeacherSubject{
			TeacherID: *teacherID,
			SubjectID: subject.ID,
		}
		if err := tx.Create(&association).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func DeleteSubject(id uint) error {
	return database.DB.Delete(&models.Subjects{}, id).Error
}

func UpdateSubject(id uint, name string) error {
	return database.DB.Model(&models.Subjects{}).Where("id = ?", id).Update("name", name).Error
}

func AssociateTeacherToSubject(teacherID, subjectID uint) error {
	association := models.TeacherSubject{
		TeacherID: teacherID,
		SubjectID: subjectID,
	}
	return database.DB.Create(&association).Error
}

func DeleteTeacherSubject(teacherID uint, subjectID uint) error {
	return database.DB.Where("teacher_id = ? AND subject_id = ?", teacherID, subjectID).Delete(&models.TeacherSubject{}).Error
}

func DeleteMultipleTeacherSubjects(teacherID uint, subjectIDs []uint) error {
	return database.DB.Where("teacher_id = ? AND subject_id IN ?", teacherID, subjectIDs).Delete(&models.TeacherSubject{}).Error
}
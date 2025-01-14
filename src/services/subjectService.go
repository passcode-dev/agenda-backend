package services

import (
	"errors"
	"agenda-backend/src/models"
	"agenda-backend/src/repository"
)

func GetAllSubjectsService(page int) ([]models.Subjects, error) {
	return repository.GetAllSubjects(page)
}

func CreateSubjectService(subject *models.Subjects, teacherID *uint) error {
	if subject.Name == "" {
		return errors.New("O campo nome é obrigatório")
	}
	return repository.CreateSubjectWithAssociation(subject, teacherID)
}

func DeleteSubjectService(id uint) error {
	return repository.DeleteSubject(id)
}

func UpdateSubjectService(id uint, name string) error {
	if name == "" {
		return errors.New("O campo nome é obrigatório")
	}
	return repository.UpdateSubject(id, name)
}

func AssociateTeacherToSubjectService(teacherID, subjectID uint) error {
	if teacherID == 0 || subjectID == 0 {
		return errors.New("IDs de professor e matéria são obrigatórios")
	}
	return repository.AssociateTeacherToSubject(teacherID, subjectID)
}

func DeleteTeacherSubjectService(teacherID uint, subjectID uint) error {
	if teacherID == 0 || subjectID == 0 {
		return errors.New("IDs de professor e matéria são obrigatórios")
	}
	return repository.DeleteTeacherSubject(teacherID, subjectID)
}

func DeleteMultipleTeacherSubjectsService(teacherID uint, subjectIDs []uint) error {
	if teacherID == 0 || len(subjectIDs) == 0 {
		return errors.New("IDs de professor e matérias são obrigatórios")
	}
	return repository.DeleteMultipleTeacherSubjects(teacherID, subjectIDs)
}

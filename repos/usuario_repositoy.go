package repos

import (
	"errors"

	"github.com/Diegolopez-ing/api_rest/models"
	"gorm.io/gorm"
)

/*
type UsuarioRepository interface {
	FindAll() ([]Usuario, error)
	FindById(id int) (models.Usuario, error)
	FindByUsername(username string) (models.Usuario, error)
	FindByUsernameAndPassword(username string, password string) (models.Usuario, error)
	Save(usuario models.Usuario) (models.Usuario, error)
	Update(usuario models.Usuario) (models.Usuario, error)
	Delete(usuario models.Usuario) (models.Usuario, error)
} */

type UsuarioRepository struct {
	conn *gorm.DB
}

func NewUsuarioRepository(conn *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{conn}
}

func (r *UsuarioRepository) Loging(rs models.RequestSession) (*models.Usuario, error) {
	usuario := &models.Usuario{}
	if res := r.conn.Find(usuario, "username = ? and password= ?", rs.Username, rs.Password); res.Error != nil {
		return nil, res.Error
	} else {
		if res.RowsAffected == 0 {
			return nil, errors.New("No Encontrado")
		}

	}
	return usuario, nil
}

/* func (repo *UsuarioRepository) FindAll() ([]*models.Usuario, error) {
	var usuarios []models.Usuario
	err := repo.conn.Find(&usuarios).Error
	return usuarios, err
}

func (repo *UsuarioRepository) FindById(id int) (models.Usuario, error) {
	var usuario Usuario
	err := repo.conn.First(&usuario, id).Error
	return usuario, err
}

func (repo *UsuarioRepository) FindByUsername(username string) (models.Usuario, error) {
	var usuario Usuario
	err := repo.conn.Where("username = ?", username).First(&usuario).Error
	return usuario, err
}

func (repo *UsuarioRepository) FindByUsernameAndPassword(username string, password string) (models.Usuario, error) {
	var usuario Usuario
	err := repo.conn.Where("username = ? AND password = ?", username, password).First(&usuario).Error
	return usuario, err
}

func (repo *UsuarioRepository) Save(usuario Usuario) (models.Usuario, error) {
	err := repo.conn.Create(&usuario).Error
	return usuario, err
}

func (repo *UsuarioRepository) Update(usuario Usuario) (models.Usuario, error) {
	err := repo.conn.Save(&usuario).Error
	return usuario, err
}

func (repo *UsuarioRepository) Delete(usuario Usuario) (models.Usuario, error) {
	err := repo.conn.Delete(&usuario).Error
	return usuario, err
}
*/
// Language: go
// Path: repos/usuario_repositoy.go
// Compare this snippet from models/usuario.go:
// package models
//

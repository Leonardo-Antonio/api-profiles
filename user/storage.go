package user

import (
	"database/sql"
	"github.com/Leonardo-Antonio/api-profiles/helper"
	"github.com/asaskevich/govalidator"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) GetAll() (users []model, err error) {
	stmt, err := s.db.Prepare(sqlGetAll)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	profileNull := sql.NullString{}

	for rows.Next() {
		user := model{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.LastName,
			&user.Phone,
			&user.Email,
			&user.Password,
			&profileNull,
		)
		if err != nil {
			return users, err
		}
		user.Profile = profileNull.String
		users = append(users, user)
	}
	return
}

func (s *Storage) SignUp(data model) error {
	stmt, err := s.db.Prepare(sqlSignUp)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if !govalidator.IsEmail(data.Email) {
		return helper.ErrEmailInvalid
	}
	if len(data.Name) < 1 && len(data.LastName) < 1 {
		return helper.ErrPersonalInformationInvalid
	}
	if len(data.Password) < 10 {
		return helper.ErrInsecurePassword
	}
	rs, err := stmt.Exec(
		helper.NullString(data.Name),
		helper.NullString(data.LastName),
		helper.NullString(data.Phone),
		helper.NullString(data.Email),
		helper.NullString(data.Password),
		helper.NullString(data.Profile),
	)
	if err != nil {
		return err
	}

	rA, err := rs.RowsAffected()
	if rA != 1 {
		return helper.ErrRowNotAffected
	}

	return nil
}

func (s *Storage) SignIn(identification model) (data model, err error) {
	stmt, err := s.db.Prepare(sqlSignIn)
	if err != nil {
		return data, helper.ErrStmtSQL
	}
	defer stmt.Close()

	profileNull := sql.NullString{}

	err = stmt.QueryRow(identification.Email, identification.Password).Scan(
		&data.ID,
		&data.Name,
		&data.LastName,
		&data.Phone,
		&data.Email,
		&profileNull,
	)
	data.Profile = profileNull.String
	if err != nil {
		return
	}
	return
}

func (s *Storage) Delete(identification model) error {
	stmt, err := s.db.Prepare(sqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(identification.Email, identification.Password)
	if err != nil {
		return err
	}
	rA, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if rA != 1 {
		return helper.ErrUserInvalid
	}
	return nil
}

func (s *Storage) Update(data model) error {
	stmt, err := s.db.Prepare(sqlUpdate)
	if err != nil {
		return helper.ErrStmtSQL
	}
	defer stmt.Close()

	if len(data.Name) < 1 && len(data.LastName) < 1 {
		return helper.ErrPersonalInformationInvalid
	}
	rs, err := stmt.Exec(
		helper.NullString(data.Name),
		helper.NullString(data.LastName),
		helper.NullString(data.Phone),
		helper.NullString(data.Profile),
		data.ID,
	)
	if err != nil {
		return err
	}

	rA, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if rA != 1 {
		return helper.ErrRowNotAffected
	}
	return nil
}

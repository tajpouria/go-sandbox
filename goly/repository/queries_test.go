package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo Repository
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	require.NoError(s.T(), err)

	s.repo = NewRepository(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) TestGetGolies() {
	g1, g2 := Goly{
		ID:       1,
		Redirect: "a",
		Goly:     "b",
		Clicked:  false,
		Random:   false,
	}, Goly{
		ID:       2,
		Redirect: "c",
		Goly:     "d",
		Clicked:  true,
		Random:   true,
	}

	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "golies"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "redirect", "goly", "clicked", "random"}).
				AddRow(g1.ID, g1.Redirect, g1.Goly, g1.Clicked, g1.Random).
				AddRow(g2.ID, g2.Redirect, g2.Goly, g2.Clicked, g2.Random),
		)

	res, err := s.repo.GetGolies()
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal([]Goly{g1, g2}, res))
}

func (s *Suite) TestGetGoly() {
	var (
		id       = uint64(1)
		redirect = "a"
		goly     = "b"
		clicked  = false
		random   = false
	)

	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "golies" WHERE id = $1 ORDER BY "golies"."id" LIMIT 1`)).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "redirect", "goly", "clicked", "random"}).
				AddRow(id, redirect, goly, clicked, random),
		)

	res, err := s.repo.GetGoly(id)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(Goly{
		ID:       id,
		Redirect: redirect,
		Goly:     goly,
		Clicked:  clicked,
		Random:   random,
	}, res))
}

func (s *Suite) TestGetGolyByURL() {
	var (
		id       = uint64(1)
		redirect = "a"
		goly     = "b"
		clicked  = false
		random   = false
	)

	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "golies" WHERE goly = $1 ORDER BY "golies"."id" LIMIT 1`)).
		WithArgs(goly).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "redirect", "goly", "clicked", "random"}).
				AddRow(id, redirect, goly, clicked, random),
		)

	res, err := s.repo.GetGolyByURL(goly)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(
		Goly{
			ID:       id,
			Redirect: redirect,
			Goly:     goly,
			Clicked:  clicked,
			Random:   random,
		}, res))
}

func (s *Suite) TestCreateGoly() {
	var (
		id       = uint64(1)
		redirect = "a"
		goly     = "b"
		clicked  = false
		random   = false
	)

	s.mock.ExpectBegin()
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "golies" ("redirect","goly","clicked","random","id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(redirect, goly, clicked, random, id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
	s.mock.ExpectCommit()

	err := s.repo.CreateGoly(&Goly{
		ID:       id,
		Redirect: redirect,
		Goly:     goly,
		Clicked:  clicked,
		Random:   random,
	})
	require.NoError(s.T(), err)
}

func (s *Suite) TestUpdateGoly() {
	var (
		id       = uint64(1)
		redirect = "a"
		goly     = "b"
		clicked  = false
		random   = false
	)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "golies" SET "redirect"=$1,"goly"=$2,"clicked"=$3,"random"=$4 WHERE "id" = $5`)).
		WithArgs(redirect, goly, clicked, random, id)
	s.mock.ExpectRollback()

	err := s.repo.UpdateGoly(&Goly{
		ID:       id,
		Redirect: redirect,
		Goly:     goly,
		Clicked:  clicked,
		Random:   random,
	})
	require.NoError(s.T(), err)
}

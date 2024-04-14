package common

import (
	"github.com/jmoiron/sqlx"
	"github.com/walnuts1018/2024-golang-linebot/common/config"

	_ "github.com/lib/pq"
)

type Storage interface {
	AddSubject(subject Subject) error
	GetSubjects() ([]Subject, error)
}

type DBClient struct {
	db *sqlx.DB
}

func NewDBClient(config config.Config) (Storage, error) {
	db, err := sqlx.Open("postgres", config.PSQLDSN)
	if err != nil {
		return nil, err
	}

	c := &DBClient{
		db: db,
	}

	if err := c.Init(); err != nil {
		c.Close()
		return nil, err
	}

	return c, nil
}

func (c *DBClient) Close() error {
	return c.db.Close()
}

func (c *DBClient) DB() *sqlx.DB {
	return c.db
}

func (c *DBClient) Init() error {
	if err := c.CreateSubjectTable(); err != nil {
		return err
	}

	return nil
}

func (c *DBClient) CreateSubjectTable() error {
	_, err := c.db.Exec("CREATE TABLE IF NOT EXISTS subjects (id SERIAL PRIMARY KEY, name TEXT, weekday TEXT, period INTEGER, room TEXT)")
	return err
}

func (c *DBClient) AddSubject(subject Subject) error {
	_, err := c.db.Exec("INSERT INTO subjects (name, weekday, period, room) VALUES ($1, $2, $3, $4)", subject.Name, subject.Weekday, subject.Period, subject.Room)
	return err
}

func (c *DBClient) GetSubjects() ([]Subject, error) {
	rows, err := c.db.Queryx("SELECT * FROM subjects")
	if err != nil {
		return nil, err
	}

	var subjects []Subject
	for rows.Next() {
		var subject Subject
		if err := rows.StructScan(&subject); err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

type InmemoryDB struct {
	subjects []Subject
}

func NewInmemoryDB() Storage {
	return &InmemoryDB{
		subjects: make([]Subject, 0),
	}
}

func (i *InmemoryDB) AddSubject(subject Subject) error {
	i.subjects = append(i.subjects, subject)
	return nil
}

func (i *InmemoryDB) GetSubjects() ([]Subject, error) {
	return i.subjects, nil
}

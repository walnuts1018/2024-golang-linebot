package common

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/walnuts1018/2024-golang-linebot/common/config"

	_ "github.com/lib/pq"
)

type Storage interface {
	AddSubject(subject Subject, userid string) error
	GetSubjects(userid string) ([]Subject, error)
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
	_, err := c.db.Exec("CREATE TABLE IF NOT EXISTS subjects (id SERIAL PRIMARY KEY, user_id TEXT PRIMARY KEY, name TEXT, weekday TEXT, period INTEGER, room TEXT)")
	return err
}

func (c *DBClient) AddSubject(subject Subject, userid string) error {
	_, err := c.db.Exec("INSERT INTO subjects (user_id, name, weekday, period, room) VALUES ($1, $2, $3, $4, $5)", userid, subject.Name, subject.Weekday, subject.Period, subject.Room)
	return err
}

func (c *DBClient) GetSubjects(userid string) ([]Subject, error) {
	rows, err := c.db.Queryx("SELECT * FROM subjects WHERE user_id = $1", userid)
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
	subjects map[string][]Subject // key: userid
}

func NewInmemoryDB() Storage {
	return &InmemoryDB{
		subjects: make(map[string][]Subject),
	}
}

func (i *InmemoryDB) AddSubject(subject Subject, userid string) error {
	if _, ok := i.subjects[userid]; !ok {
		i.subjects[userid] = []Subject{}
	}

	i.subjects[userid] = append(i.subjects[userid], subject)

	return nil

}

func (i *InmemoryDB) GetSubjects(userid string) ([]Subject, error) {
	if _, ok := i.subjects[userid]; !ok {
		return []Subject{}, nil
	}

	return i.subjects[userid], nil
}

type FileDB struct {
	subjects map[string][]Subject // key: userid
	path     string
}

func NewFileDB(path string) (Storage, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(path)
			if err != nil {
				return nil, fmt.Errorf("failed to create file: %w", err)
			}
			defer file.Close()

			_, err = file.Write([]byte("[]"))
			if err != nil {
				return nil, fmt.Errorf("failed to write file: %w", err)
			}

			return &FileDB{
				path:     path,
				subjects: make(map[string][]Subject),
			}, nil
		} else {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
	}
	defer file.Close()

	var subjects map[string][]Subject
	dec := json.NewDecoder(file)
	if err := dec.Decode(&subjects); err != nil {
		return nil, fmt.Errorf("failed to decode subjects: %w", err)
	}

	return &FileDB{
		path:     path,
		subjects: subjects,
	}, nil
}

func (f *FileDB) AddSubject(subject Subject, userid string) error {

	file, err := os.Create(f.path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, ok := f.subjects[userid]; !ok {
		f.subjects[userid] = []Subject{}
	}

	subjects := f.subjects
	subjects[userid] = append(subjects[userid], subject)

	enc := json.NewEncoder(file)
	if err := enc.Encode(subjects); err != nil {
		return fmt.Errorf("failed to encode subjects: %w", err)
	}

	f.subjects = subjects

	return nil
}

func (f *FileDB) GetSubjects(userid string) ([]Subject, error) {
	if _, ok := f.subjects[userid]; !ok {
		return []Subject{}, nil
	}

	return f.subjects[userid], nil
}

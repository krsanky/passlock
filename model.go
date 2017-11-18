package passlock

import (
	"fmt"
	"time"

	"github.com/krsanky/go-postgres-template/db"
	"oldcode.org/webplay/account"
)

//CREATE TABLE passlock (
//    id SERIAL PRIMARY KEY,
//    account_id INT REFERENCES account ON DELETE CASCADE NOT NULL,
//    title text NOT NULL,
//    password text NOT NULL,
//    ts TIMESTAMP WITHOUT TIME ZONE NOT NULL
type Passlock struct {
	Id        int
	AccountId int
	Title     string
	Password  string
	Release   time.Time
}

func Create(account_id int, title, password string, release time.Time) *Passlock {
	pl := &Passlock{}
	pl.Title = title
	pl.Password = password
	pl.AccountId = account_id
	pl.Release = release
	return pl
}

func (p *Passlock) Save() error {
	var id int
	err := db.DB.QueryRow(`
INSERT INTO passlock 
(account_id, title, password, ts)
VALUES ($1, $2, $3, $4) RETURNING id`,
		p.AccountId, p.Title, p.Password, p.Release).Scan(&id)

	p.Id = id

	return err
}

func (p *Passlock) String() string {
	return fmt.Sprintf("<id:%d aid:%d name:%s>", p.Id, p.AccountId, p.Title)
}

func (p *Passlock) Delete() error {
	_, err := db.DB.Exec(`
DELETE FROM passlock
WHERE id = $1`,
		p.Id)
	return err
}

func Get(id int) (*Passlock, error) {
	pl := &Passlock{}
	row := db.DB.QueryRow(`
SELECT id, account_id, title, password, ts
FROM passlock
WHERE id = $1`, id)

	err := row.Scan(
		&pl.Id,
		&pl.AccountId,
		&pl.Title,
		&pl.Password,
		&pl.Release)

	return pl, err
}

func GetIds(u *account.User) ([]int, error) {
	var (
		id  int
		ids []int
	)
	rows, err := db.DB.Query(`
SELECT id from passlock
WHERE account_id = $1`, u.Id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

func GetAll(u *account.User) ([]Passlock, error) {
	ids, err := GetIds(u)
	if err != nil {
		return nil, err
	}
	pls := make([]Passlock, len(ids))
	for i, id := range ids {
		pl, err := Get(id)
		if err != nil {
			return nil, err
		}
		pls[i] = *pl
	}
	return pls, nil
}

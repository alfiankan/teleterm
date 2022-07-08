package teleterm

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

type Button struct {
	Name string
	Cmd  string
}

type Persistence struct {
	db *sql.DB
}

func (p *Persistence) AddButton(ctx context.Context, name string, cmd string) error {

	if _, err := p.GetButtonByname(ctx, name); err == nil {
		return errors.New("Button Already Exist")
	}

	stmt, err := p.db.PrepareContext(ctx, "INSERT INTO buttons (id, name, cmd) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid.NewString(), name, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (p *Persistence) GetButtonByname(ctx context.Context, buttonName string) (button Button, err error) {
	row := p.db.QueryRowContext(ctx, "SELECT name, cmd FROM buttons WHERE name = ?", buttonName)
	var name string
	var cmd string
	err = row.Scan(&name, &cmd)
	if err != nil {
		return
	}
	button = Button{
		Name: name,
		Cmd:  cmd,
	}
	return
}

func (p *Persistence) GetAllButtons(ctx context.Context) (buttons []Button, err error) {
	rows, err := p.db.QueryContext(ctx, "SELECT name, cmd FROM buttons")
	if err != nil {
		return
	}

	var name string
	var cmd string

	for rows.Next() {
		err = rows.Scan(&name, &cmd)
		if err != nil {
			return
		}
		buttons = append(buttons, Button{
			Name: name,
			Cmd:  cmd,
		})
	}
	return
}

func (p *Persistence) DeleteButtonsByName(ctx context.Context, name string) error {
	stmt, err := p.db.PrepareContext(ctx, "DELETE FROM buttons WHERE name = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}
	return nil

}

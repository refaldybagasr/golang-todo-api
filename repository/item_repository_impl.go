package repository

import (
	"context"
	"database/sql"
	"refaldybagasr/golang-todo-api/helper"
	"refaldybagasr/golang-todo-api/model/domain"
)

type ItemRepositoryImpl struct {
}

func (repository *ItemRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, item domain.Item) domain.Item {
	//TODO implement me
	SQL := "insert into items(name, status) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, item.Name, item.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	item.Id = int(id)
	return item
}

func (repository *ItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, item domain.Item) domain.Item {
	//TODO implement me
	SQL := "update items set name = ?, status = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, item.Name, item.Status, item.Id)
	helper.PanicIfError(err)

	return item
}

func (repository *ItemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, item domain.Item) {
	//TODO implement me
	SQL := "delete from items where id = ?"
	_, err := tx.ExecContext(ctx, SQL, item.Id)

	helper.PanicIfError(err)
}

func (repository *ItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Item {
	//TODO implement me
	SQL := "select * from items"
	var items []domain.Item

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var item domain.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Status)
		helper.PanicIfError(err)
		items = append(items, item)
	}

	return items
}

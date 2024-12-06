package main

import "context"

type Item struct {
	ID        int
	Title     string
	Completed bool
}

type Tasks struct {
	Items          []*Item
	Count          int
	CompletedCount int
}

func FetchTasks() ([]*Item, error) {
	var items []*Item
	rows, err := DB.Query("select id, title, completed from tasks order by position")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		i := Item{}
		if err := rows.Scan(&i.ID, &i.Title, &i.Completed); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	return items, nil
}

func FetchTask(id int) (*Item, error) {
	var i Item
	if err := DB.QueryRow("select id, title, completed from tasks where id = (?)", id).Scan(&i.ID, &i.Title, &i.Completed); err != nil {
		return nil, err
	}
	return &i, nil
}

func UpdateTask(id int, title string) (*Item, error) {
	var i Item
	if err := DB.QueryRow("update tasks set title = (?) where id = (?) returning id, title, completed", title, id).Scan(&i.ID, &i.Title, &i.Completed); err != nil {
		return nil, err
	}
	return &i, nil
}

func FetchCount() (int, error) {
	var count int
	if err := DB.QueryRow("select count(*) from tasks").Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func FetchCountCompleted() (int, error) {
	var count int
	if err := DB.QueryRow("select count(*) from tasks where completed = true").Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func InsertTask(title string) (*Item, error) {
	count, err := FetchCount()
	if err != nil {
		return nil, err
	}
	var id int
	if err := DB.QueryRow("insert into tasks (title, position) values (?, ?) returning id", title, count).Scan(&id); err != nil {
		return nil, err
	}
	item := Item{
		ID:        id,
		Title:     title,
		Completed: false,
	}
	return &item, nil
}

func DeleteTask(ctx context.Context, id int) error {
	_, err := DB.Exec("delete from tasks where id = (?)", id)
	if err != nil {
		return err
	}
	rows, err := DB.Query("select id from tasks order by position")
	if err != nil {
		return err
	}
	var ids []int
	for rows.Next() {
		var i int
		if err := rows.Scan(&id); err != nil {
			return err
		}
		ids = append(ids, i)
	}
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for idx, id := range ids {
		_, err := DB.Exec("update tasks set position = (?) where id = (?)", idx, id)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func OrderTasks(ctx context.Context, values []int) error {
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for i, v := range values {
		_, err := tx.Exec("update tasks set position = (?) where id = (?)", i, v)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func ToggleTask(id int) (*Item, error) {
	var i Item
	if err := DB.QueryRow("update tasks set completed = case when completed = 1 then 0 else 1 end where id = (?) returning id, title, completed", id).Scan(&i.ID, &i.Title, &i.Completed); err != nil {
		return nil, err
	}
	return &i, nil
}

-- +goose Up
-- +goose StatementBegin
INSERT INTO coach (id, name)
VALUES (1, 'Pep Guardiola');

INSERT INTO team (name, city, coach_id)
VALUES ('Manchester City', 'Manchester', 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM coach;
DELETE FROM team;
-- +goose StatementEnd

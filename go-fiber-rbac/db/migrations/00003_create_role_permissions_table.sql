-- +goose Up
-- +goose StatementBegin
CREATE TABLE role_permissions (
  role_id INT REFERENCES roles(id) NOT NULL,
  permission_id INT REFERENCES permissions(id) NOT NULL,
  PRIMARY KEY (role_id, permission_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE role_permissions;
-- +goose StatementEnd

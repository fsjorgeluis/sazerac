package repositories

import (
    "database/sql"
    "{{ .Module }}/internal/domain/entities"
    "{{ .Module }}/internal/domain/interfaces"
)

type {{ .Entity }}MySQLRepo struct {
    DB *sql.DB
}

func New{{ .Entity }}MySQLRepo(db *sql.DB) interfaces.{{ .Entity }}Repository {
    return &{{ .Entity }}MySQLRepo{DB: db}
}

func (r *{{ .Entity }}MySQLRepo) Save(e *entities.{{ .Entity }}) error {
    // TODO: implement
    return nil
}

func (r *{{ .Entity }}MySQLRepo) FindByID(id string) (*entities.{{ .Entity }}, error) {
    // TODO: implement
    return nil, nil
}
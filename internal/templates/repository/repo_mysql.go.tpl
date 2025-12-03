package mysql

import (
    "database/sql"
    "{{ .Module }}/internal/domain/entities"
    "{{ .Module }}/internal/repository"
)

type {{ .Entity }}MySQLRepo struct {
    DB *sql.DB
}

func New{{ .Entity }}MySQLRepo(db *sql.DB) repository.{{ .Entity }}Repository {
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
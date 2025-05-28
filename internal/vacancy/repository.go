package vacancy

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"time"
)

type VacancyRepository struct {
	DbPool *pgxpool.Pool
	Logger *zerolog.Logger
}

func NewVacancyRepository(dbPool *pgxpool.Pool, logger *zerolog.Logger) *VacancyRepository {
	return &VacancyRepository{
		DbPool: dbPool,
		Logger: logger,
	}
}

func (vr *VacancyRepository) AddVacancy(form VacancyCreateForm) error {
	query := `insert into vacancies (Email, Role, Company, Salary, Type, Location, Created_at ) values (@Email, @Role, @Company, @Salary, @Type, @Location, @Created_at )`
	args := pgx.NamedArgs{
		"Email":      form.Email,
		"Role":       form.Role,
		"Company":    form.Company,
		"Salary":     form.Salary,
		"Type":       form.Type,
		"Location":   form.Location,
		"Created_at": time.Now(),
	}
	_, err := vr.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("add vacancy failed: %w", err)
	}
	return nil
}

func (vr *VacancyRepository) GetCount() int {
	query := `select count(*) from vacancies`
	var count int
	vr.DbPool.QueryRow(context.Background(), query).Scan(&count)
	return count

}

func (vr *VacancyRepository) GetAll(limit, offset int) ([]Vacancy, error) {
	query := `select * from vacancies order by created_at limit @limit offset @offset`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := vr.DbPool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}
	vacancies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vacancy])
	if err != nil {
		return nil, err
	}
	return vacancies, nil

}

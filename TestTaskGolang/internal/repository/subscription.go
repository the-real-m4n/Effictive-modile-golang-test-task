package repository

import (
	"context"
	"fmt"
	models "subscriptions-service/internal/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool" //движок для постгрес
)

type SubscriptionRepo struct {
	db *pgxpool.Pool
}

func NewSubscriprionRepo(db *pgxpool.Pool) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}

}

func (r *SubscriptionRepo) Create(ctx context.Context, s models.Subscription) error {
	query := `
		INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx, query, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate)
	return err

}

func (r *SubscriptionRepo) GetAll(ctx context.Context) ([]models.Subscription, error) {
	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []models.Subscription
	for rows.Next() {
		var sub models.Subscription
		if err := rows.Scan(
			&sub.ID,
			&sub.ServiceName,
			&sub.Price,
			&sub.UserID,
			&sub.StartDate,
			&sub.EndDate,
		); err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}

	return subs, nil
}

func (r *SubscriptionRepo) GetByID(ctx context.Context, s string) ([]models.Subscription, error) {
	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE user_id = $1`

	rows, err := r.db.Query(ctx, query, s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []models.Subscription
	for rows.Next() {
		var sub models.Subscription
		if err := rows.Scan(
			&sub.ID,
			&sub.ServiceName,
			&sub.Price,
			&sub.UserID,
			&sub.StartDate,
			&sub.EndDate,
		); err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}

	return subs, nil
}

func (r *SubscriptionRepo) Update(ctx context.Context, sub models.Subscription) error {
	query := `
        UPDATE subscriptions
        SET service_name = $1, price = $2, user_id = $3, start_date = $4, end_date = $5
        WHERE id = $6
    `
	cmdTag, err := r.db.Exec(ctx, query,
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate,
		sub.EndDate,
		sub.ID,
	)

	if cmdTag.RowsAffected() == 0 {

		return fmt.Errorf("no rows updated")
	}

	return err

}

func (r *SubscriptionRepo) Delete(ctx context.Context, id int) error {
	query := `
        DELETE FROM subscriptions WHERE id = $1
    `
	_, err := r.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	return err

}

func (r *SubscriptionRepo) GetTotalPrice(ctx context.Context, userID string, serviceName string, from, to time.Time) (int, error) {
	query := `
        SELECT COALESCE(SUM(price * (DATE_PART('year', age(end_date, start_date)) * 12 +
                                     DATE_PART('month', age(end_date, start_date)) + 1)), 0)
        FROM subscriptions
        WHERE user_id = $1
          AND service_name = $2
          AND start_date >= $3
          AND end_date <= $4
    `

	var total int
	err := r.db.QueryRow(ctx, query, userID, serviceName, from, to).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

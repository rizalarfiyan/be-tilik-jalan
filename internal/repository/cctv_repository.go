package repository

import (
	"context"
	"database/sql"

	"github.com/rizalarfiyan/be-tilik-jalan/internal/model"
	"github.com/rizalarfiyan/be-tilik-jalan/logger"
	"github.com/rs/zerolog"
)

type CCTVRepository interface {
	GetAll(ctx context.Context) (model.CCTVs, error)
}

type cctvRepository struct {
	db  *sql.DB
	log *zerolog.Logger
}

func NewCCTVRepository(db *sql.DB) CCTVRepository {
	return &cctvRepository{
		db:  db,
		log: logger.Get("cctv_repository"),
	}
}

func (r *cctvRepository) GetAll(ctx context.Context) (model.CCTVs, error) {
	var res model.CCTVs
	query := `SELECT id, title, link, latitude, longitude, width, height FROM cctvs ORDER BY title`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.CCTVItem
		err := rows.Scan(&item.Id, &item.Title, &item.Link, &item.Latitude, &item.Longitude, &item.Width, &item.Height)
		if err != nil {
			return nil, err
		}
		item.FillImage()
		res = append(res, item)
	}

	return res, nil
}

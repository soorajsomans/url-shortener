package repository

import (
	"context"

	"github.com/soorajsomans/url-shortener/internal/domain"
)

type URLRepository interface {
	Save(ctx context.Context, url *domain.URL) error

	FindByShortCode(
		ctx context.Context,
		shortCode string,
	) (*domain.URL, error)

	FindbyLongURL(
		ctx context.Context,
		longURL string,
	) (*domain.URL, error)
}

package repository

import (
	"context"
	"sync"

	"github.com/soorajsomans/url-shortener/internal/domain"
	"github.com/soorajsomans/url-shortener/internal/errors"
)

type InMemoryURLRepository struct {
	mu sync.RWMutex

	byShortCode map[string]*domain.URL
	byLongURL   map[string]*domain.URL
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		byShortCode: make(map[string]*domain.URL),
		byLongURL:   make(map[string]*domain.URL),
	}
}

func (r *InMemoryURLRepository) Save(
	ctc context.Context,
	url *domain.URL,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.byShortCode[url.ShortCode] = url
	r.byLongURL[url.LongURL] = url

	return nil
}

func (r *InMemoryURLRepository) FindByShortCode(
	ctx context.Context,
	shortCode string,
) (*domain.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, exists := r.byShortCode[shortCode]

	if !exists {
		return nil, errors.ErrURLNotFound
	}
	return url, nil
}

func (r *InMemoryURLRepository) FindbyLongURL(
	ctx context.Context,
	longURL string,
) (*domain.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, exists := r.byLongURL[longURL]

	if !exists {
		return nil, errors.ErrURLNotFound
	}
	return url, nil
}

// compiletime interface validation
var _ URLRepository = (*InMemoryURLRepository)(nil)

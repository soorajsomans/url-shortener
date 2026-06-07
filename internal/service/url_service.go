package service

import (
	"context"
	"net/url"
	"time"

	"github.com/soorajsomans/url-shortener/internal/domain"
	"github.com/soorajsomans/url-shortener/internal/errors"
	"github.com/soorajsomans/url-shortener/internal/generator"
	"github.com/soorajsomans/url-shortener/internal/repository"
)

type URLService interface {
	Shorten(
		ctx context.Context,
		longURL string,

	) (*domain.URL, error)

	Resolve(
		ctx context.Context,
		shortCode string,
	) (*domain.URL, error)
}

type urlService struct {
	repo    repository.URLRepository
	idGen   generator.IDGenerator
	codeGen generator.CodeGenerator
}

func NewURLService(
	repo repository.URLRepository,
	idGen generator.IDGenerator,
	codeGen generator.CodeGenerator,
) URLService {
	return &urlService{
		repo:    repo,
		idGen:   idGen,
		codeGen: codeGen,
	}
}

func validateURL(rawURL string) error {
	if rawURL == "" {
		return errors.ErrInvalidURL
	}

	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return errors.ErrInvalidURL
	}

	if parsed.Scheme == "" {
		return errors.ErrInvalidURL
	}

	if parsed.Host == "" {
		return errors.ErrInvalidURL
	}
	return nil
}

func (s *urlService) Shorten(
	ctx context.Context,
	longURL string,
) (*domain.URL, error) {

	if err := validateURL(longURL); err != nil {
		return nil, err
	}

	existing, err := s.repo.FindbyLongURL(ctx, longURL)
	if err == nil {
		return existing, nil
	}

	id := s.idGen.NextID()

	shortCode := s.codeGen.Generate(id)

	shortURL := &domain.URL{
		ID:        id,
		LongURL:   longURL,
		ShortCode: shortCode,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.Save(
		ctx,
		shortURL,
	); err != nil {
		return nil, err
	}
	return shortURL, nil
}

func (s *urlService) Resolve(
	ctx context.Context,
	shortCode string,
) (*domain.URL, error) {
	urlEntity, err := s.repo.FindByShortCode(
		ctx,
		shortCode,
	)

	if err != nil {
		return nil, err
	}
	if urlEntity.IsExpired() {
		return nil, errors.ErrURLExpired
	}
	return urlEntity, nil
}

var _ URLService = (*urlService)(nil)

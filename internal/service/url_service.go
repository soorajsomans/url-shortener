package service

import (
	"context"
	"log"
	"net/url"
	"time"

	"github.com/soorajsomans/url-shortener/internal/domain"
	"github.com/soorajsomans/url-shortener/internal/errors"
	"github.com/soorajsomans/url-shortener/internal/event"
	"github.com/soorajsomans/url-shortener/internal/generator"
	"github.com/soorajsomans/url-shortener/internal/messaging"
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
	repo     repository.URLRepository
	idGen    generator.IDGenerator
	codeGen  generator.CodeGenerator
	producer messaging.Producer
}

func NewURLService(
	repo repository.URLRepository,
	idGen generator.IDGenerator,
	codeGen generator.CodeGenerator,
	producer messaging.Producer,
) URLService {
	return &urlService{
		repo:     repo,
		idGen:    idGen,
		codeGen:  codeGen,
		producer: producer,
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

	// produce created event
	createdEvent := event.URLCreatedEvent{
		EventType: event.URLCreated,
		URLID:     shortURL.ID,
		ShortCode: shortURL.ShortCode,
		LongURL:   shortURL.LongURL,
		CreatedAt: shortURL.CreatedAt,
	}
	if err := s.producer.Publish(
		ctx,
		"url-events",
		createdEvent,
	); err != nil {
		log.Printf(
			"failed publishing URL_CREATED : %v", err,
		)
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

	visitedEvent := event.URLVisitedEvent{
		EventType: event.URLVisited,
		URLID:     urlEntity.ID,
		ShortCode: urlEntity.ShortCode,
		VisitedAt: time.Now().UTC(),
	}

	go func() {
		_ = s.producer.Publish(
			context.Background(),
			"url-events",
			visitedEvent,
		)
	}()
	if urlEntity.IsExpired() {
		return nil, errors.ErrURLExpired
	}
	return urlEntity, nil
}

var _ URLService = (*urlService)(nil)

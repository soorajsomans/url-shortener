package messaging

import "context"

type Producer interface {
	Publish(
		ctx context.Context,
		topic string,
		payload any,
	) error
}

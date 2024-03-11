package infra

import (
	"app/infra/messages"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewCommonDB,
	wire.NewSet(
		messages.NewSmtpEmailNoreply,
		wire.Bind(new(messages.EmailProvider), new(*messages.SmtpEmail)),
	),
	NewCache,
	NewIdGenerator,
)

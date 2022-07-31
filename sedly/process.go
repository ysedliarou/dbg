package sedly

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

const (
	TagsKey        = "sentry.tags"
	ExtrasKey      = "sentry.extras"
	ContextsKey    = "sentry.contexts"
	UserKey        = "sentry.user"
	TransactionKey = "sentry.transaction"
	LevelKey       = "sentry.level"
)

func processTags(scope *sentry.Scope, entry *logrus.Entry) {
	if tags, ok := entry.Data[TagsKey].(map[string]string); ok {
		scope.SetTags(tags)
	}
}

func processExtras(scope *sentry.Scope, entry *logrus.Entry) {
	if extras, ok := entry.Data[ExtrasKey].(map[string]interface{}); ok {
		scope.SetExtras(extras)
	}
}

func processContexts(scope *sentry.Scope, entry *logrus.Entry) {
	if contexts, ok := entry.Data[ContextsKey].(map[string]interface{}); ok {
		scope.SetContexts(contexts)
	}
}

func processUser(scope *sentry.Scope, entry *logrus.Entry) {
	if user, ok := entry.Data[UserKey].(sentry.User); ok {
		scope.SetUser(user)
	}
}

func processTransaction(scope *sentry.Scope, entry *logrus.Entry) {
	if tx, ok := entry.Data[TransactionKey].(string); ok {
		scope.SetTransaction(tx)
	}
}

func processLevel(scope *sentry.Scope, entry *logrus.Entry) {
	if text, err := entry.Level.MarshalText(); err == nil {
		scope.SetLevel(sentry.Level(text))
	}
}

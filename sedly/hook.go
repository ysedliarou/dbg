package sedly

import (
	"context"
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

const (
	hubContextKey = "sentry"
)

type SentryHook struct {
	levels          []logrus.Level
	hubContextKey   interface{}
	scopeProcessors map[string]func(*sentry.Scope, *logrus.Entry)
}

func NewSentryHook() *SentryHook {
	return &SentryHook{
		levels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		},
		hubContextKey: hubContextKey,
		scopeProcessors: map[string]func(*sentry.Scope, *logrus.Entry){
			TagsKey:        processTags,
			ExtrasKey:      processExtras,
			ContextsKey:    processContexts,
			UserKey:        processUser,
			TransactionKey: processTransaction,
			LevelKey:       processLevel,
		},
	}
}

func (sh *SentryHook) Levels() []logrus.Level {
	return sh.levels
}

func (sh *SentryHook) Fire(entry *logrus.Entry) error {
	if entry.Context != nil {
		return sh.fireWithContext(entry)
	}
	return sh.fireWithoutContext(entry)
}

func getHubFromContext(ctx context.Context, key interface{}) *sentry.Hub {
	if hub, ok := ctx.Value(key).(*sentry.Hub); ok {
		return hub
	}
	return nil
}

func (sh *SentryHook) fireWithContext(entry *logrus.Entry) error {
	_ = getHubFromContext(entry.Context, sh.hubContextKey)
	if hub := getHubFromContext(entry.Context, sh.hubContextKey); hub != nil {
		return sh.fireWithHub(entry, hub)
	}
	return sh.fireWithoutContext(entry)
}

func (sh *SentryHook) fireWithoutContext(entry *logrus.Entry) error {
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		sh.configureScope(scope, entry)
	})
	return fire(entry.Message, sentry.CaptureMessage)
}

func (sh *SentryHook) fireWithHub(entry *logrus.Entry, hub *sentry.Hub) error {
	hub.ConfigureScope(func(scope *sentry.Scope) {
		sh.configureScope(scope, entry)
	})
	return fire(entry.Message, hub.CaptureMessage)
}

func fire(message string, f func(string) *sentry.EventID) error {
	eId := f(message)
	if eId == nil {
		return errors.New("message was not captured")
	}
	return nil
}

func (sh *SentryHook) configureScope(scope *sentry.Scope, entry *logrus.Entry) {
	for _, v := range sh.scopeProcessors {
		v(scope, entry)
	}
}

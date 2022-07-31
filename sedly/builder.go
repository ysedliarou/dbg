package sedly

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

func (sh *SentryHook) WithHubContextKey(hubContextKey interface{}) *SentryHook {
	sh.hubContextKey = hubContextKey
	return sh
}

func (sh *SentryHook) WithLevels(levels []logrus.Level) *SentryHook {
	sh.levels = levels
	return sh
}

func (sh *SentryHook) WithProcessors(m map[string]func(*sentry.Scope, *logrus.Entry)) *SentryHook {
	sh.scopeProcessors = m
	return sh
}

func (sh *SentryHook) ProcessWith(key string, f func(*sentry.Scope, *logrus.Entry)) *SentryHook {
	sh.scopeProcessors[key] = f
	return sh
}

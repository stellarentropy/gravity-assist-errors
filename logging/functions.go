package logging

import (
	"github.com/go-chi/chi/v5"
)

func (l *Logger) LogRoutes(router *chi.Mux) {
	for _, v := range router.Routes() {
		l.Logger.Info().Str("pattern", v.Pattern).Msg("registered route")
	}
}

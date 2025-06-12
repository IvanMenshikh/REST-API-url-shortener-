package save

import (
	"log/slog"
	"net/http"
	mdware "url-shortener/internal/http-server/middleware/logger"
	resp "url-shortener/internal/lib/api/response"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// Важно:
//  - omitempty: в полях структуры позволяет не возвращать пустые значения в JSON ответе.
//  - validate:"required,url": используется для валидации входящих данных (например, с помощью go-playground/validator).

// Получен запрос от клиента на сохранение URL.
type Request struct {
	URL string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

// Ответ на запрос сохранения URL.
type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

// URLSaver - интерфейс для сохранения URL в хранилище.
type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

// New - обработчик HTTP запроса на сохранение URL.
func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"
		log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
			//slog.String("real_ip", mdware.getRealIP(r)),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", slog.Any("error", err))
			render.JSON(w,r, resp.Error("failed to decode request"))
			return
		}
		
		log.Info("request body decoded", slog.Any("request", req))


	}
}
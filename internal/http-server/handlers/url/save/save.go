package save

import (
	"errors"
	"log/slog"
	"net/http"

	//mdware "url-shortener/internal/http-server/middleware/logger"
	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/random"
	"url-shortener/internal/storage"

	"url-shortener/internal/config"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

// Важно:
//  - omitempty: в полях структуры позволяет не возвращать пустые значения в JSON ответе.
//  - validate:"required,url": используется для валидации входящих данных (например, с помощью go-playground/validator).

// Получен запрос от клиента на сохранение URL.
type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

// Ответ на запрос сохранения URL.
type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

// Генерируем моки для тестирования
//
//go:generate mockery --name=URLSaver
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
			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", slog.Any("error", err))
			validatorErr := err.(validator.ValidationErrors)
			// Возвращаем ошибку валидации, но ошибка не человеко-читаемая. Для теста.
			//render.JSON(w, r, resp.Error("invalid request"))
			render.JSON(w, r, resp.ValidationError(validatorErr))
			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(config.AliasMaxLength)
		}

		// TODO: Возможно в будущем алиас рандомный совпадет с существующим, надо будет придумать проверку.

		id, err := urlSaver.SaveURL(req.URL, alias)
		if errors.Is(err, storage.ErrURLExists) {
			log.Info("URL already exists", slog.String("url", req.URL))
			render.JSON(w, r, resp.Error("url already exists"))
			return
		}
		if err != nil {
			log.Error("failed to add URL", slog.Any("error", err))
			render.JSON(w, r, resp.Error("failed to add URL"))
			return
		}

		log.Info("URL added", slog.Int64("id", id))
		responseOK(w, r, alias)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Alias:    alias,
	})
}

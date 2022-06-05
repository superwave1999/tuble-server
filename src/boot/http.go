package boot

import (
	"github.com/atreugo/cors"
	"github.com/savsgio/atreugo/v11"
	"os"
	mapvalidator "tuble/src/classes/map-validator"
	"tuble/src/config"
	"tuble/src/storage/cache"
	"tuble/src/utils"
)

func StartHTTP() {
	httpAddr := os.Getenv("HTTP_LISTEN_SOCKET")
	if httpAddr == "" {
		httpAddr = "0.0.0.0:8080"
	}
	httpCorsOrigins := os.Getenv("HTTP_ALLOWED_FRONTEND_HOST")
	if httpCorsOrigins == "" {
		httpCorsOrigins = "*"
	}

	server := atreugo.New(atreugo.Config{
		Addr: httpAddr,
	})
	corsMiddleware := cors.New(cors.Config{
		AllowedOrigins:   []string{httpCorsOrigins},
		AllowedHeaders:   []string{"Content-Type"},
		AllowedMethods:   []string{"GET", "POST"},
		ExposedHeaders:   []string{},
		AllowCredentials: false, //Not needed currently.
		AllowMaxAge:      86400, //Use max available, unlikely to change.
	})
	server.UseAfter(corsMiddleware)

	//Routes.
	server.GET("/start", func(c *atreugo.RequestCtx) error {
		c.Response.Header.SetContentType("application/json; charset=utf-8")
		c.Response.SetBodyString(cache.GetCurrentMap())
		c.Response.SetStatusCode(200)
		return nil
	})
	server.GET("/config", func(c *atreugo.RequestCtx) error {
		return c.JSONResponse(config.Map, 200)
	})
	server.POST("/validate", func(c *atreugo.RequestCtx) error {
		bodyBytes := c.Request.Body()
		if len(bodyBytes) == 0 {
			return c.RawResponse("", 400)
		}
		unpackedMap, err := utils.JsonBytesToMap(bodyBytes)
		if err != nil {
			return c.RawResponse("", 400)
		}
		validation, err := mapvalidator.Validate(unpackedMap)
		if err != nil {
			return c.RawResponse("", 400)
		}
		return c.JSONResponse(validation)
	})
	server.ListenAndServe()
}

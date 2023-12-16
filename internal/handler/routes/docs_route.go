package routes

import (
  "github.com/go-chi/chi"
  httpSwagger "github.com/swaggo/http-swagger"
  "github.com/wiliamvj/api-users-golang/docs/custom"
)

var (
  docsURL = "http://localhost:8080/docs/doc.json"
)

//	@title		API users
//	@version	1.0
//	@in			header
//	@name		Authorization
func InitDocsRoutes(r chi.Router) {
  r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL),
    httpSwagger.AfterScript(custom.CustomJS),
    httpSwagger.DocExpansion("none"),
    httpSwagger.UIConfig(map[string]string{
      "defaultModelsExpandDepth": `"-1"`,
    }),
  ))
}

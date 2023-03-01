package utils

import (
	"booking/backend/errors"
	"html/template"
	"net/http"
)

func ExecuteTemplates(writer http.ResponseWriter, data interface{}, templates ...string) error {
	t, err := template.ParseFiles(templates...)
	if err != nil {
		return errors.TemplateExecuteError
	}
	err = t.ExecuteTemplate(writer, "base", data)
	if err != nil {
		return errors.TemplateExecuteError
	}
	return nil
}

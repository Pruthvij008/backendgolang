package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func Faq(tpl Template) http.HandlerFunc {
	// FAQ questions and answers
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What is Cropify?",
			Answer:   "Cropify is an AI-powered platform for enhancing farming productivity and sustainability.",
		},
		{
			Question: "How does Cropify recommend fertilizers?",
			Answer:   "Using AI and soil analysis, Cropify recommends the most suitable fertilizers for your crops.",
		},
		{
			Question: "Can I purchase fertilizers directly from Cropify?",
			Answer:   "Yes, you can purchase recommended fertilizers directly from the platform or visit our offline stores.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions) // Pass questions to the template
	}
}

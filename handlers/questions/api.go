package questions

import (
	"net/http"
)

type QuestionRoutes interface {
	GetOrgQuestions(w http.ResponseWriter, r *http.Request)
	GetQuestion(w http.ResponseWriter, r *http.Request)
	GetQuestions(w http.ResponseWriter, r *http.Request)
	GetTeamQuestions(w http.ResponseWriter, r *http.Request)
	SubmitTeamQuestion(w http.ResponseWriter, r *http.Request)
	SubmitQuestion(w http.ResponseWriter, r *http.Request)
	ViewQuestion(w http.ResponseWriter, r *http.Request)
}
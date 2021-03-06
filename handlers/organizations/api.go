package organizations

import (
	"net/http"
)

type OrganizationRoutes interface {
	DeleteOrganization(w http.ResponseWriter, r *http.Request)
	CreateOrganization(w http.ResponseWriter, r *http.Request)
	GetOrganizations(w http.ResponseWriter, r *http.Request)
	GetOrganization(w http.ResponseWriter, r *http.Request)
	GetOrganizationMembers(w http.ResponseWriter, r *http.Request)
	InsertOrganizationMember(w http.ResponseWriter, r *http.Request)
}

package api

const (
	jsonContextType = "application/json"
)

func Register(r *Router, h *Handler) {
	r.GET("/api/tutors/:languageSlug", h.GetTutors)
	r.GET("/api/tutor/:tutorSlug", h.GetTutor)
}

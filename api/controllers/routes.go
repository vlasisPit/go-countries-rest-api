package controllers

func (s *Server) initializeRoutes() {
	s.mux.HandleFunc("/countries", s.countries)
	s.mux.HandleFunc("/countries/", s.countryById)
}

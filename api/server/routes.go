package server

/**
https://dev.to/bmf_san/introduction-to-url-router-from-scratch-with-golang-3p8j
https://github.com/gsingharoy/httprouter-tutorial/tree/master/part4
Check this about routing
*/
func (s *Server) initializeRoutes() {
	s.Mux.HandleFunc("/countries", s.countries)
	s.Mux.HandleFunc("/countries/", s.countryById)
}

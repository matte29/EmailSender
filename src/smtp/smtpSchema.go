package SMTP

type SMTP struct {
	host     string
	port     string
	from     string
	password string
}

func (s *SMTP) Init(host string, port string, from string, password string) {
	s.host = host
	s.port = port
	s.from = from
	s.password = password
}

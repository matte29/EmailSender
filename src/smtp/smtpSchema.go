package sending

type smtpInfo struct {
	host     string
	port     string
	from     string
	password string
}

func (s *smtpInfo) Init(host string, port string, from string, password string) {
	s.host = host
	s.port = port
	s.from = from
	s.password = password
}

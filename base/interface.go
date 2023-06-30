package base

type WhoisService interface {
	IsAvailable(domain string) bool
	WhoisFullInformation(domain string) (map[string]interface{}, error)
}

type AuthService interface {
	Signup(email, password string) error
	Login(email, password string) error
	Logout() error
	GetToken(refreshToken ...string) error
}

type MailService interface {
	AuthService
	GetEmails(from ...string) (emails []string, err error)
}

type DomainService interface {
	AuthService
	GetDomains(opts ...string) (domains []string, err error)
}

type OTPService interface {
	Order(serviceID int) (map[string]interface{}, error)
	GetOTP(orderID int) (map[string]interface{}, error)
	CancelOrder(orderID int) error
	FinishOrder(orderID int) error
}

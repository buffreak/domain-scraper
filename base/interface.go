package base

type HelperService interface {
	StructToMap(s interface{}) (map[string]interface{}, error)
}

type WhoisService interface {
	IsAvailable(domain string) bool
	WhoisFullInformation(domain string) (map[string]interface{}, error)
}

type AuthService interface {
	HelperService
	Signup(email, password string) error
	Login(email, password string) error
	Logout() error
	GetToken(refreshToken string) error
}

type MailService interface {
	HelperService
	AuthService
	GetEmails(from ...string) (emails []string, err error)
}

type DomainService interface {
	HelperService
	AuthService
	GetDomains(opts ...string) (domains []string, err error)
}

type OTPService interface {
	HelperService
	Order(serviceID int) (map[string]interface{}, error)
	GetOTP(orderID int) (map[string]interface{}, error)
	CancelOrder(orderID int) error
	FinishOrder(orderID int) error
}

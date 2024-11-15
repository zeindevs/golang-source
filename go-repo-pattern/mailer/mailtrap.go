package mailer

type MailTrapClient struct {
}

func NewMailTrapClient(apikey, from string) (*MailTrapClient, error) {
	return &MailTrapClient{}, nil
}

func (m *MailTrapClient) Send(templateFile, username, email string, data any, isSandBox bool) (int, error) {
	return -1, nil
}

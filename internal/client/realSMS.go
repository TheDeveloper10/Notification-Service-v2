package client

import (
	"errors"
	"fmt"
	"net/http"
	"notification-service/internal/config"
	"strings"
)

func newRealSMSClientFromConfig(conf *config.SMSConfig) *realSMS {
	return &realSMS{
		httpClient:        &http.Client{},
		endpoint:          fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", conf.AccountSID),
		queryParamsFormat: fmt.Sprintf("MessagingServiceSid=%s&To=%%s&Body=%%s", conf.MessagingServiceSID),
		accountSID:        conf.AccountSID,
		authToken:         conf.AuthToken,
	}
}

type realSMS struct {
	httpClient        *http.Client
	endpoint          string
	queryParamsFormat string

	accountSID string
	authToken  string
}

func (rs *realSMS) SendSMS(title string, body string, to string) error {
	message := fmt.Sprintf("%s\r\n%s", title, body)
	relativeParameters := fmt.Sprintf(rs.queryParamsFormat, to, message)
	data := strings.NewReader(relativeParameters)

	req, err := http.NewRequest("POST", rs.endpoint, data)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(rs.accountSID, rs.authToken)
	resp, err := rs.httpClient.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return errors.New("failed to send SMS")
	}

	return nil
}

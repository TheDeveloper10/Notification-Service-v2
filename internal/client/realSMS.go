package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"notification-service/internal/config"
	"notification-service/internal/util"
	"strings"
)

func newRealSMSClientFromConfig(conf *config.SMSConfig) *realSMS {
	c := realSMS{
		httpClient:        &http.Client{},
		endpoint:          fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", conf.AccountSID),
		queryParamsFormat: fmt.Sprintf("From=%s&To=%%s&Body=%%s", conf.FromPhoneNumber),
		accountSID:        conf.AccountSID,
		authToken:         conf.AuthToken,
	}

	util.Logger.Info().Msg("Initialized SMS Client")
	return &c
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
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Failed to send SMS. Status code: %d", resp.StatusCode)
		}
		return fmt.Errorf("Failed to send SMS. Response body: %s", string(bodyBytes))
	}

	return nil
}

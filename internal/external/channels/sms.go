package channels

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/fuadop/sendchamp"
)

type SMS struct {
	To      []string
	Message string
}

type Response struct {
	Success bool
	Message string
}

func SendOTP(to string) (*Response, error) {
	client := sendchamp.NewClient(config.GetConf().ProdSmsPublicKey, sendchamp.ModeTest)

	verifyPayload := sendchamp.SendOTPPayload{
		Channel:              sendchamp.OTPChannelSMS,
		Sender:               "Everest",
		TokenType:            sendchamp.OTPTokenTypeAlphaNumeric,
		TokenLength:          "6",
		ExpirationTime:       5,
		CustomerMobileNumber: to,
		MetaData:             nil,
	}

	s, err := client.NewVerification().SendOTP(verifyPayload)
	if err != nil {
		log.Error("otp err: ", err)
		return nil, err
	}
	var response Response
	if s.Code == "200" {
		response.Success = true
		response.Message = "OTP sent successfully"
		return &response, nil
	} else {
		log.Error("otp err: ", s.Message, s.Code)
		return nil, errors.New(s.Message)
	}
}

func ConfirmOtp(code string, reference string) error {
	client := sendchamp.NewClient(config.GetConf().ProdSmsPublicKey, sendchamp.ModeTest)

	_, err := client.NewVerification().ConfirmOTP(code, reference)
	if err != nil {
		log.Error("otp err: ", err)
		return err
	}
	return nil
}

func SendSMS(sms SMS) (*Response, error) {
	client := sendchamp.NewClient(config.GetConf().ProdSmsPublicKey, sendchamp.ModeTest)

	res, err := client.NewSms().Send("Everest", sms.To, sms.Message, sendchamp.RouteInternational)

	if err != nil {
		log.Error("sms", err)
		return nil, err
	}

	return &Response{
		Success: true,
		Message: res.Message,
	}, nil
}

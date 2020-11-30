package test

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoAuthHelper struct {
	Region           string
	UserPoolID       string
	UserPoolClientID string
}

func (h *CognitoAuthHelper) SignUp(email string, password string) (*cognitoidentityprovider.SignUpOutput, error) {
	newSession := session.Must(session.NewSession())

	svc := cognitoidentityprovider.New(newSession, aws.NewConfig().WithRegion(h.Region))

	paramsSignUp := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(h.UserPoolClientID),
		Password: aws.String(password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
		Username: aws.String(email),
	}

	respSignUp, errSignUp := svc.SignUp(paramsSignUp)
	if errSignUp != nil {
		fmt.Println(errSignUp.Error())
		return nil, errSignUp
	}

	paramsConfirm := &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String(h.UserPoolID),
		Username:   aws.String(email),
	}

	_, errConfirm := svc.AdminConfirmSignUp(paramsConfirm)
	if errConfirm != nil {
		return nil, errConfirm
	}

	return respSignUp, nil
}

func (h *CognitoAuthHelper) SignIn(email string, password string) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	newSession := session.Must(session.NewSession())

	svc := cognitoidentityprovider.New(newSession, aws.NewConfig().WithRegion(h.Region))

	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(email),
			"PASSWORD": aws.String(password),
		},
		ClientId: aws.String(h.UserPoolClientID),
	}

	resp, err := svc.InitiateAuth(params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *CognitoAuthHelper) DeleteUser(accessToken string) error {
	newSession := session.Must(session.NewSession())

	svc := cognitoidentityprovider.New(newSession, aws.NewConfig().WithRegion(h.Region))

	paramsDeleteUser := &cognitoidentityprovider.DeleteUserInput{
		AccessToken: &accessToken,
	}

	_, err := svc.DeleteUser(paramsDeleteUser)
	if err != nil {
		return err
	}

	return nil
}

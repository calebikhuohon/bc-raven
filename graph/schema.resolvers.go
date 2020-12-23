package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"buycoins-raven/graph/generated"
	"buycoins-raven/graph/model"
	"buycoins-raven/internal/auth"
	database "buycoins-raven/internal/pkg/db/firestore"
	"buycoins-raven/internal/pkg/jwt"
	"buycoins-raven/internal/providers/circle"
	"buycoins-raven/internal/transactions"
	"buycoins-raven/internal/users"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.NewTransaction) (*model.Transaction, error) {
	user := auth.ForContext(ctx)
	log.Println(user)
	if user == nil {
		return &model.Transaction{}, fmt.Errorf("access denied")
	}

	sessionId, err := users.HashPassword("randomchar")
	if err != nil {
		return nil, err
	}

	ip, err := circle.GetIPAddress()
	if err != nil {
		return nil, err
	}

	card := circle.Card{
		IdempotencyKey: strings.Replace(uuid.New().String(), "-", "", -1),
		KeyId:          "key1",
		CardDetails: circle.CardDetails{
			Number: input.CardNumber,
			Cvv:    input.Cvv,
		},
		BillingDetails: circle.Billing{
			Name:       input.Name,
			City:       input.City,
			Country:    input.Country,
			Line1:      input.Address,
			PostalCode: input.PostalCode,
			ExpMonth:   int32(input.ExpMonth),
			ExpYear:    int32(input.ExpYear),
		},
		Metadata: circle.Metadata{
			Email:       input.Email,
			PhoneNumber: input.PhoneNumber,
			SessionId:   sessionId,
			IpAddress:   ip,
		},
	}

	log.Println(card)

	_, err = circle.CreateCard(&card)
	if err != nil {
		return nil, err
	}

	idemkey := strings.Replace(uuid.New().String(), "-", "", -1)

	p := circle.Payment{
		IdempotencyKey: idemkey,
		KeyId:          "key1",
		Metadata: circle.Metadata{
			Email:       input.Email,
			PhoneNumber: input.PhoneNumber,
			SessionId:   sessionId,
			IpAddress:   ip,
		},
		Amount: circle.PaymentAmount{
			Amount:   input.Amount,
			Currency: input.Currency,
		},
		Verification: "cvv",
		Source: circle.Source{
			Id:   "",
			Type: "card",
		},
		Description: input.Description,
		PaymentData: circle.PaymentData{
			Cvv: input.Cvv,
		},
	}

	payment, err := circle.MakeCardPayment(&p)
	if err != nil {
		return nil, err
	}

	var transaction transactions.Transaction

	transaction.User = *user
	transaction.SessionId = sessionId
	transaction.TrackingRef = payment["Data"].(map[string]interface{})["trackingRef"].(string)
	transaction.ChargeSuccessful = false
	transaction.WalletUpdated = false
	transaction.TransactionDate = time.Now().String()
	client, ctxx, err := database.ConnectToFirebase()
	if err != nil {
		return nil, err
	}

	transaction.ID = strings.Replace(uuid.New().String(), "-", "", -1)

	if err := transactions.Save(ctxx, client, user.ID, transaction); err != nil {
		return nil, err
	}

	return &model.Transaction{
		ID: transaction.ID,
		User: &model.User{
			ID:       user.ID,
			Username: user.UserName,
		},
		ChargeSuccessful: transaction.ChargeSuccessful,
		WalletUpdated:    transaction.WalletUpdated,
		TransactionDate:  transaction.TransactionDate,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserAuth) (string, error) {
	var user users.User
	user.UserName = input.Username
	user.Password = input.Password

	uid := uuid.New()
	user.ID = strings.Replace(uid.String(), "-", "", -1)

	client, ctxx, err := database.ConnectToFirebase()
	if err != nil {
		return "", err
	}

	if err := users.Create(ctxx, *client, &user); err != nil {
		log.Println(err)
		return "", err
	}
	token, err := jwt.GenerateToken(user.UserName)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.UserAuth) (string, error) {
	var user users.User

	user.UserName = input.Username
	user.Password = input.Password

	client, ctxx, err := database.ConnectToFirebase()
	if err != nil {
		return "", err
	}
	correct := user.Authenticate(ctxx, *client)

	if !correct {
		return "", fmt.Errorf("wrong username or password")
	}
	token, err := jwt.GenerateToken(user.UserName)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.NewRefreshToken) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return []*model.Transaction{}, fmt.Errorf("access denied")
	}

	var ts []*model.Transaction

	client, ctx, err := database.ConnectToFirebase()
	if err != nil {
		return nil, err
	}

	dbtransactions, err := transactions.GetAll(user.ID, ctx, client)
	if err != nil {
		return nil, err
	}
	for _, t := range dbtransactions {
		log.Println(t)
		ts = append(ts, &model.Transaction{
			ID: t["ID"].(string),
			User: &model.User{
				ID:       t["User"].(map[string]interface{})["ID"].(string),
				Username: t["User"].(map[string]interface{})["UserName"].(string),
			},
			TransactionDate:  t["TransactionDate"].(string),
			ChargeSuccessful: t["ChargeSuccessful"].(bool),
			WalletUpdated:    t["WalletUpdated"].(bool),
		})
	}
	return ts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

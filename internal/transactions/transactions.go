package transactions

import (
	"buycoins-raven/internal/users"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
)

type Transaction struct {
	ID               string     `json:"id"`
	User             users.User `json:"user"`
	TransactionDate  string     `json:"transaction_date"`
	ChargeSuccessful bool       `json:"charge_successful"`
	WalletUpdated    bool       `json:"wallet_updated"`
	TransactionId    string     `json:"transaction_id"`
	IdempotencyKey   string     `json:"idempotency_key"`
	TrackingRef      string     `json:"tracking_ref"`
	SessionId        string     `json:"session_id"`
}

func Save(ctx context.Context, client *firestore.Client, id string, data Transaction) error {
	_, _, err := client.Collection(id).Doc("transactionDoc").Collection("transactions").Add(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func GetAll(id string, ctx context.Context, client *firestore.Client) ([]map[string]interface{}, error) {
	iter := client.Collection(id).Doc("transactionDoc").Collection("transactions").Documents(ctx)
	var docs []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}
		docs = append(docs, doc.Data())
	}
	return docs, nil
}

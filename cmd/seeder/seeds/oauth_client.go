package seeds

import (
	"context"
	"github.com/evenyosua18/auth2/app/model"
	"github.com/evenyosua18/auth2/app/repository/authdb/oauthclient"
	"github.com/evenyosua18/auth2/app/utils/db"
	"github.com/evenyosua18/auth2/app/utils/str"
	"golang.org/x/crypto/bcrypt"
)

func GenerateOauthClient(db *db.MongoConnection) error {
	// initialize oauth client repository
	repo := oauthclient.NewOauthClientRepository(db)

	// initialize context background
	ctx := context.Background()

	// secret
	secret, err := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	// insert oauth client
	if err := repo.InsertOauthClient(ctx, model.OauthClientModel{
		ClientId:     str.GenerateString(16, ""),
		ClientSecret: string(secret),
		Uri:          "",
		Scopes:       "",
		ClientType:   "",
	}); err != nil {
		return err
	}

	return nil
}

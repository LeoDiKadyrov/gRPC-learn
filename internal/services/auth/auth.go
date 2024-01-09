package auth

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"sso/internal/domain/models"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	log *slog.Logger
	usrSaver UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (user models.User, err error)
	IsAdmin(ctx context.Context, userID int64) (isAdmin bool, err error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// New returns a new instance of the Auth service
func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		usrSaver: userSaver,
		usrProvider: userProvider,
		log: log,
		appProvider: appProvider,
		tokenTTL: tokenTTL,
	}
}

func (a *Auth) Login (
	ctx context.Context,
	email string,
	password string,
	appID int,
) (string, error) {
	panic("not implemented")
}

func (a *Auth) RegisterUser (
	ctx context.Context,
	email string,
	password string,
) (int64, error) {
	const op = "auth.RegisterUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)
	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to hash password", slog.Any("err", err))
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.usrSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Error("failed to save user", slog.Any("err", err))
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("user registered")
	return id, nil
}

func (a *Auth) IsAdmin (
	ctx context.Context,
	userID int,
) (bool, error) {
	panic("not implemented")
}
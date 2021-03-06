package user

import (
	"context"
	"fmt"
	"regexp"

	"github.com/51st-state/api/pkg/event"
	"github.com/51st-state/api/pkg/rbac"

	"github.com/pkg/errors"

	"github.com/51st-state/api/pkg/bcrypt"
)

// Manager of user objects
//go:generate counterfeiter -o ./mocks/manager.go . Manager
type Manager interface {
	Get(ctx context.Context, id Identifier) (Complete, error)
	GetByGameSerialHash(ctx context.Context, hash string) (Complete, error)
	GetByWCFUserID(ctx context.Context, wcfUserID WCFUserID) (Complete, error)
	Create(ctx context.Context, inc Incomplete) (Complete, error)
	Delete(ctx context.Context, id Identifier) error
	GetWCFInfo(ctx context.Context, name string) (*WCFUserInfo, error)
	Update(ctx context.Context, c Complete) error
	CheckPassword(ctx context.Context, id Identifier, incPw IncompletePassword) error
	GetRoles(ctx context.Context, id Identifier) (rbac.AccountRoles, error)
	SetRoles(ctx context.Context, id Identifier, roles rbac.AccountRoles) error
}

type manager struct {
	repository    Repository
	wcfRepository WCFRepository
	event         *event.Producer
	rbac          rbac.Control
}

// NewManager for user objects
//go:generate protoc -I./../../../../../../ -I ./proto --go_out=plugins=grpc:./proto ./proto/manager.proto
func NewManager(r Repository, wcf WCFRepository, prod *event.Producer, rb rbac.Control) Manager {
	return &manager{
		r,
		wcf,
		prod,
		rb,
	}
}

var errInvalidUUID = errors.New("invalid user uuid given")

// Get an user object
func (m *manager) Get(ctx context.Context, id Identifier) (Complete, error) {
	if id.UUID() == "" {
		return nil, errInvalidUUID
	}

	c, err := m.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	wcfInfo, err := m.wcfRepository.GetInfo(ctx, c.Data().WCFUserID)
	if err != nil {
		return nil, err
	}

	c.Data().WCFUsername = wcfInfo.Username
	c.Data().WCFEmail = wcfInfo.Email

	return c, nil
}

// GetByGameSerialHash returns a user filtered by its unique game serial hash
func (m *manager) GetByGameSerialHash(ctx context.Context, hash string) (Complete, error) {
	if hash == "" {
		return nil, errInvalidGameSerialHash
	}

	c, err := m.repository.GetByGameSerialHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	wcfInfo, err := m.wcfRepository.GetInfo(ctx, c.Data().WCFUserID)
	if err != nil {
		return nil, err
	}

	c.Data().WCFUsername = wcfInfo.Username
	c.Data().WCFEmail = wcfInfo.Email

	return c, nil
}

// GetByWCFUserID returns an user filtered by its wcf user id
func (m *manager) GetByWCFUserID(ctx context.Context, wcfUserID WCFUserID) (Complete, error) {
	if wcfUserID == 0 {
		return nil, errInvalidWCFUserID
	}

	c, err := m.repository.GetByWCFUserID(ctx, wcfUserID)
	if err != nil {
		return nil, err
	}

	wcfInfo, err := m.wcfRepository.GetInfo(ctx, c.Data().WCFUserID)
	if err != nil {
		return nil, err
	}

	c.Data().WCFUsername = wcfInfo.Username
	c.Data().WCFEmail = wcfInfo.Email

	return c, nil
}

var errInvalidWCFUserID = errors.New("invalid woltlab community framework user id")

// Create an user object
func (m *manager) Create(ctx context.Context, inc Incomplete) (Complete, error) {
	if inc.Data().WCFUserID == 0 {
		return nil, errInvalidWCFUserID
	}

	wcfInfo, err := m.wcfRepository.GetInfo(ctx, inc.Data().WCFUserID)
	if err != nil {
		return nil, err
	}

	c, err := m.repository.Create(ctx, inc)
	if err != nil {
		return nil, err
	}

	c.Data().WCFUsername = wcfInfo.Username
	c.Data().WCFEmail = wcfInfo.Email

	return c, m.event.Produce(ctx, CreatedEventID, &CreatedEvent{
		&event.PayloadMeta{
			Version: "1",
		},
		c,
	})
}

var (
	errInvalidEmailFormat = errors.New("invalid email format")
	emailRegexp           = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func (m *manager) checkEmail(ctx context.Context, email string) error {
	if !emailRegexp.MatchString(email) {
		return errInvalidEmailFormat
	}

	return nil
}

// Delete an user object
func (m *manager) Delete(ctx context.Context, id Identifier) error {
	if id.UUID() == "" {
		return errInvalidUUID
	}

	if err := m.repository.Delete(ctx, id); err != nil {
		return err
	}

	return m.event.Produce(ctx, DeletedEventID, &DeletedEvent{
		&event.PayloadMeta{
			Version: "1",
		},
		id,
	})
}

// GetWCFInfo returns the info of a wcf user object
func (m *manager) GetWCFInfo(ctx context.Context, name string) (*WCFUserInfo, error) {
	if emailRegexp.MatchString(name) {
		return m.wcfRepository.GetInfoByEmail(ctx, name)
	}

	return m.wcfRepository.GetInfoByUsername(ctx, name)
}

var errInvalidGameSerialHash = errors.New("invalid game serial hash")

// Update an user objects data
func (m *manager) Update(ctx context.Context, c Complete) error {
	if c.UUID() == "" {
		return errInvalidUUID
	}

	if c.Data().WCFUserID == 0 {
		return errInvalidWCFUserID
	}

	if _, err := m.wcfRepository.GetInfo(ctx, c.Data().WCFUserID); err != nil {
		return err
	}

	if err := m.repository.Update(ctx, c); err != nil {
		return err
	}

	return m.event.Produce(ctx, UpdatedEventID, &UpdatedEvent{
		&event.PayloadMeta{
			Version: "1",
		},
		c,
	})
}

// CheckPassword of a user
func (m *manager) CheckPassword(ctx context.Context, id Identifier, incPw IncompletePassword) error {
	if id.UUID() == "" {
		return errInvalidUUID
	}

	compl, err := m.Get(ctx, id)
	if err != nil {
		return err
	}

	wcfInfo, err := m.wcfRepository.GetInfo(ctx, compl.Data().WCFUserID)
	if err != nil {
		return err
	}

	pwFirstHash, err := getFirstPasswordHash(wcfInfo.Password.Hash(), []byte(incPw.Password()))
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(wcfInfo.Password.Hash(), pwFirstHash)
}

const (
	majorBcryptVersion = '2'
	minorBcryptVersion = 'a'
)

func getFirstPasswordHash(hash, password []byte) ([]byte, error) {
	hashInfo, err := bcrypt.NewFromHash(hash)
	if err != nil {
		return nil, err
	}

	pwHashOnly, err := bcrypt.Bcrypt(password, hashInfo.Cost, hashInfo.Salt)
	if err != nil {
		return nil, err
	}

	return bcrypt.NewHash(
		pwHashOnly,
		hashInfo.Salt,
		hashInfo.Cost,
		majorBcryptVersion,
		minorBcryptVersion,
	).Hash(), nil
}

// GetRoles of a user
func (m *manager) GetRoles(ctx context.Context, id Identifier) (rbac.AccountRoles, error) {
	if id.UUID() == "" {
		return nil, errInvalidUUID
	}

	roles, err := m.rbac.GetAccountRoles(ctx, rbac.AccountID(fmt.Sprintf(
		"user/%s",
		id.UUID(),
	)))
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// SetRoles of a user
func (m *manager) SetRoles(ctx context.Context, id Identifier, roles rbac.AccountRoles) error {
	if id.UUID() == "" {
		return errInvalidUUID
	}

	return m.rbac.SetAccountRoles(ctx, rbac.AccountID(fmt.Sprintf(
		"user/%s",
		id.UUID(),
	)), roles)
}

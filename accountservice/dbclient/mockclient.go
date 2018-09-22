package dbclient

import (
	"github.com/navono/go-microservice/accountservice/model"
	"github.com/stretchr/testify/mock"
)

// MockBoltClient is a mock implementation of a datastore client for testing purposes.
// Instead of the bolt.DB pointer, we're just putting a generic mock object from
// strechr/testify
type MockBoltClient struct {
	mock.Mock
}

// QueryAccount mock for interface method
// From here, we'll declare three functions that makes our MockBoltClient fufill the
// interface IBoltClient
func (m *MockBoltClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

//OpenBoltDb mock for interface method
func (m *MockBoltClient) OpenBoltDb() {
	// Does nothing
}

// Seed mock for interface method
func (m *MockBoltClient) Seed() {
	// Does nothing
}

// Check mock for interface method
func (m *MockBoltClient) Check() bool {
	args := m.Mock.Called()
	return args.Get(0).(bool)
}

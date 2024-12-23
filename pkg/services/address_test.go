package services

import (
	"errors"
	"testing"

	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/model"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepo is a mock implementation of the repository interface
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateUser(user *model.User) (uint, error) {
	args := m.Called(user)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockRepo) FindUserByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockRepo) FindUserByID(userID uint) (*model.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockRepo) UpdateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepo) GetUserList() ([]*model.User, error) {
	args := m.Called()
	return args.Get(0).([]*model.User), args.Error(1)
}

func (m *MockRepo) CreateAddress(address *model.Address) error {
	args := m.Called(address)
	return args.Error(0)
}

func (m *MockRepo) EditAddress(address *model.Address) error {
	args := m.Called(address)
	return args.Error(0)
}

func (m *MockRepo) FindAddress(userID uint) (*model.Address, error) {
	args := m.Called(userID)
	return args.Get(0).(*model.Address), args.Error(1)
}

func (m *MockRepo) GetAllAddresses(userID uint) (*[]model.Address, error) {
	args := m.Called(userID)
	return args.Get(0).(*[]model.Address), args.Error(1)
}

func (m *MockRepo) RemoveAddress(addressID, userID uint) error {
	args := m.Called(addressID, userID)
	return args.Error(0)
}

// TestAddAddressService tests the AddAddressService method
func TestAddAddressService(t *testing.T) {
	// Initialize the mock repository for mock test
	mockRepo := new(MockRepo)
	userService := &UserService{Repo: mockRepo}

	// Test data
	address := &pb.Address{
		House:   "123 Elm St",
		Street:  "Main Street",
		City:    "Gotham",
		Zip:     12345,
		State:   "NY",
		User_ID: 1,
	}

	// Mock the behavior of the CreateAddress method
	mockRepo.On("CreateAddress", mock.AnythingOfType("*model.Address")).Return(nil)

	// Call the AddAddressService method
	response, err := userService.AddAddressService(address)

	// Assert that no error occurred and the response is successful
	assert.NoError(t, err)
	assert.Equal(t, pb.Response_OK, response.Status)
	assert.Equal(t, "Address added successfully", response.Message)

	// Assert that CreateAddress was called once
	mockRepo.AssertExpectations(t)
}

// TestAddAddressService_Error tests the AddAddressService method when an error occurs
func TestAddAddressService_Error(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(MockRepo)
	userService := &UserService{Repo: mockRepo}

	// Test data
	address := &pb.Address{
		House:   "123 Elm St",
		Street:  "Main Street",
		City:    "Gotham",
		Zip:     12345,
		State:   "NY",
		User_ID: 1,
	}

	// Mock the behavior of the CreateAddress method to return an error
	mockRepo.On("CreateAddress", mock.AnythingOfType("*model.Address")).Return(errors.New("database error"))

	// Call the AddAddressService method
	response, err := userService.AddAddressService(address)

	// Assert that an error occurred and the response contains an error message
	assert.Error(t, err)
	assert.Equal(t, pb.Response_ERROR, response.Status)
	assert.Equal(t, "Error adding address", response.Message)

	// Assert that CreateAddress was called once
	mockRepo.AssertExpectations(t)
}

// TestEditAddressService tests the EditAddressService method
func TestEditAddressService(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(MockRepo)
	userService := &UserService{Repo: mockRepo}

	// Test data
	address := &pb.Address{
		ID:      1,
		House:   "123 Elm St",
		Street:  "Main Street",
		City:    "Gotham",
		Zip:     12345,
		State:   "NY",
		User_ID: 1,
	}

	// Mock the behavior of the EditAddress method
	mockRepo.On("EditAddress", mock.AnythingOfType("*model.Address")).Return(nil)

	// Call the EditAddressService method
	editedAddress, err := userService.EditAddressService(address)

	// Assert that no error occurred and the address was edited
	assert.NoError(t, err)
	assert.Equal(t, address, editedAddress)

	// Assert that EditAddress was called once
	mockRepo.AssertExpectations(t)
}

// TestEditAddressService_Error tests the EditAddressService method when an error occurs
func TestEditAddressService_Error(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(MockRepo)
	userService := &UserService{Repo: mockRepo}

	// Test data
	address := &pb.Address{
		ID:      1,
		House:   "123 Elm St",
		Street:  "Main Street",
		City:    "Gotham",
		Zip:     12345,
		State:   "NY",
		User_ID: 1,
	}

	// Mock the behavior of the EditAddress method to return an error
	mockRepo.On("EditAddress", mock.AnythingOfType("*model.Address")).Return(errors.New("database error"))

	// Call the EditAddressService method
	editedAddress, err := userService.EditAddressService(address)

	// Assert that an error occurred and the returned address is nil
	assert.Error(t, err)
	assert.Nil(t, editedAddress)

	// Assert that EditAddress was called once
	mockRepo.AssertExpectations(t)
}

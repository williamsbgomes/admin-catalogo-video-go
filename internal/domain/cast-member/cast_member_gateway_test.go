package castmember

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCastMemberGateway struct {
	mock.Mock
}

func (m *MockCastMemberGateway) Create(castMember *CastMember) (*CastMember, error) {
	args := m.Called(castMember)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*CastMember), nil
}

func (m *MockCastMemberGateway) Update(castMember *CastMember) (*CastMember, error) {
	args := m.Called(castMember)
	if err := args.Error(1); err != nil {
		return nil, err
	}
	updated := args.Get(0).(*CastMember)
	return updated, nil
}

func (m *MockCastMemberGateway) DeleteByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCastMemberGateway) FindByID(id string) (*CastMember, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*CastMember), nil
}

func (m *MockCastMemberGateway) FindAll() ([]CastMember, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]CastMember), nil
}

func TestMockCastMemberGateway_Create(t *testing.T) {
	m := new(MockCastMemberGateway)
	castMember := &CastMember{ID: "1"}
	m.On("Create", castMember).Return(castMember, nil)

	result, err := m.Create(castMember)

	assert.NoError(t, err)
	assert.Equal(t, castMember, result)
	m.AssertExpectations(t)
}

func TestMockCastMemberGateway_Create_Error(t *testing.T) {
	m := new(MockCastMemberGateway)
	castMember := &CastMember{ID: "1"}
	expectedErr := errors.New("create error")
	m.On("Create", castMember).Return(nil, expectedErr)

	result, err := m.Create(castMember)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	m.AssertExpectations(t)
}

func TestMockCastMemberGateway_Update(t *testing.T) {
	m := new(MockCastMemberGateway)
	castMember := &CastMember{ID: "1"}
	m.On("Update", castMember).Return(castMember, nil)

	result, err := m.Update(castMember)

	assert.NoError(t, err)
	assert.Equal(t, castMember, result)
	m.AssertExpectations(t)
}

func TestMockCastMemberGateway_Update_Error(t *testing.T) {
	m := new(MockCastMemberGateway)
	castMember := &CastMember{ID: "1"}
	expectedErr := errors.New("update error")
	m.On("Update", castMember).Return(nil, expectedErr)

	result, err := m.Update(castMember)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	m.AssertExpectations(t)
}

func TestMockCastMemberGateway_DeleteByID(t *testing.T) {
	m := new(MockCastMemberGateway)
	id := "1"
	m.On("DeleteByID", id).Return(nil)

	err := m.DeleteByID(id)

	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestMockCastMemberGateway_DeleteByID_Error(t *testing.T) {
	m := new(MockCastMemberGateway)
	id := "1"
	expectedErr := errors.New("delete error")
	m.On("DeleteByID", id).Return(expectedErr)

	err := m.DeleteByID(id)
	assert.Error(t, err)
}

func TestMockCastMemberGateway_FindByID(t *testing.T) {
	m := new(MockCastMemberGateway)
	id := "1"
	castMember := &CastMember{ID: id}
	m.On("FindByID", id).Return(castMember, nil)

	result, err := m.FindByID(id)

	assert.NoError(t, err)
	assert.Equal(t, castMember, result)
	m.AssertExpectations(t)
}

func TestMockCastMemberGateway_FindByID_Error(t *testing.T) {
	m := new(MockCastMemberGateway)
	id := "1"
	expectedErr := errors.New("find error")
	m.On("FindByID", id).Return(nil, expectedErr)

	result, err := m.FindByID(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	m.AssertExpectations(t)
}

func TestMockCastMemberGateway_FindAll(t *testing.T) {
	m := new(MockCastMemberGateway)
	castMembers := []CastMember{{ID: "1"}, {ID: "2"}}
	m.On("FindAll").Return(castMembers, nil)

	result, err := m.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, castMembers, result)
	m.AssertExpectations(t)
}

func TestMockCastMemberGateway_FindAll_Error(t *testing.T) {
	m := new(MockCastMemberGateway)
	expectedErr := errors.New("find all error")
	m.On("FindAll").Return(nil, expectedErr)

	result, err := m.FindAll()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	m.AssertExpectations(t)
}

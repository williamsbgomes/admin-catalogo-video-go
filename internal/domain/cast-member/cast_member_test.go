package castmember_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	castmember "github.com/williamsbgomes/admin-catalogo-video-go/internal/domain/cast-member"
)

const (
	nameEmptyErrorMessage    = "'name' should not be empty"
	nameLengthErrorMessage   = "'name' must be between 3 and 255 characters"
	validCategoryDescription = "A categoria mais assistida"
)

func TestGivenAnEmptyID_WhenCreateANewCastMember_ThenShouldReceiveAnError(t *testing.T) {
	categoryEntity := castmember.CastMember{ID: ""}
	err := categoryEntity.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'id' should not be empty")
}

func TestGivenAnEmptyName_WhenCreateANewCastMember_ThenShouldReceiveAnError(t *testing.T) {
	categoryEntity := castmember.CastMember{ID: "1234", Name: ""}
	err := categoryEntity.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), nameEmptyErrorMessage)
}

func TestGivenAnInvalidNameLengthLessThan3_WhenCreateANewCastMember_ThenShouldReceiveAnError(t *testing.T) {
	categoryEntity := castmember.CastMember{ID: "1234", Name: "ab"}
	err := categoryEntity.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), nameLengthErrorMessage)
}

func TestGivenAnInvalidNameLengthMoreThan255_WhenCreateANewCastMember_ThenShouldReceiveAnError(t *testing.T) {
	categoryEntity := castmember.CastMember{
		ID:   "1234",
		Name: strings.Repeat("a", 256),
	}
	err := categoryEntity.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), nameLengthErrorMessage)
}

func TestGivenAValidParams_WhenCallNewCastMember_ThenInstantiateACastMember(t *testing.T) {
	expectedName := "Vin Diesel"
	expectedType := castmember.Actor

	castMember, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)

	assert.NoError(t, err)
	assert.NotNil(t, castMember)
	assert.NotEmpty(t, castMember.ID)
	assert.Equal(t, expectedName, castMember.Name)
	assert.Equal(t, expectedType, castMember.Type)
	assert.NotZero(t, castMember.CreatedAt)
	assert.NotZero(t, castMember.UpdatedAt)
}

func TestGivenAnInvalidEmptyName_WhenCallNewCastMember_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := ""
	expectedType := castmember.Actor

	expectedErrorMessage := nameEmptyErrorMessage

	_, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
}

func TestGivenAnInvalidNameLengthLessThan3_WhenCallNewCastMember_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := "ab"
	expectedType := castmember.Director

	expectedErrorMessage := nameLengthErrorMessage

	_, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
}

func TestGivenAnInvalidNameLengthMoreThan255_WhenCallNewCastMember_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := strings.Repeat("a", 256)
	expectedType := castmember.Director

	expectedErrorMessage := nameLengthErrorMessage

	_, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
}

func TestGivenAnInvalidEmptyType_WhenCallNewCastMember_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := "Keanu Reeves"
	expectedType := castmember.CastMemberType("FLAVOR")

	expectedErrorMessage := "'type' must be either 'ACTOR' or 'DIRECTOR'"

	_, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
}

func TestGivenAValidCastMember_WhenCallUpdate_ThenReturnUpdatedCastMember(t *testing.T) {
	expectedName := "Jonhny Depp"
	expectedType := castmember.Director

	expectedUpdatedName := "Jonhny Depp - Atualizado"
	expectedUpdatedType := castmember.Actor

	castMember, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)

	assert.NoError(t, err)
	assert.NotNil(t, castMember)
	assert.NotEmpty(t, castMember.ID)
	assert.Equal(t, expectedName, castMember.Name)
	assert.Equal(t, expectedType, castMember.Type)
	assert.NotZero(t, castMember.CreatedAt)
	assert.NotZero(t, castMember.UpdatedAt)

	castMember.Update(expectedUpdatedName, expectedUpdatedType)

	assert.NotEmpty(t, castMember.ID)
	assert.Equal(t, expectedUpdatedName, castMember.Name)
	assert.Equal(t, expectedUpdatedType, castMember.Type)
	assert.NotZero(t, castMember.CreatedAt)
	assert.NotZero(t, castMember.UpdatedAt)
}

func TestGivenAValidCastMember_WhenCallUpdateWithInvalidEmptyName_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := "Filmes"
	expectedType := castmember.Actor

	castMember, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)

	assert.NoError(t, err)
	assert.NotNil(t, castMember)
	assert.NotEmpty(t, castMember.ID)

	// Test with empty name
	err = castMember.Update("", expectedType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), nameEmptyErrorMessage)
}

func TestGivenAValidCastMember_WhenCallUpdateWithInvalidNameLengthLessThan3_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := "Filmes"
	expectedType := castmember.Actor

	castMember, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)

	assert.NoError(t, err)
	assert.NotNil(t, castMember)
	assert.NotEmpty(t, castMember.ID)

	err = castMember.Update("ab", expectedType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), nameLengthErrorMessage)
}

func TestGivenAValidCastMember_WhenCallUpdateWithInvalidNameLengthMoreThan255_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := "Filmes"
	expectedType := castmember.Actor

	castMember, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)

	assert.NoError(t, err)
	assert.NotNil(t, castMember)
	assert.NotEmpty(t, castMember.ID)

	err = castMember.Update(strings.Repeat("a", 256), expectedType)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), nameLengthErrorMessage)
}

// Test with invalid type
func TestGivenAValidCastMember_WhenCallUpdateWithInvalidType_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := "Filmes"
	expectedType := castmember.Actor

	expectedErrorMessage := "'type' must be either 'ACTOR' or 'DIRECTOR'"

	castMember, err := castmember.NewCastMember(
		expectedName,
		expectedType,
	)
	assert.NoError(t, err)
	assert.NotNil(t, castMember)
	assert.NotEmpty(t, castMember.ID)

	err = castMember.Update("Steven Seagal", "INVALID")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
}

package category_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamsbgomes/admin-catalogo-video-go/internal/entity/category"
)

func TestGivenAnEmptyID_WhenCreateANewCategory_ThenShouldReceiveAnError(t *testing.T) {
	categoryEntity := category.Category{ID: ""}
	err := categoryEntity.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'id' should not be empty")
}

func TestGivenAnEmptyName_WhenCreateANewCategory_ThenShouldReceiveAnError(t *testing.T) {
	categoryEntity := category.Category{ID: "1234", Name: ""}
	err := categoryEntity.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'name' should not be empty")
}

func TestGivenAnInvalidNameLengthLessThan3_WhenCreateANewCategory_ThenShouldReceiveAnError(t *testing.T) {
	categoryEntity := category.Category{ID: "1234", Name: "ab"}
	err := categoryEntity.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'name' must be between 3 and 255 characters")
}

func TestGivenAnInvalidNameLengthMoreThan255_WhenCreateANewCategory_ThenShouldReceiveAnError(t *testing.T) {
	categoryEntity := category.Category{
		ID:   "1234",
		Name: strings.Repeat("a", 256),
	}
	err := categoryEntity.IsValid()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'name' must be between 3 and 255 characters")
}

func TestGivenAValidParams_WhenCallNewCategory_ThenInstantiateACategory(t *testing.T) {
	expectedName := "Filmes"
	expectedDescription := "A categoria mais assistida"
	expectedActive := true

	categoryEntity, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)

	assert.NoError(t, err)
	assert.NotNil(t, categoryEntity)
	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedName, categoryEntity.Name)
	assert.Equal(t, expectedDescription, categoryEntity.Description)
	assert.Equal(t, expectedActive, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.Nil(t, categoryEntity.DeletedAt)
}

func TestGivenAnInvalidEmptyName_WhenCallNewCategory_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := ""
	expectedDescription := "A categoria mais assistida"
	expectedActive := true

	expectedErrorMessage := "'name' should not be empty"

	_, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
}

func TestGivenAnInvalidNameLengthLessThan3_WhenCallNewCategory_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := "ab"
	expectedDescription := "A categoria mais assistida"
	expectedActive := true

	expectedErrorMessage := "'name' must be between 3 and 255 characters"

	_, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
}

func TestGivenAnInvalidNameLengthMoreThan255_WhenCallNewCategory_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := strings.Repeat("a", 256)
	expectedDescription := "A categoria mais assistida"
	expectedActive := true

	expectedErrorMessage := "'name' must be between 3 and 255 characters"

	_, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrorMessage)
}

func TestGivenAValidEmptyDescription_WhenCallNewCategoryAndValidate_ThenShouldReceiveOK(t *testing.T) {
	expectedName := "Filmes"
	expectedDescription := ""
	expectedActive := true

	categoryEntity, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)

	assert.NoError(t, err)
	assert.NotNil(t, categoryEntity)
	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedName, categoryEntity.Name)
	assert.Equal(t, expectedDescription, categoryEntity.Description)
	assert.Equal(t, expectedActive, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.Nil(t, categoryEntity.DeletedAt)
}

func TestGivenAValidFalseIsActive_WhenCallNewCategoryAndValidate_ThenShouldReceiveOK(t *testing.T) {
	expectedName := "Filmes"
	expectedDescription := "A categoria mais assistida"
	expectedActive := false

	categoryEntity, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)

	assert.NoError(t, err)
	assert.NotNil(t, categoryEntity)
	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedName, categoryEntity.Name)
	assert.Equal(t, expectedDescription, categoryEntity.Description)
	assert.Equal(t, expectedActive, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.NotNil(t, categoryEntity.DeletedAt)
}

func TestGivenAValidActiveCategory_whenCallDeactivate_thenReturnCategoryInactivated(t *testing.T) {
	expectedName := "Filmes"
	expectedDescription := "A categoria mais assistida"
	expectedActive := true

	categoryEntity, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)

	assert.NoError(t, err)
	assert.NotNil(t, categoryEntity)
	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedName, categoryEntity.Name)
	assert.Equal(t, expectedDescription, categoryEntity.Description)
	assert.Equal(t, expectedActive, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.Nil(t, categoryEntity.DeletedAt)

	categoryEntity.Deactivate()

	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedName, categoryEntity.Name)
	assert.Equal(t, expectedDescription, categoryEntity.Description)
	assert.False(t, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.NotNil(t, categoryEntity.DeletedAt)
}

func TestGivenAValidInactiveCategory_whenCallActivate_thenReturnCategoryActivated(t *testing.T) {
	expectedName := "Filmes"
	expectedDescription := "A categoria mais assistida"
	expectedActive := false

	categoryEntity, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)

	assert.NoError(t, err)
	assert.NotNil(t, categoryEntity)
	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedName, categoryEntity.Name)
	assert.Equal(t, expectedDescription, categoryEntity.Description)
	assert.Equal(t, expectedActive, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.NotNil(t, categoryEntity.DeletedAt)

	categoryEntity.Activate()

	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedName, categoryEntity.Name)
	assert.Equal(t, expectedDescription, categoryEntity.Description)
	assert.True(t, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.Nil(t, categoryEntity.DeletedAt)
}

func TestGivenAValidCategory_WhenCallUpdate_ThenReturnUpdatedCategory(t *testing.T) {
	expectedName := "Filmes"
	expectedDescription := "A categoria mais assistida"
	expectedActive := true

	expectedUpdatedName := "Filmes Atualizados"
	expectedUpdatedDescription := "A categoria mais assistida - Atualizada"

	categoryEntity, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)

	assert.NoError(t, err)
	assert.NotNil(t, categoryEntity)
	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedName, categoryEntity.Name)
	assert.Equal(t, expectedDescription, categoryEntity.Description)
	assert.Equal(t, expectedActive, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.Nil(t, categoryEntity.DeletedAt)

	categoryEntity.Update(expectedUpdatedName, expectedUpdatedDescription, false)

	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedUpdatedName, categoryEntity.Name)
	assert.Equal(t, expectedUpdatedDescription, categoryEntity.Description)
	assert.False(t, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.NotNil(t, categoryEntity.DeletedAt)
}

func TestGivenAValidCategory_WhenCallUpdateToInactive_ThenReturnUpdatedCategory(t *testing.T) {
	expectedName := "Filmes"
	expectedDescription := "A categoria mais assistida"
	expectedActive := false

	expectedUpdatedName := "Filmes Atualizados"
	expectedUpdatedDescription := "A categoria mais assistida - Atualizada"

	categoryEntity, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)

	assert.NoError(t, err)
	assert.NotNil(t, categoryEntity)
	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedName, categoryEntity.Name)
	assert.Equal(t, expectedDescription, categoryEntity.Description)
	assert.Equal(t, expectedActive, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.NotNil(t, categoryEntity.DeletedAt)

	categoryEntity.Update(expectedUpdatedName, expectedUpdatedDescription, false)

	assert.NotEmpty(t, categoryEntity.ID)
	assert.Equal(t, expectedUpdatedName, categoryEntity.Name)
	assert.Equal(t, expectedUpdatedDescription, categoryEntity.Description)
	assert.False(t, categoryEntity.Active)
	assert.NotZero(t, categoryEntity.CreatedAt)
	assert.NotZero(t, categoryEntity.UpdatedAt)
	assert.NotNil(t, categoryEntity.DeletedAt)
}

func TestGivenAValidCategory_WhenCallUpdateWithInvalidParams_ThenShouldReceiveAnError(t *testing.T) {
	expectedName := "Filmes"
	expectedDescription := "A categoria mais assistida"
	expectedActive := true

	categoryEntity, err := category.NewCategory(
		expectedName,
		expectedDescription,
		expectedActive,
	)

	assert.NoError(t, err)
	assert.NotNil(t, categoryEntity)
	assert.NotEmpty(t, categoryEntity.ID)

	err = categoryEntity.Update("", expectedDescription, expectedActive)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'name' should not be empty")

	err = categoryEntity.Update("ab", "", expectedActive)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'name' must be between 3 and 255 characters")

	err = categoryEntity.Update(strings.Repeat("a", 256), "", expectedActive)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'name' must be between 3 and 255 characters")

	err = categoryEntity.Update(expectedName, expectedDescription, true)
	assert.NoError(t, err)
}

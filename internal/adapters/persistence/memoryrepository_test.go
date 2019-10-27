package persistence_test

import (
	"testing"
	"url-at-minimal-api/internal/adapters/persistence"

	"github.com/stretchr/testify/assert"
)

func TestMemoryRepositorySaveAndFind(t *testing.T) {
	// Given
	repo := persistence.CreateRepository()

	repo.Save("https://www.google.com", "Vsdfb1")
	repo.Save("https://www.microsoft.com", "MasFgr0")

	// When
	count := repo.Count()
	entry1 := repo.Find("Vsdfb1")
	entry2 := repo.Find("MasFgr0")

	// Then
	assert.Equal(t, 2, count)
	assert.Equal(t, "https://www.google.com", entry1)
	assert.Equal(t, "https://www.microsoft.com", entry2)
}

func TestMemoryRepositoryNonExisting(t *testing.T) {
	// Given
	repo := persistence.CreateRepository()

	repo.Save("https://www.google.com", "Vsdfb1")
	repo.Save("https://www.microsoft.com", "MasFgr0")

	// When
	count := repo.Count()
	entry1 := repo.Find("random")

	// Then
	assert.Equal(t, 2, count)
	assert.Equal(t, "", entry1)
}

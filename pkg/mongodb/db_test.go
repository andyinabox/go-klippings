package database

import (
	"github.com/andyinabox/go-klippings-api/pkg/parser"
	"github.com/andyinabox/go-klippings-api/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const testDb = "../../test/tmp/test.db"
const testDataFile = "../../test/data/my_clippings.txt"

func createTestDB() (*Database, error) {
	f, err := os.Open(testDataFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := parser.Parse(f)
	if err != nil {
		return nil, err
	}

	db, err := Open(testDb)
	if err != nil {
		return nil, err
	}

	_, err = db.ProcessParseData(&data)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestOpen(t *testing.T) {
	db, err := Open(testDb)
	require.Nil(t, err)

	err = db.Destroy()
	assert.Nil(t, err)
}

func TestProcessParseData(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	db, err := createTestDB()
	require.Nil(err)

	var titles []types.Title
	db.DB.Find(&titles)
	assert.NotEmpty(titles)
	assert.Equal(len(titles), 3)

	var authors []types.Author
	db.DB.Find(&authors)
	assert.NotEmpty(authors)
	assert.Equal(len(authors), 4)

	var clippings []types.Clipping
	db.DB.Find(&clippings)
	assert.NotEmpty(clippings)
	assert.Equal(len(clippings), 23)

	err = db.Destroy()
}

func TestDuplicates(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	db, err := createTestDB()
	require.Nil(err)

	var initialClippingsCount int
	var initialTitlesCount int
	var initialAuthorsCount int
	var secondClippingsCount int
	var secondTitlesCount int
	var secondAuthorsCount int

	db.DB.Model(&types.Clipping{}).Count(&initialClippingsCount)
	db.DB.Model(&types.Title{}).Count(&initialTitlesCount)
	db.DB.Model(&types.Title{}).Count(&initialAuthorsCount)

	f, err := os.Open(testDataFile)
	require.Nil(err)
	defer f.Close()
	data, err := parser.Parse(f)
	require.Nil(err)
	result, err := db.ProcessParseData(&data)
	require.Nil(err)

	// result should reflect no updates
	assert.Empty(result.Clippings)
	assert.Empty(result.Authors)
	assert.Empty(result.Titles)

	db.DB.Model(&types.Clipping{}).Count(&secondClippingsCount)
	db.DB.Model(&types.Title{}).Count(&secondTitlesCount)
	db.DB.Model(&types.Title{}).Count(&secondAuthorsCount)

	assert.Equal(initialClippingsCount, secondClippingsCount)
	assert.Equal(initialTitlesCount, secondTitlesCount)
	assert.Equal(initialAuthorsCount, secondAuthorsCount)

	err = db.Destroy()
}

// func TestGetTitlesDeep(t *testing.T) {
// 	db, err := createTestDB()
// 	require.Nil(t, err)

// 	var titles []types.Title
// 	db.GetTitlesDeep(&titles)

// 	if assert.NotEqual(t, len(titles), 0, "There should be titles") {
// 		assert.NotNil(t, titles[0].Authors, "Authors should not be nil")
// 		assert.NotNil(t, titles[0].Clippings, "Clippings should not be nil")
// 	}

// 	db.Destroy()
// }

// func TestGetAuthorsDeep(t *testing.T) {
// 	db, err := createTestDB()
// 	require.Nil(t, err)

// 	var authors []types.Author
// 	db.GetAuthorsDeep(&authors)

// 	if assert.NotEqual(t, len(authors), 0, "There should be authors") {
// 		if assert.NotNil(t, authors[0].Titles, "Titles should not be nil") {
// 			assert.NotNil(t, authors[0].Titles[0].Clippings, "Clippings should not be nil")
// 		}
// 	}

// 	db.Destroy()
// }

// func TestGetClippingsDeep(t *testing.T) {
// 	db, err := createTestDB()
// 	require.Nil(t, err)

// 	var clippings []types.Clipping
// 	db.GetClippingsDeep(&clippings)

// 	if assert.NotEqual(t, len(clippings), 0, "There should be clippings") {
// 		// t.Logf("%#v\n", clippings[0])
// 		if assert.NotNil(t, clippings[0].Title, "Title should not be nil") {
// 			assert.NotNil(t, clippings[0].Title.Authors, "Authors should not be nil")
// 		}
// 	}

// 	db.Destroy()
// }

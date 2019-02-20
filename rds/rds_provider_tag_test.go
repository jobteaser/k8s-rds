package rds

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTagsFromString(t *testing.T) {
	t.Run("nominal case", func(t *testing.T) {
		input := `{ "key": "value" }`
		actual, err := ParseTagsFromString(input)
		expected := map[string]string{"key": "value"}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})
	t.Run("empty case", func(t *testing.T) {
		input := "{}"
		actual, err := ParseTagsFromString(input)
		expected := map[string]string{}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})
	t.Run("empty string case", func(t *testing.T) {
		input := ""
		actual, err := ParseTagsFromString(input)
		expected := map[string]string{}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})
	t.Run("error case with unexpected structure", func(t *testing.T) {
		input := `"unexpected"`
		_, err := ParseTagsFromString(input)

		assert.Error(t, err)
	})
}

func TestExpandTagString(t *testing.T) {
	namespace := "jobteaser"
	database := "jobteaser-11056-mysql"

	t.Run("no expansion", func(t *testing.T) {
		input := "string with no expansion char"
		actual, err := ExpandTagString(input, namespace, database)
		expected := "string with no expansion char"

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("no expansion 2", func(t *testing.T) {
		input := "string 2 with no expansion char"
		actual, err := ExpandTagString(input, namespace, database)
		expected := "string 2 with no expansion char"

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("%% character reduction", func(t *testing.T) {
		input := "contains %% character"
		actual, err := ExpandTagString(input, namespace, database)
		expected := "contains % character"

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("%% character reduction (multiple)", func(t *testing.T) {
		input := "contains %%%% character"
		actual, err := ExpandTagString(input, namespace, database)
		expected := "contains %% character"

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("%n and %d character expansion", func(t *testing.T) {
		input := "namespace is %n and database is %d"
		actual, err := ExpandTagString(input, namespace, database)
		expected := "namespace is jobteaser and database is jobteaser-11056-mysql"

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("%n and %d character expansion with annoying %%", func(t *testing.T) {
		input := "%%n is not a namespace, namespace is %n and database is %d"
		actual, err := ExpandTagString(input, namespace, database)
		expected := "%n is not a namespace, namespace is jobteaser and database is jobteaser-11056-mysql"

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("unknown expansion", func(t *testing.T) {
		input := "expansion %f will generate an error"
		_, err := ExpandTagString(input, namespace, database)

		assert.Error(t, err)
	})

	t.Run("trailing % character", func(t *testing.T) {
		input := "expansion will generate an error if last character is %"
		_, err := ExpandTagString(input, namespace, database)

		assert.Error(t, err)
	})

}

func TestGetAwsTags(t *testing.T) {
	namespace := "jobteaser"
	database := "jobteaser-11056-mysql"

	t.Run("no tags", func(t *testing.T) {
		specTags := map[string]string{}
		defaultTags := map[string]string{}
		actual, err := GetAwsTags(specTags, defaultTags, namespace, database)
		expected := []rds.Tag{}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("a few tags from spec only", func(t *testing.T) {
		specTags := map[string]string{
			"com.example.component": "database",
		}
		defaultTags := map[string]string{}
		actual, err := GetAwsTags(specTags, defaultTags, namespace, database)
		expected := []rds.Tag{
			{Key: aws.String("com.example.component"), Value: aws.String("database")},
		}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("a few distinct tags from both specs and default", func(t *testing.T) {
		specTags := map[string]string{
			"com.example.component": "database",
		}
		defaultTags := map[string]string{
			"com.example.environment": "dev",
		}
		actual, err := GetAwsTags(specTags, defaultTags, namespace, database)
		expected := []rds.Tag{
			{Key: aws.String("com.example.component"), Value: aws.String("database")},
			{Key: aws.String("com.example.environment"), Value: aws.String("dev")},
		}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("default tags have higher prio on key present in both default and spec tags", func(t *testing.T) {
		specTags := map[string]string{
			"com.example.environment": "prod",
		}
		defaultTags := map[string]string{
			"com.example.environment": "dev",
		}
		actual, err := GetAwsTags(specTags, defaultTags, namespace, database)
		expected := []rds.Tag{
			{Key: aws.String("com.example.environment"), Value: aws.String("dev")},
		}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})
}

func TestGetTags(t *testing.T) {
	namespace := "jobteaser"
	database := "jobteaser-11056-mysql"

	t.Run("no tags", func(t *testing.T) {
		specTags := map[string]string{}
		defaultTags := map[string]string{}
		actual, err := GetTags(specTags, defaultTags, namespace, database)
		expected := map[string]string{}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("a few tags from spec only", func(t *testing.T) {
		specTags := map[string]string{
			"com.example.component": "database",
		}
		defaultTags := map[string]string{}
		actual, err := GetTags(specTags, defaultTags, namespace, database)
		expected := map[string]string{
			"com.example.component": "database",
		}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("a few distinct tags from both specs and default", func(t *testing.T) {
		specTags := map[string]string{
			"com.example.component": "database",
		}
		defaultTags := map[string]string{
			"com.example.environment": "dev",
		}
		actual, err := GetTags(specTags, defaultTags, namespace, database)
		expected := map[string]string{
			"com.example.component":   "database",
			"com.example.environment": "dev",
		}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("default tags have higher prio on key present in both default and spec tags", func(t *testing.T) {
		specTags := map[string]string{
			"com.example.environment": "prod",
		}
		defaultTags := map[string]string{
			"com.example.environment": "dev",
		}
		actual, err := GetTags(specTags, defaultTags, namespace, database)
		expected := map[string]string{
			"com.example.environment": "dev",
		}

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})
}

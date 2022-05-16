package faker

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	mu                   = &sync.Mutex{}
	generateUniqueValues = false
	uniqueValues         = map[string][]interface{}{}
)

var (
	ErrUniqueFailure = "Failed to generate a unique value for field \"%s\""
	ErrNoUniqueTag   = "Failed to generate fake data without unique values for tag \"%s\""
)

const (
	maxRetry               = 10000
	TitleMaleTag           = "title_male_tag"
	TitleFemaleTag         = "title_female_tag"
	FirstNameMaleTag       = "first_name_male_tag"
	FirstNameFemaleTag     = "first_name_female_tag"
	LastNameMaleTag        = "last_name_male_tag"
	LastNameFemaleTag      = "last_name_female_tag"
	GenderTag              = "gender_tag"
	FirstNameTag           = "first_name_tag"
	LastNameTag            = "last_name_tag"
	RandomFirstNameTag     = "random_first_name_tag"
	RandomLastNameTag      = "random_last_name_tag"
	TitleMaleTagID         = "title_male_id_tag"
	TitleFemaleTagID       = "title_female_id_tag"
	FirstNameMaleTagID     = "first_name_male_id_tag"
	FirstNameFemaleTagID   = "first_name_female_id_tag"
	LastNameMaleTagID      = "last_name_male_id_tag"
	LastNameFemaleTagID    = "last_name_female_id_tag"
	GenderTagID            = "gender_id_tag"
	TitleMaleTagJP         = "title_male_jp_tag"
	TitleMaleTagJPEN       = "title_male_jp_en_tag"
	TitleFemaleTagJP       = "title_female_jp_tag"
	TitleFemaleTagJPEN     = "title_female_jp_en_tag"
	FirstNameMaleTagJP     = "first_name_male_jp_tag"
	FirstNameMaleTagJPEN   = "first_name_male_jp_en_tag"
	FirstNameFemaleTagJP   = "first_name_female_jp_tag"
	FirstNameFemaleTagJPEN = "first_name_female_jp_en_tag"
	LastNameTagJP          = "last_name_jp_tag"
	LastNameTagJPEN        = "last_name_jp_en"
	CountryTag             = "country_tag"
)

type TaggedFunction func(reflect.Value) interface{}

func isUnique(tag string, value interface{}) bool {
	if _, ok := uniqueValues[tag]; !ok {
		uniqueValues[tag] = []interface{}{value}
		return true
	}

	for _, v := range uniqueValues[tag] {
		if v == value {
			return false
		}
	}

	uniqueValues[tag] = append(uniqueValues[tag], value)
	return true
}

func uniqueTag(tag string, fn TaggedFunction) (interface{}, error) {
	if _, ok := uniqueValues[tag]; !ok {
		return fn(reflect.Value{}), nil
	}

	for i := 0; i < maxRetry; i++ {
		res := fn(reflect.Value{})
		if isUnique(tag, res) {
			return res, nil
		}
	}

	return nil, fmt.Errorf(ErrUniqueFailure, tag)
}

func fakeData(tag string, fn TaggedFunction) interface{} {
	if generateUniqueValues {
		res, err := uniqueTag(tag, fn)
		if err != nil {
			panic(fmt.Errorf(ErrNoUniqueTag, tag))
		}

		return res
	}

	return fn(reflect.Value{})
}

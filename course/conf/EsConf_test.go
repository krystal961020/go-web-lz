package conf

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

const mapping = `
{
    "mappings": {
        "properties": {
            "id": {
                "type": "long"
            },
            "title": {
                "type": "text"
            },
            "genres": {
                "type": "keyword"
            }
        }
    }
}`
const index = "krystal"

var (
	subject Subject
)

type Subject struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}

func TestInit(t *testing.T) {
	ctx := context.Background()
	exists, err := EsClient.IndexExists(index).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err = EsClient.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
	subject = Subject{
		ID:     1,
		Title:  "肖恩克的救赎",
		Genres: []string{"犯罪", "剧情"},
	}
	result, err := EsClient.Get().Index(index).
		Id(strconv.Itoa(subject.ID)).Do(ctx)
	if err != nil {
		panic(err)
	}
	if result.Found {
		fmt.Printf("Got document %v (version=%d, index=%s, type=%s)\n",
			result.Id, result.Version, result.Index, result.Type)
	}
}

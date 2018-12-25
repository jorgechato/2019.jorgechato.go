package controller

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/service"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func GetMisc(p graphql.ResolveParams) (interface{}, error) {
	misc := service.GetMiscByID(p.Args["id"].(string))

	return misc, nil
}

func GetMiscs(p graphql.ResolveParams) (interface{}, error) {
	misc := service.GetMiscList(
		p.Args["first"].(int),
		p.Args["offset"].(int),
	)

	return misc, nil
}

func UpdateMisc(p graphql.ResolveParams) (interface{}, error) {
	var misc Misc

	misc.ID = p.Args["id"].(string)
	if misc.ID == "" {
		return nil, nil
	}

	misc.Title = p.Args["title"].(string)
	misc.Content = p.Args["content"].(string)
	misc.Url = p.Args["url"].(string)
	misc.Public = p.Args["public"].(bool)

	switch arg := p.Args["published_at"].(type) {
	case string:
		time, _ := time.Parse(time.RFC3339, arg)
		misc.PublishedAt = time
	case time.Time:
		misc.PublishedAt = arg
	}

	if obj, ok := p.Args["thumbnail"].(string); ok == true {
		misc.Thumbmail = obj
	}

	service.UpdateMisc(&misc)

	return misc, nil
}

func CreateMisc(p graphql.ResolveParams) (interface{}, error) {
	var misc Misc

	misc.Title = p.Args["title"].(string)
	misc.Content = p.Args["content"].(string)
	misc.Url = p.Args["url"].(string)
	misc.Public = p.Args["public"].(bool)

	switch arg := p.Args["published_at"].(type) {
	case string:
		time, _ := time.Parse(time.RFC3339, arg)
		misc.PublishedAt = time
	case time.Time:
		misc.PublishedAt = arg
	}

	if obj, ok := p.Args["thumbnail"].(string); ok == true {
		misc.Thumbmail = obj
	}

	service.CreateMisc(&misc)

	return misc, nil
}

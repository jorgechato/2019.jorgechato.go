package misc

import (
	"time"

	"github.com/graphql-go/graphql"

	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

// CreateMisc create a misc
func CreateMisc() (field *graphql.Field) {
	field = &graphql.Field{
		Type: MiscType,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
			"content": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
			"url": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
			"thumbnail": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"public": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
			"published_at": &graphql.ArgumentConfig{
				Type:         graphql.DateTime,
				DefaultValue: time.Now(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

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

			createMisc(&misc)

			return misc, nil
		},
	}

	return
}

// UpdateMisc create a misc
func UpdateMisc() (field *graphql.Field) {
	field = &graphql.Field{
		Type: MiscType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"title": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
			"content": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
			"url": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
			"thumbnail": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"public": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
			"published_at": &graphql.ArgumentConfig{
				Type:         graphql.DateTime,
				DefaultValue: time.Now(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			var misc Misc

			misc.ID = p.Args["id"].(string)
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

			updateMisc(&misc)

			return misc, nil
		},
	}

	return
}

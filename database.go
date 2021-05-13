package notion

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const databaseEndpoint = "databases"

// Database describes the property schema of a database in Notion.
type Database struct {
	Object         string              `json:"object"`
	ID             string              `json:"id"`
	CreatedTime    string              `json:"created_time"`
	LastEditedTime string              `json:"last_edited_time"`
	Title          []RichText          `json:"title"`
	Properties     map[string]Property `json:"properties"`
}

// RichText contains data for displaying formatted text, mentions, and equations.
type RichText struct {
	PlainText   string     `json:"plain_text"`
	Href        string     `json:"href,omitempty"`
	Annotations Annotation `json:"annotations"`
	Type        string     `json:"type"`
	Text        Text       `json:"text,omitempty"`
	Mention     Mention    `json:"mention,omitempty"`
}

// Annotation contains style information which applies to the whole rich text object.
type Annotation struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

// Text contains the following information within the text property.
type Text struct {
	Content string
	Link    *Link `json:"link,omitempty"`
}

// Link contains a type key whose value is always "url" and a url key whose value is a web address.
type Link struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// Mention represents an inline mention of a user, page, database, or date.
// In the app these are created by typing @ followed by the name of a user, page, database, or a date.
type Mention struct {
	Type string `json:"type"`
	User User   `json:"user,omitempty"`
}

// Property has metadata that controls how a database property behaves.
type Property struct {
	ID          string       `json:"id,omitempty"`
	Type        string       `json:"type,omitempty"`
	Number      *Number      `json:"number,omitempty"`
	Select      *Select      `json:"select,omitempty"`
	MultiSelect *MultiSelect `json:"multi_select,omitempty"`
	Formula     *Formula     `json:"formula,omitempty"`
	Relation    *Relation    `json:"relation,omitempty"`
	Rollup      *Rollup      `json:"rollup,omitempty"`
}

// Number database property contains the following configuration within the number property.
type Number struct {
	Format string `json:"format"`
}

// Select database property contains the following configuration within the select property.
type Select struct {
	Options []SelectOption `json:"options"`
}

// SelectOption encapsulates the properties for the type Select.
type SelectOption struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Color string `json:"color"`
}

// Multi-select database property contains the following configuration within the multi_select property.
type MultiSelect struct {
	Options []MultiSelectOption `json:"options"`
}

// SelectOption encapsulates the properties for the type MultiSelect.
type MultiSelectOption struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Color string `json:"color"`
}

// Formula database property contains the following configuration within the formula property.
type Formula struct {
	Expression string `json:"expression"`
}

// Relation database property contains the following configuration within the relation property.
type Relation struct {
	DatabaseID         string `json:"database_id"`
	SyncedPropertyName string `json:"synced_property_name,omitempty"`
	SyncedPropertyID   string `json:"synced_property_id,omitempty"`
}

// Rollup database property contains the following configuration within the rollup property.
type Rollup struct {
	RelationPropertyName string `json:"relation_property_name"`
	RelationPropertyID   string `json:"relation_property_id"`
	RollupPropertyName   string `json:"rollup_property_name"`
	RollupPropertyID     string `json:"rollup_property_id"`
	Function             string `json:"function"`
}

// GetDatabase retrieves a Database using the ID specified.
func (c Client) GetDatabase(id string) (Database, error) {
	url := c.Config.BaseURL + "/" + c.Config.APIVersion + "/" + databaseEndpoint + "/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Database{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", c.Config.HeaderVersion)
	resp, err := c.Client.Do(req)
	if err != nil {
		return Database{}, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return Database{}, errors.New("database not found")
	}
	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusTooManyRequests {
		return Database{}, errors.New("request exceeded the request limits")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Database{}, err
	}
	var database Database
	err = json.Unmarshal(body, &database)
	if err != nil {
		return Database{}, err
	}
	return database, nil
}

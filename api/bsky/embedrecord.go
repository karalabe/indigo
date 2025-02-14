package bsky

import (
	"encoding/json"
	"fmt"

	comatprototypes "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/lex/util"
)

// schema: app.bsky.embed.record

func init() {
	util.RegisterType("app.bsky.embed.record#main", &EmbedRecord{})
}

// RECORDTYPE: EmbedRecord
type EmbedRecord struct {
	LexiconTypeID string                         `json:"$type,const=app.bsky.embed.record" cborgen:"$type,const=app.bsky.embed.record"`
	Record        *comatprototypes.RepoStrongRef `json:"record" cborgen:"record"`
}

// RECORDTYPE: EmbedRecord_View
type EmbedRecord_View struct {
	LexiconTypeID string                   `json:"$type,const=app.bsky.embed.record" cborgen:"$type,const=app.bsky.embed.record"`
	Record        *EmbedRecord_View_Record `json:"record" cborgen:"record"`
}

// RECORDTYPE: EmbedRecord_ViewNotFound
type EmbedRecord_ViewNotFound struct {
	LexiconTypeID string `json:"$type,const=app.bsky.embed.record" cborgen:"$type,const=app.bsky.embed.record"`
	Uri           string `json:"uri" cborgen:"uri"`
}

// RECORDTYPE: EmbedRecord_ViewRecord
type EmbedRecord_ViewRecord struct {
	LexiconTypeID string                                `json:"$type,const=app.bsky.embed.record" cborgen:"$type,const=app.bsky.embed.record"`
	Author        *ActorDefs_ProfileViewBasic           `json:"author" cborgen:"author"`
	Cid           string                                `json:"cid" cborgen:"cid"`
	Embeds        []*EmbedRecord_ViewRecord_Embeds_Elem `json:"embeds,omitempty" cborgen:"embeds,omitempty"`
	IndexedAt     string                                `json:"indexedAt" cborgen:"indexedAt"`
	Uri           string                                `json:"uri" cborgen:"uri"`
	Value         *util.LexiconTypeDecoder              `json:"value" cborgen:"value"`
}

type EmbedRecord_ViewRecord_Embeds_Elem struct {
	EmbedImages_View          *EmbedImages_View
	EmbedExternal_View        *EmbedExternal_View
	EmbedRecord_View          *EmbedRecord_View
	EmbedRecordWithMedia_View *EmbedRecordWithMedia_View
}

func (t *EmbedRecord_ViewRecord_Embeds_Elem) MarshalJSON() ([]byte, error) {
	if t.EmbedImages_View != nil {
		t.EmbedImages_View.LexiconTypeID = "app.bsky.embed.images#view"
		return json.Marshal(t.EmbedImages_View)
	}
	if t.EmbedExternal_View != nil {
		t.EmbedExternal_View.LexiconTypeID = "app.bsky.embed.external#view"
		return json.Marshal(t.EmbedExternal_View)
	}
	if t.EmbedRecord_View != nil {
		t.EmbedRecord_View.LexiconTypeID = "app.bsky.embed.record#view"
		return json.Marshal(t.EmbedRecord_View)
	}
	if t.EmbedRecordWithMedia_View != nil {
		t.EmbedRecordWithMedia_View.LexiconTypeID = "app.bsky.embed.recordWithMedia#view"
		return json.Marshal(t.EmbedRecordWithMedia_View)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *EmbedRecord_ViewRecord_Embeds_Elem) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.embed.images#view":
		t.EmbedImages_View = new(EmbedImages_View)
		return json.Unmarshal(b, t.EmbedImages_View)
	case "app.bsky.embed.external#view":
		t.EmbedExternal_View = new(EmbedExternal_View)
		return json.Unmarshal(b, t.EmbedExternal_View)
	case "app.bsky.embed.record#view":
		t.EmbedRecord_View = new(EmbedRecord_View)
		return json.Unmarshal(b, t.EmbedRecord_View)
	case "app.bsky.embed.recordWithMedia#view":
		t.EmbedRecordWithMedia_View = new(EmbedRecordWithMedia_View)
		return json.Unmarshal(b, t.EmbedRecordWithMedia_View)

	default:
		return nil
	}
}

type EmbedRecord_View_Record struct {
	EmbedRecord_ViewRecord   *EmbedRecord_ViewRecord
	EmbedRecord_ViewNotFound *EmbedRecord_ViewNotFound
}

func (t *EmbedRecord_View_Record) MarshalJSON() ([]byte, error) {
	if t.EmbedRecord_ViewRecord != nil {
		t.EmbedRecord_ViewRecord.LexiconTypeID = "app.bsky.embed.record#viewRecord"
		return json.Marshal(t.EmbedRecord_ViewRecord)
	}
	if t.EmbedRecord_ViewNotFound != nil {
		t.EmbedRecord_ViewNotFound.LexiconTypeID = "app.bsky.embed.record#viewNotFound"
		return json.Marshal(t.EmbedRecord_ViewNotFound)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *EmbedRecord_View_Record) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.embed.record#viewRecord":
		t.EmbedRecord_ViewRecord = new(EmbedRecord_ViewRecord)
		return json.Unmarshal(b, t.EmbedRecord_ViewRecord)
	case "app.bsky.embed.record#viewNotFound":
		t.EmbedRecord_ViewNotFound = new(EmbedRecord_ViewNotFound)
		return json.Unmarshal(b, t.EmbedRecord_ViewNotFound)

	default:
		return nil
	}
}

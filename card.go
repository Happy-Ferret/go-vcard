// Package vcard implements the vCard format, defined in RFC 6350.
package vcard

import (
	"strings"
)

// Card property parameters.
const (
	ParamLanguage = "LANGUAGE"
	ParamValue = "VALUE"
	ParamPreferred = "PREF"
	ParamAltID = "ALTID"
	ParamPID = "PID"
	ParamType = "TYPE"
	ParamMediaType = "MEDIATYPE"
	ParamCalendarScale = "CALSCALE"
	ParamSortAs = "SORT-AS"
	ParamGeo = "GEO"
	ParamTimezone = "TZ"
)

// Card properties.
const (
	// General Properties
	FieldSource = "SOURCE"
	FieldKind = "KIND"
	FieldXML = "XML"

	// Identification Properties
	FieldFormattedName = "FN"
	FieldName = "N"
	FieldNickname = "NICKNAME"
	FieldPhoto = "PHOTO"
	FieldBirthday = "BDAY"
	FieldAnniversary = "ANIVERSARY"
	FieldGender = "GENDER"

	// Delivery Addressing Properties
	FieldAddress = "ADR"

	// Communications Properties
	FieldTelephone = "TEL"
	FieldEmail = "EMAIL"
	FieldIMPP = "IMPP" // Instant Messaging and Presence Protocol
	FieldLanguage = "LANG"

	// Geographical Properties
	FieldTimezone = "TZ"
	FieldGeo = "GEO"

	// Organizational Properties
	FieldTitle = "TITLE"
	FieldRole = "ROLE"
	FieldLogo = "LOGO"
	FieldOrganization = "ORG"
	FieldMember = "MEMBER"
	FieldRelated = "RELATED"

	// Explanatory Properties
	FieldCategories = "CATEGORIES"
	FieldNote = "NOTE"
	FieldProductID = "PRODID"
	FieldRevision = "REV"
	FieldSound = "SOUND"
	FieldUID = "UID"
	FieldClientPIDMap = "CLIENTPIDMAP"
	FieldURL = "URL"
	FieldVersion = "VERSION"

	// Security Properties
	FieldKey = "KEY"

	// Calendar Properties
	FieldFBURL = "FBURL"
	FieldCalendarAddressURI = "CALADRURI"
	FieldCalendarURI = "CALURI"
)

func maybeGet(l []string, i int) string {
	if i < len(l) {
		return l[i]
	}
	return ""
}

func fieldValues(fields []*Field) []string {
	if fields == nil {
		return nil
	}

	values := make([]string, len(fields))
	for i, f := range fields {
		values[i] = f.Value
	}
	return values
}

// A Card is an address book entry.
type Card map[string][]*Field

// Get returns the first field of the card. If there is no such field, it
// returns nil.
func (c Card) Get(k string) *Field {
	fields := c[k]
	if len(fields) == 0 {
		return nil
	}
	return fields[0]
}

// Preferred returns the preferred field of the card.
func (c Card) Preferred(k string) *Field {
	fields := c[k]
	if len(fields) == 0 {
		return nil
	}

	for _, f := range fields {
		if f.Params[ParamPreferred] == "1" {
			return f
		}
	}
	return fields[0]
}

// Value returns the first field value of the card. If there is no such field,
// it returns an empty string.
func (c Card) Value(k string) string {
	f := c.Get(k)
	if f == nil {
		return ""
	}
	return f.Value
}

// PreferredValue returns the preferred field value of the card.
func (c Card) PreferredValue(k string) string {
	f := c.Preferred(k)
	if f == nil {
		return ""
	}
	return f.Value
}

// Sources return a list of sources of directory information contained in the
// content type.
func (c Card) Sources() []string {
	return fieldValues(c[FieldSource])
}

// Kind returns the kind of the object represented by this card. If it isn't
// specified, it returns an empty string. The default is KindIndividual.
func (c Card) Kind() string {
	return strings.ToLower(c.Value(FieldKind))
}

// FormattedNames returns formatted names of the card. The length of the result
// is always greater or equal to 1.
func (c Card) FormattedNames() []*Field {
	fns := c[FieldFormattedName]
	if len(fns) == 0 {
		return []*Field{{Value: ""}}
	}
	return fns
}

// Names returns names of the card.
func (c Card) Names() []*Name {
	ns := c[FieldName]
	if ns == nil {
		return nil
	}

	names := make([]*Name, len(ns))
	for i, n := range ns {
		names[i] = newName(n)
	}
	return names
}

// Name returns the preferred name of the card. If it isn't specified, it
// returns nil.
func (c Card) Name() *Name {
	n := c.Preferred(FieldName)
	if n == nil {
		return nil
	}
	return newName(n)
}

// A field contains a value and some parameters.
type Field struct {
	Value string
	Params map[string]string
	Group string
}

const (
	KindIndividual = "individual"
	KindGroup = "group"
	KindOrg = "org"
	KindLocation = "location"
)

const (
	TypeHome = "home"
	TypeWork = "work"
)

type Name struct {
	FamilyName string
	GivenName string
	AdditionalName string
	HonorificPrefix string
	HonorificSuffix string

	Params map[string]string
	Group string
}

func newName(field *Field) *Name {
	components := strings.Split(field.Value, ";")
	return &Name{
		maybeGet(components, 0),
		maybeGet(components, 1),
		maybeGet(components, 2),
		maybeGet(components, 3),
		maybeGet(components, 4),
		field.Params,
		field.Group,
	}
}

type Gender string

const (
	GenderFemale Gender = "F"
	GenderMale = "M"
	GenderOther = "O"
	GenderNone = "N"
	GenderUnknown = "U"
)

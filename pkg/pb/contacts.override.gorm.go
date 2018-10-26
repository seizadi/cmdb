package pb

import (
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"

	"github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
	"github.com/infobloxopen/atlas-app-toolkit/query"
)

// AfterToORM will add the primary e-mail to the list of e-mails if it isn't
// present already
func (m *Contact) AfterToORM(ctx context.Context, c *ContactORM) error {
	for _, g := range m.Groups {
		groupId, err := resource.DecodeInt64(&Group{}, g)
		if err != nil {
			return err
		}
		c.Groups = append(c.Groups, &GroupORM{Id: groupId})
	}

	if m.PrimaryEmail == "" {
		return nil
	}

	var primary *EmailORM

	emails := []*EmailORM{}
	for _, e := range c.Emails {
		if e.Address != m.PrimaryEmail {
			e.IsPrimary = new(bool)
			*e.IsPrimary = false
			emails = append(emails, e)
		} else {
			e.IsPrimary = new(bool)
			*e.IsPrimary = true
			primary = e
		}
	}

	if primary == nil {
		if e, err := (&Email{Address: m.PrimaryEmail}).ToORM(ctx); err != nil {
			return err
		} else {
			primary = &e
		}

		primary.IsPrimary = new(bool)
		*primary.IsPrimary = true
	}

	emails = append(emails, primary)
	c.Emails = emails

	return nil
}

// AfterToPB copies the primary e-mail address from the DB to the special PB field
func (m *ContactORM) AfterToPB(ctx context.Context, c *Contact) error {
	// find the primary e-mail in list of e-mails from DB
	for _, addr := range m.Emails {
		if addr != nil && addr.IsPrimary != nil && *addr.IsPrimary {
			c.PrimaryEmail = addr.Address
			break
		}
	}

	for _, g := range m.Groups {
		groupId, err := resource.Encode(&Group{}, g.Id)
		if err != nil {
			return err
		}
		c.Groups = append(c.Groups, groupId)
	}
	return nil
}

func (m *ContactORM) BeforeStrictUpdateSave(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	if err := db.Model(m).Association("Groups").Replace(m.Groups).Error; err != nil {
		return nil, err
	}
	return db, nil
}

func (m *ContactORM) BeforeReadApplyQuery(ctx context.Context, db *gorm.DB, fs *query.FieldSelection) (*gorm.DB, error) {
	for _, field := range fs.GetFields() {
		if field.GetName() == "primary_email" {
			db = db.Preload("Emails")
		}
	}
	return db, nil
}

func (m *ContactORM) BeforeListApplyQuery(ctx context.Context, db *gorm.DB, f *query.Filtering, s *query.Sorting, p *query.Pagination, fs *query.FieldSelection) (*gorm.DB, error) {
	syntheticFields := make(map[string]struct{})
	if f != nil {
		syntheticFields = IterateFiltering(f, supportSynteticFields())
	}
	for _, cr := range s.GetCriterias() {
		if cr.GetTag() == "primary_email" {
			cr.Tag = "primary_email.address"
			syntheticFields["primary_email"] = struct{}{}
		}
	}
	if _, ok := syntheticFields["primary_email"]; ok {
		db = db.Joins("join emails primary_email on contacts.id = primary_email.contact_id and primary_email.is_primary = true")
	}
	for _, field := range fs.GetFields() {
		if field.GetName() == "primary_email" {
			db = db.Preload("Emails")
		}
	}
	return db, nil
}

type FilteringIteratorCallback func(path []string, f interface{}) string

// IterateFiltering calls callback function for each condtion struct of *Filtering.
func IterateFiltering(f *query.Filtering, callback FilteringIteratorCallback) map[string]struct{} {
	syntheticFields := make(map[string]struct{})

	var getOperator func(interface{}) interface{}

	doCallback := func(path []string, f interface{}) {
		field := callback(path, f)
		if field != "" {
			syntheticFields[field] = struct{}{}
		}
	}

	getOperator = func(f interface{}) interface{} {
		val := f.(*query.LogicalOperator)

		left := val.GetLeft()
		switch leftVal := left.(type) {
		case *query.LogicalOperator_LeftOperator:
			val.SetLeft(getOperator(leftVal.LeftOperator))

		case *query.LogicalOperator_LeftStringCondition:
			doCallback(leftVal.LeftStringCondition.GetFieldPath(), leftVal.LeftStringCondition)

		case *query.LogicalOperator_LeftNumberCondition:
			doCallback(leftVal.LeftNumberCondition.GetFieldPath(), leftVal.LeftNumberCondition)

		case *query.LogicalOperator_LeftNullCondition:
			doCallback(leftVal.LeftNullCondition.GetFieldPath(), leftVal.LeftNullCondition)
		}

		right := val.GetRight()
		switch rightVal := right.(type) {
		case *query.LogicalOperator_RightOperator:
			getOperator(rightVal.RightOperator)

		case *query.LogicalOperator_RightStringCondition:
			doCallback(rightVal.RightStringCondition.GetFieldPath(), rightVal.RightStringCondition)

		case *query.LogicalOperator_RightNumberCondition:
			doCallback(rightVal.RightNumberCondition.GetFieldPath(), rightVal.RightNumberCondition)

		case *query.LogicalOperator_RightNullCondition:
			doCallback(rightVal.RightNullCondition.GetFieldPath(), rightVal.RightNullCondition)
		}
		return val
	}

	root := f.GetRoot()
	switch val := root.(type) {
	case *query.Filtering_Operator:
		getOperator(val.Operator)

	case *query.Filtering_StringCondition:
		doCallback(val.StringCondition.GetFieldPath(), val.StringCondition)

	case *query.Filtering_NumberCondition:
		doCallback(val.NumberCondition.GetFieldPath(), val.NumberCondition)

	case *query.Filtering_NullCondition:
		doCallback(val.NullCondition.GetFieldPath(), val.NullCondition)
	}
	return syntheticFields
}

// callback function for IterateFiltering to support "primary_email" (synthetic field) filtering
func supportSynteticFields() FilteringIteratorCallback {
	return func(path []string, f interface{}) string {
		switch strings.Join(path, ".") {
		case "primary_email":
			sc, ok := f.(*query.StringCondition)
			if ok {
				sc.FieldPath = []string{"primary_email", "address"}
				return "primary_email"
			}
		}
		return ""
	}
}

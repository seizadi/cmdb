package options

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/infobloxopen/atlas-app-toolkit/query"
)

type FilteringOptions struct {
	AllowMissingFields bool
	Options            map[string]FilteringOption
}

type FilteringOption struct {
	ValueType QueryValidate_ValueType
	Deny      []QueryValidate_FilterOperator
}

func getFieldInfo(path []string, messageInfo map[string]FilteringOption) (FilteringOption, error) {
	fieldTag := strings.Join(path, ".")
	if fieldInfo, ok := messageInfo[fieldTag]; ok {
		return fieldInfo, nil
	}

	if len(path) >= 1 {
		fdPrefixTag := strings.Join(path[:1], ".") + ".*"
		if fieldInfo, ok := messageInfo[fdPrefixTag]; ok {
			return fieldInfo, nil
		}
	}

	return FilteringOption{}, fmt.Errorf("Unknown field: %s", fieldTag)
}

func ValidateFiltering(f *query.Filtering, messageInfo map[string]FilteringOption) error {
	var getOperator func(interface{}) error
	var fieldInfo FilteringOption

	validate := func(path []string, f interface{}) error {
		var err error
		fieldTag := strings.Join(path, ".")
		fieldInfo, err = getFieldInfo(path, messageInfo)
		if err != nil {
			return err
		}

		if fieldInfo.ValueType == QueryValidate_DEFAULT {
			return fmt.Errorf("Filtering is not supported for field %s", fieldTag)
		}

		tp := ""

		switch x := f.(type) {
		case *query.StringCondition:
			if fieldInfo.ValueType != QueryValidate_STRING && fieldInfo.ValueType != QueryValidate_BOOL {
				return fmt.Errorf("Got invalid literal type for %s, expect %s", fieldTag, fieldInfo.ValueType)
			}

			if fieldInfo.ValueType == QueryValidate_BOOL {

				if x.Type != query.StringCondition_EQ {
					return fmt.Errorf("Operation %s is not allowed for %q", query.StringCondition_Type_name[int32(x.Type)], fieldTag)
				}

				if _, err := strconv.ParseBool(x.Value); err != nil {
					return fmt.Errorf("Got invalid literal for field %q of type %s, expect 'true' or 'false'", fieldTag, fieldInfo.ValueType)
				}
			}

			sc := &query.Filtering_StringCondition{x}
			tp = query.StringCondition_Type_name[int32(sc.StringCondition.Type)]
		case *query.NumberCondition:
			if fieldInfo.ValueType != QueryValidate_NUMBER {
				return fmt.Errorf("Got invalid literal type for %s, expect %s", fieldTag, fieldInfo.ValueType)
			}
			nc := &query.Filtering_NumberCondition{x}
			tp = query.NumberCondition_Type_name[int32(nc.NumberCondition.Type)]
		case *query.StringArrayCondition:
			if fieldInfo.ValueType != QueryValidate_STRING && fieldInfo.ValueType != QueryValidate_BOOL {
				return fmt.Errorf("Got invalid literal type for %s, expect %s", fieldTag, fieldInfo.ValueType)
			}

			if fieldInfo.ValueType == QueryValidate_BOOL {
				for i, xv := range x.Values {
					if _, err := strconv.ParseBool(xv); err != nil {
						return fmt.Errorf("Got invalid literal for field %q of type %s at position %d, expect 'true' or 'false'", fieldTag, fieldInfo.ValueType, i)
					}
				}
			}

			nc := &query.Filtering_StringArrayCondition{x}
			tp = query.StringArrayCondition_Type_name[int32(nc.StringArrayCondition.Type)]
		case *query.NumberArrayCondition:
			if fieldInfo.ValueType != QueryValidate_NUMBER {
				return fmt.Errorf("Got invalid literal type for %s, expect %s", fieldTag, fieldInfo.ValueType)
			}
			nc := &query.Filtering_NumberArrayCondition{x}
			tp = query.NumberArrayCondition_Type_name[int32(nc.NumberArrayCondition.Type)]
		default:

			return nil
		}
		for _, val := range fieldInfo.Deny {
			if val == QueryValidate_ALL {
				return fmt.Errorf("Operation %s is not allowed for '%s'", tp, fieldTag)
			}
			if val.String() == tp {
				return fmt.Errorf("Operation %s is not allowed for '%s'", tp, fieldTag)
			}
		}
		return nil
	}

	var vres error

	getOperator = func(f interface{}) error {
		val := f.(*query.LogicalOperator)
		var vres error
		left := val.GetLeft()
		switch leftVal := left.(type) {
		case *query.LogicalOperator_LeftOperator:
			vres = getOperator(leftVal.LeftOperator)

		case *query.LogicalOperator_LeftStringCondition:
			vres = validate(leftVal.LeftStringCondition.GetFieldPath(), leftVal.LeftStringCondition)

		case *query.LogicalOperator_LeftNumberCondition:
			vres = validate(leftVal.LeftNumberCondition.GetFieldPath(), leftVal.LeftNumberCondition)

		case *query.LogicalOperator_LeftNullCondition:
			vres = validate(leftVal.LeftNullCondition.GetFieldPath(), leftVal.LeftNullCondition)

		case *query.LogicalOperator_LeftStringArrayCondition:
			vres = validate(leftVal.LeftStringArrayCondition.GetFieldPath(), leftVal.LeftStringArrayCondition)

		case *query.LogicalOperator_LeftNumberArrayCondition:
			vres = validate(leftVal.LeftNumberArrayCondition.GetFieldPath(), leftVal.LeftNumberArrayCondition)
		}

		if vres != nil {
			return vres
		}

		right := val.GetRight()
		switch rightVal := right.(type) {
		case *query.LogicalOperator_RightOperator:
			getOperator(rightVal.RightOperator)

		case *query.LogicalOperator_RightStringCondition:
			vres = validate(rightVal.RightStringCondition.GetFieldPath(), rightVal.RightStringCondition)

		case *query.LogicalOperator_RightNumberCondition:
			vres = validate(rightVal.RightNumberCondition.GetFieldPath(), rightVal.RightNumberCondition)

		case *query.LogicalOperator_RightNullCondition:
			vres = validate(rightVal.RightNullCondition.GetFieldPath(), rightVal.RightNullCondition)

		case *query.LogicalOperator_RightStringArrayCondition:
			vres = validate(rightVal.RightStringArrayCondition.GetFieldPath(), rightVal.RightStringArrayCondition)

		case *query.LogicalOperator_RightNumberArrayCondition:
			vres = validate(rightVal.RightNumberArrayCondition.GetFieldPath(), rightVal.RightNumberArrayCondition)
		}

		return vres
	}

	if f != nil {
		root := f.GetRoot()
		switch val := root.(type) {
		case *query.Filtering_Operator:
			vres = getOperator(val.Operator)

		case *query.Filtering_StringCondition:
			vres = validate(val.StringCondition.GetFieldPath(), val.StringCondition)

		case *query.Filtering_NumberCondition:
			vres = validate(val.NumberCondition.GetFieldPath(), val.NumberCondition)

		case *query.Filtering_NullCondition:
			vres = validate(val.NullCondition.GetFieldPath(), val.NullCondition)

		case *query.Filtering_StringArrayCondition:
			vres = validate(val.StringArrayCondition.GetFieldPath(), val.StringArrayCondition)

		case *query.Filtering_NumberArrayCondition:
			vres = validate(val.NumberArrayCondition.GetFieldPath(), val.NumberArrayCondition)
		}
	}
	return vres
}

func ValidateSorting(p *query.Sorting, fields []string) error {
	if p != nil {
		for _, criteria := range p.GetCriterias() {
			tag := criteria.GetTag()
			var ok bool
			for _, v := range fields {
				if v == tag {
					ok = true
				}
			}
			if !ok {
				return fmt.Errorf("Sorting is not allowed for '%s'", tag)
			}
		}
	}
	return nil
}

func ValidateFieldSelection(fs *query.FieldSelection, allowedFields []string) error {
	var flatten func(fields map[string]*query.Field) []string
	flatten = func(fields map[string]*query.Field) []string {
		var flatFields []string
		for _, v := range fields {
			if v.GetSubs() != nil {
				subFields := flatten(v.GetSubs())
				for _, i := range subFields {
					flatFields = append(flatFields, v.GetName()+"."+i)
				}
			}
			flatFields = append(flatFields, v.GetName())
		}
		return flatFields
	}
	flatFields := flatten(fs.GetFields())
	for _, f := range flatFields {
		var ok bool
		for _, v := range allowedFields {
			if f == v {
				ok = true
			}
		}
		if !ok {
			return fmt.Errorf("Unknown field: '%s'", f)
		}
	}
	return nil
}

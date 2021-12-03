// Code generated by volcago. DO NOT EDIT.
// generated version: 1.12.0
package repository

import (
	"fmt"
	"reflect"

	"cloud.google.com/go/firestore"
	"golang.org/x/xerrors"
)

const (
	alreadyEqual      = "equality operator already used for this field"
	alreadyRange      = "range operator already used for this field"
	alreadyRangeValid = "already has a valid range field: "
	alreadyUsingOnce  = "already using `in` or `array-contains-any`"
	alreadyEqualOrIn  = "this field already contains an `in` clause or equality operator"
	alreadySortOrder  = "sort order is already specified"
	allowArrayOrSlice = "value is not a array or slice"
	cannotCombineAny  = "cannot be combined with `array-contains-any`"
	notLessThan10     = "value is not less than 10"
)

type errs struct {
	TryOperator string
	TryField    string
	TryValue    interface{}
	Message     string
}

func newErrs(op, field, message string, value interface{}) errs {
	return errs{
		TryOperator: op,
		TryField:    field,
		TryValue:    value,
		Message:     message,
	}
}

// QueryBuilder - query builder
type QueryBuilder struct {
	q               firestore.Query
	errs            []errs
	usedIn          bool
	usedAny         bool
	validRangeField string
	equalCounter    map[string]struct{}
	rangeCounter    map[string]struct{}
	orderCounter    map[string]struct{}
}

// NewQueryBuilder - constructor
func NewQueryBuilder(collection *firestore.CollectionRef) *QueryBuilder {
	return &QueryBuilder{
		q:            collection.Query,
		equalCounter: make(map[string]struct{}),
		rangeCounter: make(map[string]struct{}),
		orderCounter: make(map[string]struct{}),
	}
}

// Query - return firestore.Query
func (qb *QueryBuilder) Query() *firestore.Query {
	return &qb.q
}

// Check - condition check
func (qb *QueryBuilder) Check() error {
	if len(qb.errs) <= 0 {
		return nil
	}
	var result string
	for _, err := range qb.errs {
		result += fmt.Sprintf(
			"tryOp: %s, tryField: %s, tryValue: %v, message: %s\n",
			err.TryOperator, err.TryField, err.TryValue, err.Message,
		)
	}
	return xerrors.New(result)
}

// Equal - equality filter ( `==` )
func (qb *QueryBuilder) Equal(path string, value interface{}) *QueryBuilder {
	if _, ok := qb.rangeCounter[path]; ok {
		qb.errs = append(qb.errs, newErrs(OpTypeEqual, path, alreadyRange, value))
		return qb
	}
	qb.equalCounter[path] = struct{}{}
	qb.q = qb.q.Where(path, OpTypeEqual, value)
	return qb
}

// NotEqual - inequality filter ( `!=` )
func (qb *QueryBuilder) NotEqual(path string, value interface{}) *QueryBuilder {
	if _, ok := qb.rangeCounter[path]; ok {
		qb.errs = append(qb.errs, newErrs(OpTypeNotEqual, path, alreadyRange, value))
		return qb
	}
	qb.equalCounter[path] = struct{}{}
	qb.q = qb.q.Where(path, OpTypeNotEqual, value)
	return qb
}

// LessThan - range filter ( `<` )
func (qb *QueryBuilder) LessThan(path string, value interface{}) *QueryBuilder {
	if _, ok := qb.equalCounter[path]; ok {
		qb.errs = append(qb.errs, newErrs(OpTypeLessThan, path, alreadyEqual, value))
		return qb
	}
	if qb.validRangeField != "" {
		qb.errs = append(qb.errs, newErrs(OpTypeLessThan, path, alreadyRangeValid+qb.validRangeField, value))
		return qb
	}
	qb.rangeCounter[path] = struct{}{}
	qb.q = qb.q.Where(path, OpTypeLessThan, value)
	return qb
}

// LessThanOrEqual - range filter ( `<=` )
func (qb *QueryBuilder) LessThanOrEqual(path string, value interface{}) *QueryBuilder {
	if _, ok := qb.equalCounter[path]; ok {
		qb.errs = append(qb.errs, newErrs(OpTypeLessThanOrEqual, path, alreadyEqual, value))
		return qb
	}
	if qb.validRangeField != "" {
		qb.errs = append(qb.errs, newErrs(OpTypeLessThanOrEqual, path, alreadyRangeValid+qb.validRangeField, value))
		return qb
	}
	qb.rangeCounter[path] = struct{}{}
	qb.q = qb.q.Where(path, OpTypeLessThanOrEqual, value)
	return qb
}

// GreaterThan - range filter ( `>` )
func (qb *QueryBuilder) GreaterThan(path string, value interface{}) *QueryBuilder {
	if _, ok := qb.equalCounter[path]; ok {
		qb.errs = append(qb.errs, newErrs(OpTypeGreaterThan, path, alreadyEqual, value))
		return qb
	}
	if qb.validRangeField != "" {
		qb.errs = append(qb.errs, newErrs(OpTypeGreaterThan, path, alreadyRangeValid+qb.validRangeField, value))
		return qb
	}
	qb.rangeCounter[path] = struct{}{}
	qb.q = qb.q.Where(path, OpTypeGreaterThan, value)
	return qb
}

// GreaterThanOrEqual - range filter ( `>=` )
func (qb *QueryBuilder) GreaterThanOrEqual(path string, value interface{}) *QueryBuilder {
	if _, ok := qb.equalCounter[path]; ok {
		qb.errs = append(qb.errs, newErrs(OpTypeGreaterThanOrEqual, path, alreadyEqual, value))
		return qb
	}
	if qb.validRangeField != "" {
		qb.errs = append(qb.errs, newErrs(OpTypeGreaterThanOrEqual, path, alreadyRangeValid+qb.validRangeField, value))
		return qb
	}
	qb.rangeCounter[path] = struct{}{}
	qb.q = qb.q.Where(path, OpTypeGreaterThanOrEqual, value)
	return qb
}

// In - array filter ( `in` )
func (qb *QueryBuilder) In(path string, value interface{}) *QueryBuilder {
	switch rv := reflect.ValueOf(value); rv.Kind() {
	case reflect.Slice, reflect.Array:
		if rv.Len() > 10 {
			qb.errs = append(qb.errs, newErrs(OpTypeIn, path, notLessThan10, value))
			return qb
		}
	default:
		qb.errs = append(qb.errs, newErrs(OpTypeIn, path, allowArrayOrSlice, value))
		return qb
	}
	if qb.usedIn || qb.usedAny {
		qb.errs = append(qb.errs, newErrs(OpTypeIn, path, alreadyUsingOnce, value))
		return qb
	}
	qb.usedIn = true
	qb.q = qb.q.Where(path, OpTypeIn, value)
	return qb
}

// NotIn - array filter ( `not-in` )
func (qb *QueryBuilder) NotIn(path string, value interface{}) *QueryBuilder {
	switch rv := reflect.ValueOf(value); rv.Kind() {
	case reflect.Slice, reflect.Array:
		if rv.Len() > 10 {
			qb.errs = append(qb.errs, newErrs(OpTypeNotIn, path, notLessThan10, value))
			return qb
		}
	default:
		qb.errs = append(qb.errs, newErrs(OpTypeNotIn, path, allowArrayOrSlice, value))
		return qb
	}
	if qb.usedIn || qb.usedAny {
		qb.errs = append(qb.errs, newErrs(OpTypeNotIn, path, alreadyUsingOnce, value))
		return qb
	}
	qb.usedIn = true
	qb.q = qb.q.Where(path, OpTypeNotIn, value)
	return qb
}

// ArrayContains - array filter ( `array-contains` )
func (qb *QueryBuilder) ArrayContains(path string, value interface{}) *QueryBuilder {
	if qb.usedAny {
		qb.errs = append(qb.errs, newErrs(OpTypeArrayContains, path, cannotCombineAny, value))
		return qb
	}
	qb.q = qb.q.Where(path, OpTypeArrayContains, value)
	return qb
}

// ArrayContainsAny - array filter ( `array-contains-any` )
func (qb *QueryBuilder) ArrayContainsAny(path string, value interface{}) *QueryBuilder {
	switch rv := reflect.ValueOf(value); rv.Kind() {
	case reflect.Slice, reflect.Array:
		if rv.Len() > 10 {
			qb.errs = append(qb.errs, newErrs(OpTypeArrayContainsAny, path, notLessThan10, value))
			return qb
		}
	default:
		qb.errs = append(qb.errs, newErrs(OpTypeArrayContainsAny, path, allowArrayOrSlice, value))
		return qb
	}
	if qb.usedIn || qb.usedAny {
		qb.errs = append(qb.errs, newErrs(OpTypeArrayContainsAny, path, alreadyUsingOnce, value))
		return qb
	}
	qb.usedAny = true
	qb.q = qb.q.Where(path, OpTypeArrayContainsAny, value)
	return qb
}

// Asc - order
func (qb *QueryBuilder) Asc(path string) *QueryBuilder {
	if _, ok := qb.equalCounter[path]; ok || qb.usedIn {
		qb.errs = append(qb.errs, newErrs("Asc", path, alreadyEqualOrIn, ""))
		return qb
	}
	if _, ok := qb.orderCounter[path]; ok {
		qb.errs = append(qb.errs, newErrs("Asc", path, alreadySortOrder, ""))
		return qb
	}
	qb.q = qb.q.OrderBy(path, firestore.Asc)
	return qb
}

// Desc - order
func (qb *QueryBuilder) Desc(path string) *QueryBuilder {
	if _, ok := qb.equalCounter[path]; ok || qb.usedIn {
		qb.errs = append(qb.errs, newErrs("Desc", path, alreadyEqualOrIn, ""))
		return qb
	}
	if _, ok := qb.orderCounter[path]; ok {
		qb.errs = append(qb.errs, newErrs("Desc", path, alreadySortOrder, ""))
		return qb
	}
	qb.q = qb.q.OrderBy(path, firestore.Desc)
	return qb
}

// Limit - limit
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.q = qb.q.Limit(limit)
	return qb
}

// StartAt - start at
func (qb *QueryBuilder) StartAt(docSnapshotOrFieldValues ...interface{}) *QueryBuilder {
	qb.q = qb.q.StartAt(docSnapshotOrFieldValues...)
	return qb
}

// StartAfter - start after
func (qb *QueryBuilder) StartAfter(docSnapshotOrFieldValues ...interface{}) *QueryBuilder {
	qb.q = qb.q.StartAfter(docSnapshotOrFieldValues...)
	return qb
}

// EndAt - end at
func (qb *QueryBuilder) EndAt(docSnapshotOrFieldValues ...interface{}) *QueryBuilder {
	qb.q = qb.q.EndAt(docSnapshotOrFieldValues...)
	return qb
}

// EndBefore - end before
func (qb *QueryBuilder) EndBefore(docSnapshotOrFieldValues ...interface{}) *QueryBuilder {
	qb.q = qb.q.EndBefore(docSnapshotOrFieldValues...)
	return qb
}

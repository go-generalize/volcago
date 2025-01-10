package misc

import (
	"reflect"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/type/latlng"
)

type Meta struct {
	CreatedAt time.Time  `json:"createdAt" firestore:"createdAt"`
	CreatedBy string     `json:"createdBy" firestore:"createdBy"`
	UpdatedAt time.Time  `json:"updatedAt" firestore:"updatedAt"`
	UpdatedBy string     `json:"updatedBy" firestore:"updatedBy"`
	DeletedAt *time.Time `json:"deletedAt" firestore:"deletedAt"`
	DeletedBy string     `json:"deletedBy" firestore:"deletedBy"`
	Version   int        `json:"version"   firestore:"version"`
}

type article struct {
	User      user   `json:"user"      firestore:"user"`
	Page      string `json:"page"      firestore:"page"`
	Published bool   `json:"published" firestore:"published"`
	Price     *int   `json:"price"     firestore:"price"`
	Meta
}

type user struct {
	Name     string     `json:"name"     firestore:"name"`
	Age      int        `json:"age"      firestore:"age"`
	BirthDay *time.Time `json:"birthDay" firestore:"birthDay"`
	IsAdult  bool       `json:"isAdult"  firestore:"isAdult"`
	Address  address    `json:"address"  firestore:"address"`
}

type address struct {
	LatLng *latlng.LatLng `json:"latLng"`
}

type articleUpdateParam struct {
	User      interface{}
	Page      interface{}
	Published interface{}
	Price     interface{}
	CreatedAt interface{}
	CreatedBy interface{}
	UpdatedAt interface{}
	UpdatedBy interface{}
	DeletedAt interface{}
	DeletedBy interface{}
	Version   interface{}
}

func Test_IsReserveType(t *testing.T) {
	type args struct {
		value reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				value: reflect.ValueOf(1),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(2)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(int8(3)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(int8(4))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(int16(5)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(int16(6))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(int32(7)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(int32(8))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(int64(9)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(int64(10))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(uint(11)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(uint(12))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(uint8(13)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(uint8(14))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(uint16(15)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(uint16(16))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(uint32(17)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(uint16(18))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(uint64(19)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(uint64(20))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(float32(100.1)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(float32(100.2))),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(100.3),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(100.4)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(true),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(false)),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf("string"),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr("string ptr")),
			},
			want: true,
		},
		{
			args: args{
				value: reflect.ValueOf(article{}),
			},
			want: false,
		},
		{
			args: args{
				value: reflect.ValueOf(lo.ToPtr(article{})),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt // escape: Using the variable on range scope `tt` in loop literal
		t.Run(tt.name, func(t *testing.T) {
			got := isReservedType(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("updateBuilder() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}

func Test_updateBuilder(t *testing.T) {
	type args struct {
		v     interface{}
		param *articleUpdateParam
	}

	unix := time.Unix(0, 0)
	age := time.Now().Year() - unix.Year()
	latLng := &latlng.LatLng{
		Latitude:  35.678803,
		Longitude: 139.756263,
	}

	tests := []struct {
		name string
		args args
		want map[string]firestore.Update
	}{
		{
			args: args{
				v: article{},
				param: &articleUpdateParam{
					User: user{
						Name:     "john",
						Age:      age,
						BirthDay: &unix,
						IsAdult:  false,
						Address: address{
							LatLng: latLng,
						},
					},
					Page:      "section",
					Price:     2,
					Published: false,
					CreatedAt: unix,
					CreatedBy: "operator",
					UpdatedAt: unix,
					Version:   1,
				},
			},
			want: map[string]firestore.Update{
				"user.Name":           {FieldPath: firestore.FieldPath{"user", "name"}, Value: "john"},
				"user.Age":            {FieldPath: firestore.FieldPath{"user", "age"}, Value: age},
				"user.BirthDay":       {FieldPath: firestore.FieldPath{"user", "birthDay"}, Value: &unix},
				"user.address.LatLng": {FieldPath: firestore.FieldPath{"user", "address", "LatLng"}, Value: latLng},
				"Page":                {FieldPath: firestore.FieldPath{"page"}, Value: "section"},
				"Published":           {FieldPath: firestore.FieldPath{"published"}, Value: false},
				"Price":               {FieldPath: firestore.FieldPath{"price"}, Value: 2},
				"CreatedAt":           {FieldPath: firestore.FieldPath{"createdAt"}, Value: unix},
				"CreatedBy":           {FieldPath: firestore.FieldPath{"createdBy"}, Value: "operator"},
				"UpdatedAt":           {FieldPath: firestore.FieldPath{"updatedAt"}, Value: unix},
				"Version":             {FieldPath: firestore.FieldPath{"version"}, Value: 1},
			},
		},
	}

	for _, tt := range tests {
		tt := tt // escape: Using the variable on range scope `tt` in loop literal
		t.Run(tt.name, func(t *testing.T) {
			got := updateBuilder(tt.args.v, tt.args.param)
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(latlng.LatLng{})); diff != "" {
				t.Errorf("updateBuilder() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}

func Test_updater(t *testing.T) {
	type args struct {
		v     interface{}
		param *articleUpdateParam
	}

	unix := time.Unix(0, 0)
	age := time.Now().Year() - unix.Year()
	latLng := &latlng.LatLng{
		Latitude:  35.678803,
		Longitude: 139.756263,
	}

	tests := []struct {
		name string
		args args
		want []firestore.Update
	}{
		{
			args: args{
				v: article{},
				param: &articleUpdateParam{
					User: user{
						Name:     "john",
						Age:      age,
						BirthDay: &unix,
						IsAdult:  false,
						Address: address{
							LatLng: latLng,
						},
					},
					Page:      "section",
					Price:     2,
					Published: false,
					CreatedAt: unix,
					CreatedBy: "operator",
					UpdatedAt: unix,
					Version:   1,
				},
			},
			want: []firestore.Update{
				{FieldPath: firestore.FieldPath{"createdAt"}, Value: unix},
				{FieldPath: firestore.FieldPath{"createdBy"}, Value: "operator"},
				{FieldPath: firestore.FieldPath{"page"}, Value: "section"},
				{FieldPath: firestore.FieldPath{"price"}, Value: 2},
				{FieldPath: firestore.FieldPath{"published"}, Value: false},
				{FieldPath: firestore.FieldPath{"updatedAt"}, Value: unix},
				{FieldPath: firestore.FieldPath{"user", "address", "LatLng"}, Value: latLng},
				{FieldPath: firestore.FieldPath{"user", "age"}, Value: age},
				{FieldPath: firestore.FieldPath{"user", "birthDay"}, Value: &unix},
				{FieldPath: firestore.FieldPath{"user", "name"}, Value: "john"},
				{FieldPath: firestore.FieldPath{"version"}, Value: 1},
			},
		},
		{
			args: args{
				v: article{},
				param: &articleUpdateParam{
					Page: "section",
				},
			},
			want: []firestore.Update{
				{FieldPath: firestore.FieldPath{"page"}, Value: "section"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt // escape: Using the variable on range scope `tt` in loop literal
		t.Run(tt.name, func(t *testing.T) {
			got := updater(tt.args.v, tt.args.param)
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(latlng.LatLng{})); diff != "" {
				t.Errorf("updater() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}

func Test_tagMap(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			args: args{
				v: article{},
			},
			want: map[string]string{
				"CreatedAt":           "createdAt",
				"CreatedBy":           "createdBy",
				"DeletedAt":           "deletedAt",
				"DeletedBy":           "deletedBy",
				"Page":                "page",
				"Price":               "price",
				"Published":           "published",
				"UpdatedAt":           "updatedAt",
				"UpdatedBy":           "updatedBy",
				"Version":             "version",
				"user.Age":            "user.age",
				"user.BirthDay":       "user.birthDay",
				"user.IsAdult":        "user.isAdult",
				"user.Name":           "user.name",
				"user.address.LatLng": "user.address.LatLng",
			},
		},
	}

	for _, tt := range tests {
		tt := tt // escape: Using the variable on range scope `tt` in loop literal
		t.Run(tt.name, func(t *testing.T) {
			got := tagMap(tt.args.v)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("tagMap() = %v, want %v\n%s", got, tt.want, diff)
			}
		})
	}
}

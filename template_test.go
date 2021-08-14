package main

import (
	"bytes"
	"testing"
	"text/template"
)

type fields struct {
	tpl *template.Template
}
type args struct {
	templateName string
	data         layer
}

type testTable struct {
	name    string
	fields  fields
	args    args
	wantWr  string
	wantErr bool
}

type testTables []testTable

func TestTemplate_Create(t1 *testing.T) {
	tests := testTables{}
	tests = append(tests, getDomainLayerTests()...)

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := textTemplate{
				tpl: tt.fields.tpl,
			}
			wr := &bytes.Buffer{}
			err := t.create(wr, tt.args.templateName, tt.args.data)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWr := wr.String(); gotWr != tt.wantWr {
				t1.Errorf("Create() gotWr = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}

func getDomainLayerTests() testTables {
	tpl := template.Must(template.New("domain.gotpl").Funcs(getTemplateFunctions()).ParseFiles("templates/domain/domain.gotpl"))

	return testTables{
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "domain.gotpl",
				data: layer{
					ModelName:   "User",
					TableName:   "users",
					Fields:      nil,
					ProjectPath: "",
				},
			},
			wantWr: `package user

type UseCase interface {
	GetTx() (model.Transaction, error)

	Create(m *model.User) error
	Update(m *model.User) error
	Delete(ID uint) error

	GetWhere(fields model.Fields, sort model.SortFields) (model.User, error)
	GetAllWhere(fields model.Fields, sort model.SortFields, pag model.Pagination) (model.Users, error)
}
`,
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "domain.gotpl",
				data: layer{
					ModelName:   "UserLogin",
					TableName:   "user_logins",
					Fields:      nil,
					ProjectPath: "",
				},
			},
			wantWr: `package userlogin

type UseCase interface {
	GetTx() (model.Transaction, error)

	Create(m *model.UserLogin) error
	Update(m *model.UserLogin) error
	Delete(ID uint) error

	GetWhere(fields model.Fields, sort model.SortFields) (model.UserLogin, error)
	GetAllWhere(fields model.Fields, sort model.SortFields, pag model.Pagination) (model.UserLogins, error)
}
`,
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "domain.gotpl",
				data: layer{
					ModelName:   "UserLogin",
					TableName:   "user_logins",
					Fields:      nil,
					ProjectPath: "",
				},
			},
			wantWr: `package userlogin

type UseCase interface {
	GetTx() (model.Transaction, error)

	Create(m *model.UserLogin) error
	Update(m *model.UserLogin) error
	Delete(ID uint) error

	GetWhere(fields model.Fields, sort model.SortFields) (model.UserLogin, error)
	GetAllWhere(fields model.Fields, sort model.SortFields, pag model.Pagination) (model.UserLogins, error)
}
`,
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "domain.gotpl",
				data: layer{
					ModelName:   "Role",
					TableName:   "roles",
					Fields:      nil,
					ProjectPath: "",
				},
			},
			wantWr: `package role

type UseCase interface {
	GetTx() (model.Transaction, error)

	Create(m *model.Role) error
	Update(m *model.Role) error
	Delete(ID uint) error

	GetWhere(fields model.Fields, sort model.SortFields) (model.Role, error)
	GetAllWhere(fields model.Fields, sort model.SortFields, pag model.Pagination) (model.Roles, error)
}
`,
			wantErr: false,
		},
	}
}

func getDomainUseCaseLayerTests() testTables {
	tpl := template.Must(template.New("usecase.gotpl").Funcs(getTemplateFunctions()).ParseFiles("templates/domain/usecase.gotpl"))

	return testTables{
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "usecase.gotpl",
				data: layer{
					ModelName:   "User",
					TableName:   "users",
					Fields:      nil,
					ProjectPath: "",
				},
			},
			wantWr: `package user

type UseCase interface {
	GetTx() (model.Transaction, error)

	Create(m *model.User) error
	Update(m *model.User) error
	Delete(ID uint) error

	GetWhere(fields model.Fields, sort model.SortFields) (model.User, error)
	GetAllWhere(fields model.Fields, sort model.SortFields, pag model.Pagination) (model.Users, error)
}
`,
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "usecase.gotpl",
				data: layer{
					ModelName:   "UserLogin",
					TableName:   "user_logins",
					Fields:      nil,
					ProjectPath: "",
				},
			},
			wantWr: `package userlogin

type UseCase interface {
	GetTx() (model.Transaction, error)

	Create(m *model.UserLogin) error
	Update(m *model.UserLogin) error
	Delete(ID uint) error

	GetWhere(fields model.Fields, sort model.SortFields) (model.UserLogin, error)
	GetAllWhere(fields model.Fields, sort model.SortFields, pag model.Pagination) (model.UserLogins, error)
}
`,
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "usecase.gotpl",
				data: layer{
					ModelName:   "UserLogin",
					TableName:   "user_logins",
					Fields:      nil,
					ProjectPath: "",
				},
			},
			wantWr: `package userlogin

type UseCase interface {
	GetTx() (model.Transaction, error)

	Create(m *model.UserLogin) error
	Update(m *model.UserLogin) error
	Delete(ID uint) error

	GetWhere(fields model.Fields, sort model.SortFields) (model.UserLogin, error)
	GetAllWhere(fields model.Fields, sort model.SortFields, pag model.Pagination) (model.UserLogins, error)
}
`,
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "usecase.gotpl",
				data: layer{
					ModelName:   "Role",
					TableName:   "roles",
					Fields:      nil,
					ProjectPath: "",
				},
			},
			wantWr: `package role

type UseCase interface {
	GetTx() (model.Transaction, error)

	Create(m *model.Role) error
	Update(m *model.Role) error
	Delete(ID uint) error

	GetWhere(fields model.Fields, sort model.SortFields) (model.Role, error)
	GetAllWhere(fields model.Fields, sort model.SortFields, pag model.Pagination) (model.Roles, error)
}
`,
			wantErr: false,
		},
	}
}

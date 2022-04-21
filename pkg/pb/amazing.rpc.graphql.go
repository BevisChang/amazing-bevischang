// Code generated by protoc-gen-svc. DO NOT EDIT.
package pb

import (
	"context"
	"errors"
	"time"

	"github.com/graphql-go/graphql"
	ms "github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
)

const BevisChangGrpcContextKey = "grpc-client-BevisChang"

type deferFunc func()

func ContextWithGrpcClient(ctx context.Context, addr string) (context.Context, deferFunc, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, err
	}
	client := NewBevisChangClient(conn)
	handleDeferFunc := func() {
		conn.Close()
	}
	return context.WithValue(ctx, BevisChangGrpcContextKey, &client), handleDeferFunc, nil
}

func RefiningBevisChangGrpcClientFromContext(ctx context.Context) (*BevisChangClient, error) {
	client, ok := ctx.Value(BevisChangGrpcContextKey).(*BevisChangClient)
	if !ok {
		return nil, errors.New("RefiningBevisChangGrpcClientFromContext failed")
	}
	return client, nil
}

var ListMembersReqObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListMembersReqObject",
	Fields: graphql.Fields{
		"birthdayBefore": &graphql.Field{Type: graphql.String},
	},
	Description: "",
})

var GetRecordReqObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetRecordReqObject",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
	},
	Description: "",
})

var GetRecordResObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetRecordResObject",
	Fields: graphql.Fields{
		"record": &graphql.Field{Type: RecordObject},
	},
	Description: "",
})

var ListRecordReqObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListRecordReqObject",
	Fields: graphql.Fields{
		"size": &graphql.Field{Type: graphql.String},
		"page": &graphql.Field{Type: graphql.String},
	},
	Description: "",
})

var MemberObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "MemberObject",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"name":       &graphql.Field{Type: graphql.String},
		"birthday":   &graphql.Field{Type: graphql.Int},
		"created_at": &graphql.Field{Type: graphql.String},
		"updated_at": &graphql.Field{Type: graphql.String},
	},
	Description: "",
})

var HealthReqObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "HealthReqObject",
	Fields:      graphql.Fields{},
	Description: "",
})

var HealthResObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "HealthResObject",
	Fields: graphql.Fields{
		"ok": &graphql.Field{Type: graphql.Boolean},
	},
	Description: "",
})

var CreateMemberResObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateMemberResObject",
	Fields: graphql.Fields{
		"member": &graphql.Field{Type: MemberObject},
	},
	Description: "",
})

var CreateMemberReqObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateMemberReqObject",
	Fields: graphql.Fields{
		"name":     &graphql.Field{Type: graphql.String},
		"birthday": &graphql.Field{Type: graphql.Int},
	},
	Description: "",
})

var UpdateMemberReqObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "UpdateMemberReqObject",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.String},
		"name":     &graphql.Field{Type: graphql.String},
		"birthday": &graphql.Field{Type: graphql.Int},
	},
	Description: "",
})

var ListMembersResObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListMembersResObject",
	Fields: graphql.Fields{
		"member": &graphql.Field{Type: graphql.NewList(MemberObject)},
	},
	Description: "",
})

var ConfigReqObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ConfigReqObject",
	Fields:      graphql.Fields{},
	Description: "",
})

var ConfigResObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "ConfigResObject",
	Fields: graphql.Fields{
		"enable": &graphql.Field{Type: graphql.Boolean},
		"num":    &graphql.Field{Type: graphql.Int},
		"str":    &graphql.Field{Type: graphql.String},
	},
	Description: "",
})

var CreateRecordResObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateRecordResObject",
	Fields: graphql.Fields{
		"record": &graphql.Field{Type: RecordObject},
	},
	Description: "",
})

var ListRecordResObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListRecordResObject",
	Fields: graphql.Fields{
		"records": &graphql.Field{Type: graphql.NewList(RecordObject)},
	},
	Description: "",
})

var DeleteMemberResObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DeleteMemberResObject",
	Fields:      graphql.Fields{},
	Description: "",
})

var RecordObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "RecordObject",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.String},
		"the_num":    &graphql.Field{Type: graphql.Int},
		"the_str":    &graphql.Field{Type: graphql.String},
		"created_at": &graphql.Field{Type: graphql.String},
		"updated_at": &graphql.Field{Type: graphql.String},
	},
	Description: "",
})

var CreateRecordReqObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateRecordReqObject",
	Fields: graphql.Fields{
		"the_num":    &graphql.Field{Type: graphql.Int},
		"the_str":    &graphql.Field{Type: graphql.String},
		"created_at": &graphql.Field{Type: graphql.String},
	},
	Description: "",
})

var UpdateMemberResObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "UpdateMemberResObject",
	Fields: graphql.Fields{
		"member": &graphql.Field{Type: MemberObject},
	},
	Description: "",
})

var DeleteMemberReqObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeleteMemberReqObject",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
	},
	Description: "",
})

var HealthArguments = graphql.FieldConfigArgument{}

var HealthQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "HealthQueryType",
	Fields: graphql.Fields{
		"ok": &graphql.Field{Type: graphql.Boolean},
	},
	Description: "",
})

func BevisChangHealthResolver(p graphql.ResolveParams) (interface{}, error) {
	type result struct {
		data interface{}
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		defer close(ch)

		client, err := RefiningBevisChangGrpcClientFromContext(p.Context)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		req := HealthReq{}
		if len(p.Args) != 0 {
			err = ms.Decode(p.Args, &req)
			if err != nil {
				ch <- result{data: nil, err: err}
				return
			}
		}

		res, err := (*client).Health(ctx, &req)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}
		ch <- result{data: res, err: nil}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.err
	}, nil
}

var ConfigArguments = graphql.FieldConfigArgument{}

var ConfigQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ConfigQueryType",
	Fields: graphql.Fields{
		"enable": &graphql.Field{Type: graphql.Boolean},
		"num":    &graphql.Field{Type: graphql.Int},
		"str":    &graphql.Field{Type: graphql.String},
	},
	Description: "",
})

func BevisChangConfigResolver(p graphql.ResolveParams) (interface{}, error) {
	type result struct {
		data interface{}
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		defer close(ch)

		client, err := RefiningBevisChangGrpcClientFromContext(p.Context)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		req := ConfigReq{}
		if len(p.Args) != 0 {
			err = ms.Decode(p.Args, &req)
			if err != nil {
				ch <- result{data: nil, err: err}
				return
			}
		}

		res, err := (*client).Config(ctx, &req)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}
		ch <- result{data: res, err: nil}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.err
	}, nil
}

var CreateRecordArguments = graphql.FieldConfigArgument{
	"the_num":    &graphql.ArgumentConfig{Type: graphql.Int},
	"the_str":    &graphql.ArgumentConfig{Type: graphql.String},
	"created_at": &graphql.ArgumentConfig{Type: graphql.String},
}

var CreateRecordQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateRecordQueryType",
	Fields: graphql.Fields{
		"record": &graphql.Field{Type: RecordObject},
	},
	Description: "",
})

func BevisChangCreateRecordResolver(p graphql.ResolveParams) (interface{}, error) {
	type result struct {
		data interface{}
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		defer close(ch)

		client, err := RefiningBevisChangGrpcClientFromContext(p.Context)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		req := CreateRecordReq{}
		if len(p.Args) != 0 {
			err = ms.Decode(p.Args, &req)
			if err != nil {
				ch <- result{data: nil, err: err}
				return
			}
		}

		res, err := (*client).CreateRecord(ctx, &req)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}
		ch <- result{data: res, err: nil}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.err
	}, nil
}

var GetRecordArguments = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{Type: graphql.String},
}

var GetRecordQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetRecordQueryType",
	Fields: graphql.Fields{
		"record": &graphql.Field{Type: RecordObject},
	},
	Description: "",
})

func BevisChangGetRecordResolver(p graphql.ResolveParams) (interface{}, error) {
	type result struct {
		data interface{}
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		defer close(ch)

		client, err := RefiningBevisChangGrpcClientFromContext(p.Context)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		req := GetRecordReq{}
		if len(p.Args) != 0 {
			err = ms.Decode(p.Args, &req)
			if err != nil {
				ch <- result{data: nil, err: err}
				return
			}
		}

		res, err := (*client).GetRecord(ctx, &req)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}
		ch <- result{data: res, err: nil}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.err
	}, nil
}

var ListRecordArguments = graphql.FieldConfigArgument{
	"size": &graphql.ArgumentConfig{Type: graphql.String},
	"page": &graphql.ArgumentConfig{Type: graphql.String},
}

var ListRecordQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListRecordQueryType",
	Fields: graphql.Fields{
		"records": &graphql.Field{Type: graphql.NewList(RecordObject)},
	},
	Description: "",
})

func BevisChangListRecordResolver(p graphql.ResolveParams) (interface{}, error) {
	type result struct {
		data interface{}
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		defer close(ch)

		client, err := RefiningBevisChangGrpcClientFromContext(p.Context)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		req := ListRecordReq{}
		if len(p.Args) != 0 {
			err = ms.Decode(p.Args, &req)
			if err != nil {
				ch <- result{data: nil, err: err}
				return
			}
		}

		res, err := (*client).ListRecord(ctx, &req)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}
		ch <- result{data: res, err: nil}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.err
	}, nil
}

var CreateMemberArguments = graphql.FieldConfigArgument{
	"name":     &graphql.ArgumentConfig{Type: graphql.String},
	"birthday": &graphql.ArgumentConfig{Type: graphql.Int},
}

var CreateMemberQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateMemberQueryType",
	Fields: graphql.Fields{
		"member": &graphql.Field{Type: MemberObject},
	},
	Description: "",
})

func BevisChangCreateMemberResolver(p graphql.ResolveParams) (interface{}, error) {
	type result struct {
		data interface{}
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		defer close(ch)

		client, err := RefiningBevisChangGrpcClientFromContext(p.Context)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		req := CreateMemberReq{}
		if len(p.Args) != 0 {
			err = ms.Decode(p.Args, &req)
			if err != nil {
				ch <- result{data: nil, err: err}
				return
			}
		}

		res, err := (*client).CreateMember(ctx, &req)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}
		ch <- result{data: res, err: nil}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.err
	}, nil
}

var UpdateMemberArguments = graphql.FieldConfigArgument{
	"id":       &graphql.ArgumentConfig{Type: graphql.String},
	"name":     &graphql.ArgumentConfig{Type: graphql.String},
	"birthday": &graphql.ArgumentConfig{Type: graphql.Int},
}

var UpdateMemberQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UpdateMemberQueryType",
	Fields: graphql.Fields{
		"member": &graphql.Field{Type: MemberObject},
	},
	Description: "",
})

func BevisChangUpdateMemberResolver(p graphql.ResolveParams) (interface{}, error) {
	type result struct {
		data interface{}
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		defer close(ch)

		client, err := RefiningBevisChangGrpcClientFromContext(p.Context)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		req := UpdateMemberReq{}
		if len(p.Args) != 0 {
			err = ms.Decode(p.Args, &req)
			if err != nil {
				ch <- result{data: nil, err: err}
				return
			}
		}

		res, err := (*client).UpdateMember(ctx, &req)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}
		ch <- result{data: res, err: nil}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.err
	}, nil
}

var ListMembersArguments = graphql.FieldConfigArgument{
	"birthdayBefore": &graphql.ArgumentConfig{Type: graphql.String},
}

var ListMembersQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListMembersQueryType",
	Fields: graphql.Fields{
		"member": &graphql.Field{Type: graphql.NewList(MemberObject)},
	},
	Description: "",
})

func BevisChangListMembersResolver(p graphql.ResolveParams) (interface{}, error) {
	type result struct {
		data interface{}
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		defer close(ch)

		client, err := RefiningBevisChangGrpcClientFromContext(p.Context)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		req := ListMembersReq{}
		if len(p.Args) != 0 {
			err = ms.Decode(p.Args, &req)
			if err != nil {
				ch <- result{data: nil, err: err}
				return
			}
		}

		res, err := (*client).ListMembers(ctx, &req)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}
		ch <- result{data: res, err: nil}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.err
	}, nil
}

var DeleteMemberArguments = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{Type: graphql.String},
}

var DeleteMemberQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DeleteMemberQueryType",
	Fields:      graphql.Fields{},
	Description: "",
})

func BevisChangDeleteMemberResolver(p graphql.ResolveParams) (interface{}, error) {
	type result struct {
		data interface{}
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		defer close(ch)

		client, err := RefiningBevisChangGrpcClientFromContext(p.Context)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		req := DeleteMemberReq{}
		if len(p.Args) != 0 {
			err = ms.Decode(p.Args, &req)
			if err != nil {
				ch <- result{data: nil, err: err}
				return
			}
		}

		res, err := (*client).DeleteMember(ctx, &req)
		if err != nil {
			ch <- result{data: nil, err: err}
			return
		}
		ch <- result{data: res, err: nil}
	}()
	return func() (interface{}, error) {
		r := <-ch
		return r.data, r.err
	}, nil
}

var internalBevisChangRootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "BevisChangQuery",
	Fields: graphql.Fields{
		"Health": &graphql.Field{
			Name:    "Health",
			Type:    HealthQueryType,
			Args:    HealthArguments,
			Resolve: BevisChangHealthResolver,
		},
		"Config": &graphql.Field{
			Name:    "Config",
			Type:    ConfigQueryType,
			Args:    ConfigArguments,
			Resolve: BevisChangConfigResolver,
		},
		"GetRecord": &graphql.Field{
			Name:    "GetRecord",
			Type:    GetRecordQueryType,
			Args:    GetRecordArguments,
			Resolve: BevisChangGetRecordResolver,
		},
		"ListRecord": &graphql.Field{
			Name:    "ListRecord",
			Type:    ListRecordQueryType,
			Args:    ListRecordArguments,
			Resolve: BevisChangListRecordResolver,
		},
		"ListMembers": &graphql.Field{
			Name:    "ListMembers",
			Type:    ListMembersQueryType,
			Args:    ListMembersArguments,
			Resolve: BevisChangListMembersResolver,
		},
	},
})

var BevisChangRootQueryField = graphql.Field{
	Name: "BevisChang",
	Type: internalBevisChangRootQuery,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return func() (interface{}, error) {
			return p.Info, nil
		}, nil
	},
}

var internalBevisChangRootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "BevisChangMutation",
	Fields: graphql.Fields{
		"CreateRecord": &graphql.Field{
			Name:    "CreateRecord",
			Type:    CreateRecordQueryType,
			Args:    CreateRecordArguments,
			Resolve: BevisChangCreateRecordResolver,
		},
		"CreateMember": &graphql.Field{
			Name:    "CreateMember",
			Type:    CreateMemberQueryType,
			Args:    CreateMemberArguments,
			Resolve: BevisChangCreateMemberResolver,
		},
		"UpdateMember": &graphql.Field{
			Name:    "UpdateMember",
			Type:    UpdateMemberQueryType,
			Args:    UpdateMemberArguments,
			Resolve: BevisChangUpdateMemberResolver,
		},
		"DeleteMember": &graphql.Field{
			Name:    "DeleteMember",
			Type:    DeleteMemberQueryType,
			Args:    DeleteMemberArguments,
			Resolve: BevisChangDeleteMemberResolver,
		},
	},
})

var BevisChangRootMutationField = graphql.Field{
	Name: "BevisChang",
	Type: internalBevisChangRootMutation,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return func() (interface{}, error) {
			return p.Info, nil
		}, nil
	},
}
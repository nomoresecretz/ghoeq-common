package object

import (
	"reflect"

	"github.com/nomoresecretz/ghoeq-common/eqStruct"
)

var StructActions = []struct {
	Name         string
	StructAction *StructAction
}{
	{
		"EQRead",
		&StructAction{Fn: func(s eqStruct.EQStruct, r reflect.Value, b []byte, args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of args: got=%d, want=2", len(args))
			}

			var field string
			switch arg := args[1].(type) {
			case *String:
				field = arg.Value
			default:
				return newError("argument to `EQRead` not supported, got=%s", args[0].Type())
			}

			var l int
			switch arg := args[1].(type) {
			case *Integer:
				l = int(arg.Value)
			default:
				return newError("argument to `EQRead` not supported, got=%s", args[0].Type())
			}

			f := r.FieldByName(field).Addr()
			if err := eqStruct.EQRead(b, s, f, l); err != nil {
				return newError("err: ", err)
			}

			return nil
		}},
	},
	{
		"EQReadBigEndian",
		&StructAction{Fn: func(s eqStruct.EQStruct, r reflect.Value, b []byte, args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of args: got=%d, want=2", len(args))
			}

			var field string
			switch arg := args[1].(type) {
			case *String:
				field = arg.Value
			default:
				return newError("argument to `EQRead` not supported, got=%s", args[0].Type())
			}

			var l int
			switch arg := args[1].(type) {
			case *Integer:
				l = int(arg.Value)
			default:
				return newError("argument to `EQRead` not supported, got=%s", args[0].Type())
			}

			f := r.FieldByName(field).Addr()
			if err := eqStruct.EQReadBigEndian(b, s, f, l); err != nil {
				return newError("err: ", err)
			}

			return nil
		}},
	},
	{
		"SetPointer",
		&StructAction{Fn: func(s eqStruct.EQStruct, r reflect.Value, b []byte, args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of args: got=%d, want=2", len(args))
			}

			var p int
			switch arg := args[1].(type) {
			case *Integer:
				p = int(arg.Value)
			default:
				return newError("argument to `EQRead` not supported, got=%s", args[0].Type())
			}

			s.SetPointer(p)

			return nil
		}},
	},
}

func GetEQActionByName(name string) *StructAction {
	for _, def := range StructActions {
		if def.Name == name {
			return def.StructAction
		}
	}
	return nil
}

package ros

import (
	"strings"

	"github.com/samber/lo"

	"github.com/0x587/guardeye/test-client/reporter/ros/rosmsg"
	"github.com/0x587/guardeye/test-client/reporter/ros/rosmsg/walk"
)

func (i *impl) getInterfaceMsgDefine(typeUrl string) (*MsgDefine, error) {
	typePaths := strings.Split(typeUrl, "/")
	packageName, typeName := typePaths[0], typePaths[len(typePaths)-1]
	interfaceDefine, err := i.rosInterfaceShow(typeUrl)
	if err != nil {
		return nil, err
	}
	tree, err := rosmsg.ParseMsg(interfaceDefine)
	if err != nil {
		return nil, err
	}
	msg := walk.Walk(tree)[0]
	block := &BlockDefine{
		Fields: lo.Map(msg.Fields, func(f *walk.FieldDefine, _ int) *FieldDefine {
			return &FieldDefine{
				Name:        f.Name,
				Type:        f.Type.Name,
				IsArray:     f.Type.IsArray,
				IsComplex:   f.Type.IsCustom,
				ComplexType: nil,
			}
		}),
	}
	if err := i.processMsg(block, packageName); err != nil {
		return nil, err
	}
	return &MsgDefine{
		Package: packageName,
		Name:    typeName,
		Block:   block,
	}, nil
}

func (i *impl) getInterfaceSrvDefine(typeUrl string) (*SrvDefine, error) {
	typePaths := strings.Split(typeUrl, "/")
	packageName, typeName := typePaths[0], typePaths[len(typePaths)-1]
	interfaceDefine, err := i.rosInterfaceShow(typeUrl)
	if err != nil {
		return nil, err
	}
	tree, err := rosmsg.ParseSrv(interfaceDefine)
	if err != nil {
		return nil, err
	}
	msg := walk.Walk(tree)[0]
	req := &BlockDefine{
		Fields: lo.Map(msg.Fields, func(f *walk.FieldDefine, _ int) *FieldDefine {
			return &FieldDefine{
				Name:        f.Name,
				Type:        f.Type.Name,
				IsArray:     f.Type.IsArray,
				IsComplex:   f.Type.IsCustom,
				ComplexType: nil,
			}
		}),
	}
	if err := i.processMsg(req, packageName); err != nil {
		return nil, err
	}
	msg = walk.Walk(tree)[1]
	rsp := &BlockDefine{
		Fields: lo.Map(msg.Fields, func(f *walk.FieldDefine, _ int) *FieldDefine {
			return &FieldDefine{
				Name:        f.Name,
				Type:        f.Type.Name,
				IsArray:     f.Type.IsArray,
				IsComplex:   f.Type.IsCustom,
				ComplexType: nil,
			}
		}),
	}
	if err := i.processMsg(rsp, packageName); err != nil {
		return nil, err
	}
	return &SrvDefine{
		Package: packageName,
		Name:    typeName,
		Req:     req,
		Rsp:     rsp,
	}, nil
}

func (i *impl) processMsg(msg *BlockDefine, packageName string) error {
	for _, f := range msg.Fields {
		fieldPackageName := packageName
		if !f.IsComplex {
			continue
		}
		fullTypeName := f.Type
		ns := strings.Split(f.Type, "/")
		if len(ns) == 1 {
			fullTypeName = fieldPackageName + "/msg/" + f.Type
		}
		if len(ns) == 2 {
			fullTypeName = ns[0] + "/msg/" + ns[1]
			f.Type = ns[1]
			fieldPackageName = ns[0]
		}
		interfaceDefine, err := i.rosInterfaceShow(fullTypeName)
		if err != nil {
			return err
		}
		tree, err := rosmsg.ParseMsg(interfaceDefine)
		if err != nil {
			return err
		}
		m := walk.Walk(tree)[0]
		block := &BlockDefine{
			Fields: lo.Map(m.Fields, func(f *walk.FieldDefine, _ int) *FieldDefine {
				return &FieldDefine{
					Name:        f.Name,
					Type:        f.Type.Name,
					IsArray:     f.Type.IsArray,
					IsComplex:   f.Type.IsCustom,
					ComplexType: nil,
				}
			}),
		}
		if err := i.processMsg(block, fieldPackageName); err != nil {
			return err
		}
		f.ComplexType = &MsgDefine{
			Package: fieldPackageName,
			Name:    f.Type,
			Block:   block,
		}
	}
	return nil
}

var interfaceCache = make(map[string]string)

func (i *impl) rosInterfaceShow(typeName string) (string, error) {
	if v, ok := interfaceCache[typeName]; ok {
		return v, nil
	}
	out, err := i.rosExec("ros2 interface show --no-comments " + typeName)
	if err != nil {
		return "", err
	}
	lines := strings.Split(out, "\n")
	lines = lo.Filter(lines, func(l string, _ int) bool {
		return !strings.HasPrefix(l, "\t") && len(l) > 0
	})
	interfaceCache[typeName] = strings.Join(lines, "\n")
	return strings.Join(lines, "\n"), nil
}

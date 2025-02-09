package main

type ObjType int

const (
	OBJ_STRING ObjType = iota
	OBJ_FUNC
)

type Object struct {
	objType ObjType
}

func (o *Object) ObjType() ObjType {
	return o.objType
}

type ObjectString struct {
	obj    Object
	length int
	chars  string
}

type ObjectFunc struct {
	obj   Object
	arity int
	name  *ObjectString
	chunk Chunk
}

func NewFunction() *ObjectFunc {
	objFunc := ObjectFunc{
		name:  nil,
		chunk: NewChunck(),
		arity: 0,
		obj:   Object{OBJ_FUNC},
	}
	return &objFunc
}

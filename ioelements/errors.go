package ioelements

import "fmt"

type FallbackIOValueError struct {
	Model string
	AvlId uint16
	Err   error
}

func (e FallbackIOValueError) Error() string {
	return fmt.Sprintf("decode io element for model %q avl id %d", e.Model, e.AvlId)
}

func (e FallbackIOValueError) Unwrap() error {
	return e.Err
}

//type UndefinedElementId struct {
//	Model string
//	AvlId uint16
//	Err   error
//}
//
//func (e UndefinedElementId) Error() string {
//	return fmt.Sprintf("undefined io element for model %q avl id %d", e.Model, e.AvlId)
//}
//
//func (e UndefinedElementId) Unwrap() error {
//	return e.Err
//}

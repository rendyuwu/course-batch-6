package helper

import "github.com/jinzhu/copier"

func CopyStruct(to any, from any) {
	copier.CopyWithOption(to, from, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}

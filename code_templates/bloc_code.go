package code_templates

import "fmt"

const blocCode = `import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

part "%v"
part "%v"

class %v extends Bloc<%v, %v> {
  %v() : super(%v()) {}
}`

func GetBlocCode(config BlocConfig) string {
	code := fmt.Sprintf(blocCode,
		config.StateFileName,
		config.EventFileName,
		config.BlocName,
		config.EventName,
		config.StateName,
		config.BlocName,
		config.StateName)

	return code
}

type BlocConfig struct {
	BlocName      string
	StateName     string
	StateFileName string
	EventName     string
	EventFileName string
}

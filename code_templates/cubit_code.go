package code_templates

import "fmt"

func CubitCode(cubitName string, stateName string) string {
	code := `import 'package:flutter_bloc/flutter_bloc.dart';

class %v extends Cubit<%v> {
			void doSomething() {}
}`

	return fmt.Sprintf(code, cubitName, stateName)
}

func CubitCodeWithCustomState(stateFileName string, cubitName string, stateName string) string {
	code := `import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:equatable/equatable.dart';

part '%v';

class %v extends Cubit<%v> {
	void doSomething() {}
}`

	return fmt.Sprintf(code, stateFileName, cubitName, stateName)
}

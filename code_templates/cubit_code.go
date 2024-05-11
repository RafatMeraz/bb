package code_templates

import (
	"bb/utility"
	"fmt"
)

func CubitCode(cubitName string, stateName string) string {
	code := `import 'package:flutter_bloc/flutter_bloc.dart';

class %v extends Cubit<%v> {
	%v(): super(%v());

			void doSomething() {}
}`

	return fmt.Sprintf(code, cubitName, stateName, cubitName, utility.CreateInitialStateName(stateName))
}

func CubitCodeWithCustomState(stateFileName string, cubitName string, stateName string) string {
	code := `import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:equatable/equatable.dart';

part '%v';

class %v extends Cubit<%v> {
	%v(): super(%v());

	void doSomething() {}
}`

	return fmt.Sprintf(code, stateFileName, cubitName, stateName, cubitName, utility.CreateInitialStateName(stateName))
}

package code_templates

import "fmt"

func CubitCode(cubitName string, stateName string) string {
	return fmt.Sprintf(`import 'package:flutter_bloc/flutter_bloc.dart';

class %v extends Cubit<%v> {
			void doSomething() {}
}`, cubitName, stateName)
}

func CubitCodeWithCustomState(stateFileName string, cubitName string, stateName string) string {
	return fmt.Sprintf(`import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:equatable/equatable.dart';

part '%v';

class %v extends Cubit<%v> {
			void doSomething() {}
}`, stateFileName, cubitName, stateName)
}

func StateCode(stateName string, cubitFileName string) string {
	return fmt.Sprintf(`part of '%v';

sealed class %v extends Equatable {}
`, cubitFileName, stateName)
}

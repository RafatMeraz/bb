package code_templates

import (
	"bb/utility"
	"fmt"
)

func StateCodeWithOutPartOf(stateName string) string {
	return fmt.Sprintf(`import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:equatable/equatable.dart';

sealed class %v extends Equatable {}

class %v extends %v {
	@override
    List<Object?> get props => [];
}
`, stateName, utility.CreateInitialStateName(stateName), stateName)
}

func StateCodeWithPartOf(stateName string, partOfFileName string) string {
	return fmt.Sprintf(`part of '%v';

sealed class %v extends Equatable {}

class %v extends %v {
	@override
    List<Object?> get props => [];
}
`, partOfFileName, stateName, utility.CreateInitialStateName(stateName), stateName)
}

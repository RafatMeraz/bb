package code_templates

import "fmt"

func StateCodeWithOutPartOf(stateName string) string {
	return fmt.Sprintf(`import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:equatable/equatable.dart';

sealed class %v extends Equatable {}
`, stateName)
}

func StateCodeWithPartOf(stateName string, partOfFileName string) string {
	return fmt.Sprintf(`part of '%v';

sealed class %v extends Equatable {}
`, partOfFileName, stateName)
}

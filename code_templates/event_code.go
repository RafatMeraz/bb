package code_templates

import "fmt"

const eventWithoutPartOf = `import 'package:equatable/equatable.dart';

abstract class %v extends Equatable {}
`

const eventWithPartOf = `part of '%v';

abstract class %v extends Equatable {}
`

func GetEventCodeWithPartOf(eventName string, partOfFileName string) string {
	code := fmt.Sprintf(eventWithPartOf, partOfFileName, eventName)

	return code
}

func GetEventCodeWithoutPartOf(eventName string) string {
	code := fmt.Sprintf(eventWithoutPartOf, eventName)

	return code
}

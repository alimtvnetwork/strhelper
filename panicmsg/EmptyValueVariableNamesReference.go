package panicmsg

func EmptyValueVariableNamesReference(names ...string) *[]ReferenceValue {
	references := make([]ReferenceValue, len(names))

	for i, name := range names {
		references[i] = ReferenceValue{
			VariableName: name,
			Value:        "",
		}
	}

	return &references
}

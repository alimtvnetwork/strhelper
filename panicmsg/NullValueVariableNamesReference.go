package panicmsg

func NullValueVariableNamesReference(names ...string) *[]ReferenceValue {
	references := make([]ReferenceValue, len(names))

	for i, name := range names {
		references[i] = ReferenceValue{
			VariableName: name,
			Value:        nil,
		}
	}

	return &references
}

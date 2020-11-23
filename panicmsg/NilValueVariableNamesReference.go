package panicmsg

func NilValueVariableNamesReference(names ...string) *[]ReferenceValue {
	references := make([]ReferenceValue, len(names))

	for i, name := range names {
		references[i] = ReferenceValue{
			VariableName: name,
			Value:        nil,
		}
	}

	return &references
}

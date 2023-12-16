	// note(Julian-Chu): SetAttributes API without any attributes
	// will return MalformedInput(message: End of list found where not expected) error. We need to set minimal one default value,
	// if there are no attributes in the input
	if len(input.Attributes) == 0 {
		input.Attributes = map[string]*string{
			"DelaySeconds": latest.ko.Spec.DelaySeconds,
		}
	}

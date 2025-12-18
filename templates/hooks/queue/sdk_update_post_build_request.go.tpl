	// ignore updates if attributes are not defined.
	if len(input.Attributes) == 0 {
		return &resource{desired.ko.DeepCopy()}, nil
	}

	if tags, err := rm.getTags(ctx, r); err != nil {
		return nil, err
	} else {
		ko.Spec.Tags = FromACKTags(tags)
	}

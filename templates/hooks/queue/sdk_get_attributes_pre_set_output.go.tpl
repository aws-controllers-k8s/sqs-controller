	// If the QueueName field is empty, populate it with the last part of the queue ARN
	// This is a workaround for the fact that the QueueName field is required by the
	// CreateQueue API call, but not by the GetQueueAttributes API call
	// Use case: adopting an existing queue by queue URL
	if ko.Spec.QueueName == nil {
		split := strings.Split(string(tmpARN), ":")
		ko.Spec.QueueName = &split[len(split)-1]
	}

	// If the QueueName field is empty, populate it with the last part of the queue ARN
	// This is a workaround for the fact that the QueueName field is required by the
	// CreateQueue API call, but not by the GetQueueAttributes API call
	// Use case: adopting an existing queue by queue URL
	if ko.Spec.QueueName == nil {
		queueName, err := rm.getQueueNameFromARN(tmpARN)
		if err != nil {
			return nil, err
		}
		ko.Spec.QueueName = &queueName
	}

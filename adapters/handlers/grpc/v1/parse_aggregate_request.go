//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package v1

import (
	"github.com/weaviate/weaviate/entities/aggregation"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
	pb "github.com/weaviate/weaviate/grpc/generated/protocol/v1"
)

type AggregateParser struct {
	authorizedGetClass func(string) (*models.Class, error)
}

func NewAggregateParser(authorizedGetClass func(string) (*models.Class, error)) *AggregateParser {
	return &AggregateParser{
		authorizedGetClass: authorizedGetClass,
	}
}

func (p *AggregateParser) Aggregate(req *pb.AggregateRequest) (*aggregation.Params, error) {
	params := &aggregation.Params{}
	class, err := p.authorizedGetClass(req.Collection)
	if err != nil {
		return nil, err
	}

	params.ClassName = schema.ClassName(class.Class)
	params.IncludeMetaCount = req.MetaCount
	params.Tenant = req.Tenant

	if req.ObjectLimit != nil {
		objectLimit := int(*req.ObjectLimit)
		params.ObjectLimit = &objectLimit
	}

	return params, nil
}

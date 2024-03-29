package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"

	"github.com/Gromitmugs/temporal-playground/thirdparty/internal/graph/gqlgen"
	"github.com/Gromitmugs/temporal-playground/thirdparty/internal/graph/gqlmodel"
)

// MessageCreate is the resolver for the messageCreate field.
func (r *mutationResolver) MessageCreate(ctx context.Context, content string) (*gqlmodel.Message, error) {
	r.DBLastIndex += 1
	r.DB[r.DBLastIndex] = content

	return &gqlmodel.Message{
		ID:      r.DBLastIndex,
		Content: content,
	}, nil
}

// ErrorCreate is the resolver for the errorCreate field.
func (r *mutationResolver) ErrorCreate(ctx context.Context, errMsg string) (string, error) {
	return "", fmt.Errorf("an error occured: %s", errMsg)
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*gqlmodel.Message, error) {
	result := []*gqlmodel.Message{}
	for i := 0; i < r.DBLastIndex+1; i++ {
		result = append(result, &gqlmodel.Message{
			ID:      i,
			Content: r.DB[i],
		})
	}
	return result, nil
}

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

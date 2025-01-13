package resolvers

import (
	"context"
	"fmt"
	"github.com/GabsMeloTI/go-graphql/internal/generated"
	"github.com/GabsMeloTI/go-graphql/internal/model"
	"github.com/GabsMeloTI/go-graphql/internal/repository"
)

type InterfaceResolver interface {
}

type ResolverBasicInfo struct {
	InterfaceResolver repository.InterfaceRepository
}

func NewBasicInfoResolver(InterfaceResolver repository.InterfaceRepository) *ResolverBasicInfo {
	return &ResolverBasicInfo{InterfaceResolver}
}

func (r *mutationResolver) CreateBasicInfo(ctx context.Context, data *model.BasicInfoInput) (*model.BasicInfo, error) {
	newBasicInfo := model.BasicInfo{
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		AdditionalName: *data.AdditionalName,
		Pronouns:       data.Pronouns,
		HeadLine:       data.HeadLine,
	}

	query := `
		INSERT INTO basic_info (first_name, last_name, additional_name, pronouns, head_line)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var newID int64
	err := r.DB.QueryRowContext(
		ctx,
		query,
		newBasicInfo.FirstName,
		newBasicInfo.LastName,
		newBasicInfo.AdditionalName,
		newBasicInfo.Pronouns,
		newBasicInfo.HeadLine,
	).Scan(&newID)

	if err != nil {
		return nil, fmt.Errorf("failed to insert basic info: %w", err)
	}

	return &model.BasicInfo{
		ID:             fmt.Sprintf("%d", newID),
		FirstName:      newBasicInfo.FirstName,
		LastName:       newBasicInfo.LastName,
		AdditionalName: newBasicInfo.AdditionalName,
		Pronouns:       newBasicInfo.Pronouns,
		HeadLine:       newBasicInfo.HeadLine,
	}, nil
}

func (r *mutationResolver) UpdateBasicInfo(ctx context.Context, id *string, data *model.BasicInfoInput) (*model.BasicInfo, error) {
	newBasicInfo := model.BasicInfo{
		ID:             *id,
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		AdditionalName: *data.AdditionalName,
		Pronouns:       data.Pronouns,
		HeadLine:       data.HeadLine,
	}

	query := `
		UPDATE public.basic_info
		SET first_name=$1, last_name=$2, additional_name=$3, pronouns=$4, head_line=$5
		WHERE id=$6;
	`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		newBasicInfo.FirstName,
		newBasicInfo.LastName,
		newBasicInfo.AdditionalName,
		newBasicInfo.Pronouns,
		newBasicInfo.HeadLine,
		*id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert basic info: %w", err)
	}

	return &model.BasicInfo{
		ID:             *id,
		FirstName:      newBasicInfo.FirstName,
		LastName:       newBasicInfo.LastName,
		AdditionalName: newBasicInfo.AdditionalName,
		Pronouns:       newBasicInfo.Pronouns,
		HeadLine:       newBasicInfo.HeadLine,
	}, nil
}

func (r *mutationResolver) DeleteBasicInfo(ctx context.Context, id *string) (*bool, error) {
	query := `DELETE FROM basic_info WHERE id=$1;`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		*id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete basic info: %w", err)
	}

	return nil, nil
}

func (r *queryResolver) BaseInfo(ctx context.Context, id string) (*model.BasicInfo, error) {
	query := `SELECT * FROM basic_info WHERE id=$1;`

	newBasicInfo := &model.BasicInfo{}

	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&newBasicInfo.ID,
		&newBasicInfo.FirstName,
		&newBasicInfo.LastName,
		&newBasicInfo.AdditionalName,
		&newBasicInfo.Pronouns,
		&newBasicInfo.HeadLine,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get basic info: %w", err)
	}

	return newBasicInfo, nil
}

func (r *queryResolver) BaseInfos(ctx context.Context) ([]*model.BasicInfo, error) {
	query := `SELECT * FROM basic_info;`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get basic infos: %w", err)
	}
	defer rows.Close()

	var basicInfos []*model.BasicInfo
	for rows.Next() {
		newBasicInfo := &model.BasicInfo{}
		err = rows.Scan(
			&newBasicInfo.ID,
			&newBasicInfo.FirstName,
			&newBasicInfo.LastName,
			&newBasicInfo.AdditionalName,
			&newBasicInfo.Pronouns,
			&newBasicInfo.HeadLine,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get basic infos: %w", err)
		}
		basicInfos = append(basicInfos, newBasicInfo)
	}

	return basicInfos, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package image

import (
	"context"

	"github.com/disintegration/imaging"

	"github.com/seivanov1986/gocart/internal/repository/image"
)

type ImageCreateIn struct {
	Name      string
	ParentID  int64
	Path      string
	FType     int64
	CreatedAT int64
}

func (u *service) Create(ctx context.Context, in ImageCreateIn) error {
	// GET path from parent

	err := u.hub.Image().Create(ctx, image.ImageCreateInput{
		Name:      in.Name,
		ParentID:  in.ParentID,
		Path:      "/",
		FType:     in.FType,
		CreatedAT: in.CreatedAT,
	})
	if err != nil {
		return err
	}

	return u.makeThumb(ctx, in.Path, in.Name)
}

func (u *service) makeThumb(ctx context.Context, path, name string) error {
	filePath := path + name

	_, err := imaging.Open(filePath, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}

	return nil
}

package google

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"math/rand"
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"google.golang.org/api/customsearch/v1"

	"github.com/johnmanjiro13/dokkoi/command"

	_ "embed"
	_ "image/jpeg"
	_ "image/png"
)

func init() {
	viper.BindEnv("customsearch.api.key", "CUSTOMSEARCH_API_KEY")
	viper.BindEnv("customsearch.engine.id", "CUSTOMSEARCH_ENGINE_ID")

	viper.SetDefault("customsearch.api.key", "")
	viper.SetDefault("customsearch.engine.id", "")
}

const imageNum = 5

var (
	//go:embed assets/Aileron-Regular.otf
	fontBytes []byte
)

type customSearchRepository struct {
	svc      *customsearch.Service
	engineID string
}

func NewCustomSearchRepository(service *customsearch.Service, engineID string) command.CustomSearchRepository {
	return &customSearchRepository{
		svc:      service,
		engineID: engineID,
	}
}

func (r *customSearchRepository) SearchImage(ctx context.Context, query string) (*customsearch.Result, error) {
	images, err := r.searchImage(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(images) <= 0 {
		return nil, command.ErrImageNotFound
	}
	return images[rand.Intn(len(images))], nil
}

func (r *customSearchRepository) LGTM(ctx context.Context, query string) (io.Reader, error) {
	images, err := r.searchImage(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(images) <= 0 {
		return nil, command.ErrImageNotFound
	}

	res, err := http.Get(images[rand.Intn(len(images))].Link)
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	dst, err := drawStringToImage(img, "LGTM")
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, dst, &jpeg.Options{Quality: 100}); err != nil {
		return nil, fmt.Errorf("failed to encode dst image: %w", err)
	}
	return buf, nil
}

func (r *customSearchRepository) searchImage(ctx context.Context, query string) ([]*customsearch.Result, error) {
	search := r.svc.Cse.List().Context(ctx).Cx(r.engineID).
		SearchType("image").
		Num(imageNum).
		Q(query).
		Start(1)
	res, err := search.Do()
	if err != nil {
		return nil, fmt.Errorf("failed to search image: %w", err)
	}
	return res.Items, nil
}

func drawStringToImage(img image.Image, text string) (*image.RGBA, error) {
	dst := image.NewRGBA(img.Bounds())
	draw.Draw(dst, dst.Bounds(), img, image.Point{}, draw.Src)

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    float64(img.Bounds().Dx() / 5),
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create new face: %w", err)
	}

	x := img.Bounds().Dx() / 10 * 2
	y := img.Bounds().Dy() / 5 * 3

	d := &font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(color.RGBA{0, 0, 0, 255}),
		Face: face,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6((x + 1) * 64),
			Y: fixed.Int26_6((y + 1) * 64),
		},
	}
	d.DrawString(text)

	d = &font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6(img.Bounds().Dx() / 10 * 2 * 64),
			Y: fixed.Int26_6(img.Bounds().Dy() / 5 * 3 * 64),
		},
	}
	d.DrawString(text)
	return dst, nil
}

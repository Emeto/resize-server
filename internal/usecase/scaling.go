package usecase

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"resize-server/config"
	"resize-server/internal/entity"
)

type ScalingUseCase struct {
	Scale   entity.Scaling
	Presets map[string]Preset
}

type Preset struct {
	Width  uint
	Height uint
}

func New(presets []config.Preset) *ScalingUseCase {
	prMap := make(map[string]Preset, len(presets))

	for _, preset := range presets {
		prMap[preset.Name] = Preset{
			Width:  preset.Width,
			Height: preset.Height,
		}
	}

	return &ScalingUseCase{
		Scale: entity.Scaling{
			Width:     600,
			Height:    0,
			InterFunc: resize.Lanczos3,
		},
		Presets: prMap,
	}
}

func (uc *ScalingUseCase) GetImageBytes(s entity.Scaling) ([]byte, error) {
	res, err := http.Get(s.Path)
	if err != nil {
		return nil, fmt.Errorf("ScalingUseCase - GetImageBytes - http.Get: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ScalingUseCase - GetImageBytes - io.ReadAll: %w", err)
	}

	err = res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("ScalingUseCase - GetImageBytes - res.Body.Close: %w", err)
	}

	return body, nil
}

func (uc *ScalingUseCase) Decode(r io.Reader, s entity.Scaling) (image.Image, error) {
	if s.Ext == "jpeg" || s.Ext == "jpg" {
		return jpeg.Decode(r)
	}

	if s.Ext == "png" {
		return png.Decode(r)
	}

	if s.Ext == "gif" {
		return gif.Decode(r)
	}

	return nil, fmt.Errorf("ScalingUseCase(Decode) - unsupported file extension")
}

func (uc *ScalingUseCase) Encode(m image.Image, s entity.Scaling) (*bytes.Buffer, error) {
	var buf bytes.Buffer

	if s.Ext == "jpeg" || s.Ext == "jpg" {
		err := jpeg.Encode(&buf, m, nil)
		if err != nil {
			return nil, fmt.Errorf("ScalingUseCase - Encode - jpeg.Encode: %w", err)
		}

		return &buf, nil
	}

	if s.Ext == "png" {
		err := png.Encode(&buf, m)
		if err != nil {
			return nil, fmt.Errorf("ScalingUseCase - Encode - png.Encode: %w", err)
		}

		return &buf, nil
	}

	if s.Ext == "gif" {
		err := gif.Encode(&buf, m, nil)
		if err != nil {
			return nil, fmt.Errorf("ScalingUseCase - Encode - gif.Encode: %w", err)
		}

		return &buf, nil
	}

	return nil, fmt.Errorf("ScalingUseCase(Encode) - unsupported file extension")
}

func (uc *ScalingUseCase) SetSizeFromPreset(p string) bool {
	_, ok := uc.Presets[p]
	if ok {
		uc.Scale.Width = uc.Presets[p].Width
		uc.Scale.Height = uc.Presets[p].Height

		return true
	} else {
		return false
	}
}

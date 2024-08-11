package util

import (
	"bytes"
	"image"
	"image/png"
	"os/exec"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CompressBytes(input []byte) ([]byte, error) {
	cmd := exec.Command("pngquant", "-", "--quality", "60-80", "--speed", "11")
	cmd.Stdin = bytes.NewReader(input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to run pngquant: %s, %s", err, stderr.String())
	}

	return stdout.Bytes(), nil
}

func Compress(input image.Image) ([]byte, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to encode image to PNG: %s", err)
	}

	compressed, err := CompressBytes(buf.Bytes())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to compress image: %s", err)
	}

	return compressed, nil
}

func CompressImage(inputImageBase64 []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(inputImageBase64))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to decode image: %s", err)
	}

	compressed, err := Compress(img)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to compress image: %s", err)
	}

	// Compare sizes to determine whether to return compressed or original data
	if len(compressed) < len(inputImageBase64) {
		return compressed, nil
	}
	return inputImageBase64, nil
}

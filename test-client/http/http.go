package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func Post[ReqT, RspT proto.Message](ctx context.Context, url string, req ReqT, rsp RspT) error {
	m := jsonpb.Marshaler{}
	s, err := m.MarshalToString(req)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(
		ctx,
		"POST",
		url,
		bytes.NewBuffer([]byte(s)),
	)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.Errorf("report http error: %v", response.Status)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = jsonpb.UnmarshalString(string(body), rsp)
	if err != nil {
		return err
	}
	return nil
}

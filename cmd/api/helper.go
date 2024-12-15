package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, r *http.Request, status int, data envelope, header http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range header {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, input interface{}) error {
	maxByte := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&input)
	if err != nil {
		var (
			syntaxError           *json.SyntaxError
			unmarshalTypeError    *json.UnmarshalTypeError
			invalidUnmarshalError *json.InvalidUnmarshalError
		)

		switch {
		case errors.As(err, &syntaxError): // not valid JSON
			return fmt.Errorf("payload contains badly-formed JSON", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("payload contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError): // invalid type
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("incorrect JSON type for field  %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("incorrect JSON  charactor at %d", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF): // empty body
			return fmt.Errorf("body cannot be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{}) // empty anonomous struct to make sure the body only contains a single JSON value
	if !errors.Is(err, io.EOF) {
		return fmt.Errorf("body must only contain a single JSON value")
	}
	return nil
}

package types

import (
	"bytes"
	stdJSON "encoding/json"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/json"
)

type StringUtil struct{}

func (StringUtil) ToRawJSON(prettyJSON string) string {
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	rawJSON, err := m.String("application/json", prettyJSON)
	if err != nil {
		log.Fatal().Msgf("ToRawJSON: %v", err)
	}
	return rawJSON
}

func (StringUtil) ToPrettyJSON(rawJSON []byte) string {
	var prettyJSON bytes.Buffer
	if err := stdJSON.Indent(&prettyJSON, rawJSON, "", "  "); err != nil {
		log.Fatal().Msgf("ToPrettyJSON: %v", err)
	}
	return prettyJSON.String()
}

func (StringUtil) ToRawSQL(prettySQL string) (rawSQL string) {
	rawSQL = prettySQL
	for _, s := range []string{"\t", "\n"} {
		rawSQL = strings.ReplaceAll(rawSQL, s, "")
	}
	return rawSQL
}

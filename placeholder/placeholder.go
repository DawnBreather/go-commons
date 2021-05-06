package placeholder

import (
	"github.com/DawnBreather/go-commons/env_var"
	"github.com/DawnBreather/go-commons/logger"
	"regexp"
	"strings"
)

var _logger = logger.New()

func New(prefix, suffix, core string) *Placeholder {
	return &Placeholder{
		prefix:        prefix,
		suffix:        suffix,
		coreRegexMask: core,
	}
}

type Placeholder struct {
	prefix        string
	suffix        string
	coreRegexMask string
}

func (p *Placeholder) getRegexPattern() *regexp.Regexp {
	regexMask := p.prefix + p.coreRegexMask + p.suffix
	pattern, err := regexp.Compile(regexMask)
	if err != nil {
		_logger.Errorf("Unable to compile regexp mask '%s': %v", regexMask, err)
	}

	return pattern
}

func (p *Placeholder) findAllStringMatches(string string) []string {
	pattern := p.getRegexPattern()
	return pattern.FindAllString(string, -1)
}

func (p *Placeholder) ReplacePlaceholdersWithEnvVars(string string) string {

	var res = string

	for _, match := range p.findAllStringMatches(string) {
		envVarName := strings.ReplaceAll(match, p.prefix, "")
		envVarName = strings.ReplaceAll(envVarName, p.suffix, "")
		envVar := env_var.New(envVarName)
		envVarValue := envVar.Value()
		res = strings.ReplaceAll(res, match, envVarValue)
		if !envVar.IsExist() {
			_logger.Warnf("Environment environment not found: {%s}", envVarName)
		}
	}

	return res
}

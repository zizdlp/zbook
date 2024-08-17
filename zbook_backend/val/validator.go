package val

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/zizdlp/zbook/util"
)

var (
	isValidateUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
)

// 判断是否为有效的时区
func ValidTimeZone(timezone string) error {
	if timezone == "" {
		return fmt.Errorf("timezone cannot be empty")
	}
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Errorf("invalide timezone:%s", timezone)
	}
	return nil
}

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}
func ValidateRepoVisibility(value string) error {
	if value != util.VisibilityChosed && value != util.VisibilityPrivate && value != util.VisibilityPublic && value != util.VisibilitySigned {
		return fmt.Errorf("not valid visibility level")
	}
	return nil
}
func ValidateRepoSideBarTheme(value string) error {
	if value != util.ThemeSideBarFold && value != util.ThemeSideBarUnfold {
		return fmt.Errorf("invalid sidebar theme")
	}
	return nil
}
func ValidateLang(value string) error {
	if value != util.LangEn && value != util.LangZh {
		return fmt.Errorf("invalid language")
	}
	return nil
}
func ValidateRepoThemeColor(value string) error {
	if value != util.ThemeColorViolet && value != util.ThemeColorGreen && value != util.ThemeColorRed && value != util.ThemeColorYellow && value != util.ThemeColorTeal && value != util.ThemeColorSky && value != util.ThemeColorCyan && value != util.ThemeColorPink && value != util.ThemeColorIndigo {
		return fmt.Errorf("invalid theme color")
	}
	return nil
}
func ValidateTitle(value string) error {
	return ValidateString(value, 1, 100)
}
func ValidateID(value int64) error {
	if value <= 0 {
		return fmt.Errorf("ID must greater than 0")
	}
	return nil
}
func ValidatePageSize(value int32) error {
	if value <= 0 {
		return fmt.Errorf("page_size must greater than 0")
	}
	if value > 10 {
		return fmt.Errorf("page_szie must not greater than 10")
	}
	return nil
}
func ValidateInt32ID(value int32) error {
	if value <= 0 {
		return fmt.Errorf("ID must greater than 0")
	}
	return nil
}
func ValidateListUserType(value int64) error {
	if value <= 0 {
		return fmt.Errorf("ID must greater than 0")
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidateUsername(value) {
		return fmt.Errorf("must contain only lower letters,digits, or underscore")
	}
	return nil
}
func ValidateRepoName(repoName string) error {
	if len(repoName) < 2 || len(repoName) > 64 {
		return fmt.Errorf("repository name length is not within the valid range:[2,64]")
	}

	// Characters not allowed in URLs, typically include: '/', '?', ':', '@', '&', '=', '+', '$', ',', '#'
	illegalChars := `/?:@&=+$,#~%`
	if strings.ContainsAny(repoName, illegalChars) {
		return errors.New("repository name contains illegal characters")
	}

	return nil
}
func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}

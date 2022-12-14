package error

import (
	"errors"
	"fmt"

	"github.com/bells307/adv-service/internal/domain/advertisment/service/constant"
)

var (
	ErrNotFound      = errors.New("advertisment not found")
	ErrMaxNameLength = fmt.Errorf("maximum name length %d exceeded", constant.MAX_NAME_LENGTH)
	ErrMaxDescLength = fmt.Errorf("maximum description length %d exceeded", constant.MAX_DESC_LENGTH)
	ErrMaxPhotoCount = fmt.Errorf("maximum advertisment photo count %d exceeded", constant.MAX_PHOTO_COUNT)
)

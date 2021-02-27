package upload

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"

	. "github.com/merq-rodriguez/twitter-go/common/constants"

	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/common/response/errors"
)

func AddFileToStorage(c echo.Context, name string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return BadRequestError(c, "", err)
	}

	defer src.Close()

	ext := strings.Split(file.Filename, ".")[1]
	log.Println("|" + ext + "|")

	if ext != "png" && ext != "jpg" && ext != "jpeg" {
		return NotAcceptableError(c, "Format image invalid", nil)
	}

	dst, err := os.Create(GetUploatPathAvatar() + name + "." + ext)

	if err != nil {
		return BadRequestError(c, "", err)
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return BadRequestError(c, "", err)
	}
	return nil
}

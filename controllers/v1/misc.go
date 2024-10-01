package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/imbee-backend/config"
	"github.com/kalleriakronos24/imbee-backend/dto"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Response{Data: "pongin"})
}

func UploadFileMultiple(c *gin.Context) {
	// Multipart form
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401, Data: nil, Error: "No Form file input found"})
	}
	files := form.File["upload[]"]
	for _, file := range files {
		log.Println(file.Filename)
		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, ".")
		if err != nil {
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func UploadFileSingle(c *gin.Context) {
	// Multipart form
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401, Data: nil, Error: "No Form file input found"})
	}
	workdir, _ := os.Getwd()

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, filepath.Join(workdir, "../files-uploaded", file.Filename))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401, Data: nil, Error: err})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func RestoreDatabase(c *gin.Context) {

	fileName, _ := c.Params.Get("fileName")

	restoreDbKeyPayload, _ := c.GetPostForm("key")

	if fileName == "" {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401,
			Data:  "Ó´¬¬ø⁄ ˇ˙ˆß †´≈† ˆß ƒø® †´ß†ˆ˜© †˙´ †®å˜ß¬å†ø®. ˆƒ å˜¥†˙ˆ˜© ¬øø˚ß ∑®ø˜© ∑ˆ†˙ ˆ†, π¬´åß´ ¬´† µ´ ˚˜ø∑",
			Error: "Ó´¬¬ø⁄ ˇ˙ˆß †´≈† ˆß ƒø® †´ß†ˆ˜© †˙´ †®å˜ß¬å†ø®. ˆƒ å˜¥†˙ˆ˜© ¬øø˚ß ∑®ø˜© ∑ˆ†˙ ˆ†, π¬´åß´ ¬´† µ´ ˚˜ø∑"})
		return
	}

	if restoreDbKeyPayload == "" {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401, Data: nil, Error: nil})
		return
	}

	if restoreDbKeyPayload != config.AppConfig.DBRestoreApiKey {
		c.JSON(http.StatusBadRequest, dto.Response{Code: 401,
			Data:  "Ó´¬¬ø⁄ ˇ˙ˆß †´≈† ˆß ƒø® †´ß†ˆ˜© †˙´ †®å˜ß¬å†ø®. ˆƒ å˜¥†˙ˆ˜© ¬øø˚ß ∑®ø˜© ∑ˆ†˙ ˆ†, π¬´åß´ ¬´† µ´ ˚˜ø∑",
			Error: "Ó´¬¬ø⁄ ˇ˙ˆß †´≈† ˆß ƒø® †´ß†ˆ˜© †˙´ †®å˜ß¬å†ø®. ˆƒ å˜¥†˙ˆ˜© ¬øø˚ß ∑®ø˜© ∑ˆ†˙ ˆ†, π¬´åß´ ¬´† µ´ ˚˜ø∑"})
		return
	}

	workdir, err := os.Getwd()

	if err != nil {
		fmt.Printf("failed to get os workdir : %s", err)
	}

	dbBackupPath := filepath.Join(workdir, fmt.Sprintf("../database-backup/%s.sql", fileName))
	restoreCommand := fmt.Sprintf("docker exec -i %s /bin/bash -c \"PGPASSWORD=%s psql --username %s %s\" < %s.sql", config.AppConfig.DBContainerName, config.AppConfig.DBPassword, config.AppConfig.DBUsername, config.AppConfig.DBDatabase, dbBackupPath)
	cmd, err := exec.Command(restoreCommand).CombinedOutput()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(cmd))
		return
	}

	// Print the output
	fmt.Printf("database backed up : %s \n", fileName)
	c.JSON(http.StatusBadRequest, dto.Response{Code: 201,
		Data:  "ZAÓ´¬¬ø⁄ ˇ˙ˆß †´≈† ˆß ƒø® †´ß†ˆ˜© †˙´ †®å˜ß¬å†ø®. ˆƒ å˜¥†˙ˆ˜© ¬øø˚ß ∑®ø˜© ∑ˆ†˙ ˆ†, π¬´åß´ ¬´† µ´ ˚˜ø∑",
		Error: "XZÓ´¬¬ø⁄ ˇ˙ˆß †´≈† ˆß ƒø® †´ß†ˆ˜© †˙´ †®å˜ß¬å†ø®. ˆƒ å˜¥†˙ˆ˜© ¬øø˚ß ∑®ø˜© ∑ˆ†˙ ˆ†, π¬´åß´ ¬´† µ´ ˚˜ø∑"})

}

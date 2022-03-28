package upload

import (
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"
	"time"

	minio "github.com/GolfPhumrapee/finance/Minio"
	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/GolfPhumrapee/finance/internal/handlers/render"
	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	UpLoads(ctx *fiber.Ctx) error
}

type endpoint struct {
	config *config.Configs
	result *config.ReturnResult
}

// NewEndpoint new migrate endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		config: config.CF,
		result: config.RR,
	}
}

//Randnumber
func RandomString(num int) string {
	var Rand_num int = rand.Intn(num)
	Str := strconv.Itoa(Rand_num)
	return Str
}

// uploadfile
// @Tags Data
// @Summary อัพโหลดไฟล์
// @Description เฉพาะแบบร่างโครงการ สำหรับอัพโหลดไฟล์ หมายเหตุ : บังคับฟิล project_id ทุกครั้ง (ไม่สามารถ required ผ่าน swagger ได้เนื่องจากติดข้อจำกัด)
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)"
// @Param project_id formData string true "ID โครงการ"
// @Param budget_id formData string false "ID Item งบประมาณ"
// @Param fileupload[0] formData file false "input file to upload1"
// @Param fileupload[1] formData file false "input file to upload2"
// @Success 200 {object} config.SwaggerInfoResult
// @Failure 400 {object} config.SwaggerInfoResult
// @Security ApiKeyAuth
// @Router /upload/Upload [post]
func (ep *endpoint) UpLoads(ctx *fiber.Ctx) error {
	// response := &models.ReportSearchPoDateResponse{}
	// err := utils.CheckValiDateValue(ctx, &Uploadfile{})
	// if err != nil {
	// 	return err
	// }
	c := context.New(ctx)
	currentTime := time.Now()
	// DatetimeNow := currentTime.Format("2006-01-02 15:04:05")
	DateNow := currentTime.Format("2006010210")
	YearNow := currentTime.Format("2006")
	//ส่วนลบไฟล์
	// rForm := r.Form.Get("attach_id")

	// if rForm != "" {

	// 	s := strings.Split(rForm, ",")
	// 	updateData := make(map[string]interface{})
	// 	updateData["status"] = 0
	// 	updateData["last_update_at"] = DatetimeNow
	// 	updateData["last_update_user"] = userSession.Id
	// 	result := sql.Database.Model(models.FilesModel{}).Where("attach_id IN(?) and status= ? ", s, 1).Updates(&updateData)
	// 	log.Println(s)
	// 	log.Println(result.RowsAffected) //จำนวนรายการที่มีผล
	// }

	mForm, _ := c.MultipartForm()
	//ส่วนอัพโหลดไฟล์
	var numUpload = int(0) // นับจำนวนที่ Upload
	if ProID := c.FormValue("project_id"); ProID != "" {
		ProID = c.FormValue("project_id")
		if ProID > "0" {
			log.Println(ProID)
			number, _ := strconv.ParseUint(ProID, 10, 32)
			ProID := int(number)
			println(ProID)

			bucketName := config.CF.Minio.MINIO_BUCKET
			// println(bucketName)
			conf := minio.Configuration{
				Host:            config.CF.Minio.MINIO_ENDPOINT,
				AccessKeyID:     config.CF.Minio.MINIO_ACCESSKEY,
				SecretAccessKey: config.CF.Minio.MINIO_SECRETKEY,
				UseSSL:          config.CF.Minio.UseSSL,
			}
			if err := minio.NewConnection(conf); err != nil {
				panic(err)
			}
			client := minio.GetClient()

			for k, v := range mForm.File {
				fileHeader, err := c.FormFile(k)
				if err != nil {
					fmt.Println("inovke FormFile error:", err, v)
					return err
				}
				file, _ := fileHeader.Open()
				objectupload, err := client.OsOpenFile(fileHeader, file)
				if err != nil {
					log.Println(err)
				}
				defer objectupload.Close()
				// defer func() { file.Close() }()
				ext := filepath.Ext(fileHeader.Filename) //นามสกุลไฟล์

				NewName := c.GetEmployeeID() + "-" + DateNow + "-" + RandomString(2) + "-" + RandomString(3) + ext
				FilePath := YearNow + "/" + config.CF.Minio.MINIO_PATH
				FilePush := FilePath + "/" + NewName
				// fun อัพโหลดไฟล์
				// println("bucketName", bucketName)
				// println("FilePush", FilePush)
				// println("file", file)
				// println("fileHeader.Size", fileHeader.Size)
				// println("fileHeader.Header", fileHeader.Header["Content-Type"][0])
				// println("objectupload", objectupload)
				errfile := client.UploadFile1(bucketName, FilePush, fileHeader.Size, objectupload)
				if errfile != nil {
					log.Println("err")
				} else {
					log.Println("success : " + fileHeader.Filename)
					numUpload++
				}
			}
		}
	}
	return render.JSON(ctx, nil)
}

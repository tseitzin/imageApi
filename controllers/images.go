package controllers

import (
	"fmt"
	"imageApi/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/barasher/go-exiftool"
	"github.com/gin-gonic/gin"
)

type CreateImageInput struct {
	ImageFileName    string  `json:"imagefilename" binding:"required"`
	ImageDateTime    string  `json:"imagedatetime" binding:"required"`
	ImageDirLocation string  `json:"imagedirlocation" binding:"required"`
	ImageYear        int     `json:"imageyear"`
	ImageMonth       int     `json:"imagemonth"`
	ImageDay         int     `json:"imageday"`
	ImageWidth       int     `json:"imagewidth"`
	ImageHeight      int     `json:"imageheight"`
	ImageLat         string  `json:"imagelat"`
	ImageLon         string  `json:"imagelon"`
	ImageSize        string  `json:"imagesize"`
	ImageType        string  `json:"imagetype"`
	ImageMegaPixels  float64 `json:"imagemegapixels"`
	ImageFileSize    string  `json:"imagefilesize"`
}

type UpdateImageInput struct {
	ImageFileName    string  `json:"imagefilename"`
	ImageDateTime    string  `json:"imagedatetime"`
	ImageDirLocation string  `json:"imagedirlocation"`
	ImageYear        int     `json:"imageyear"`
	ImageMonth       int     `json:"imagemonth"`
	ImageDay         int     `json:"imageday"`
	ImageWidth       int     `json:"imagewidth"`
	ImageHeight      int     `json:"imageheight"`
	ImageLat         string  `json:"imagelat"`
	ImageLon         string  `json:"imagelon"`
	ImageSize        string  `json:"imagesize"`
	ImageType        string  `json:"imagetype"`
	ImageMegaPixels  float64 `json:"imagemegapixels"`
	ImageFileSize    string  `json:"imagefilesize"`
}

// GetImages responds with the list of all images as JSON.
// GetImages             godoc
// @Summary      Get Images array
// @Description  Responds with the list of all images as JSON.
// @Tags         images
// @Produce      json
// @Success      200  {array}  models.Image
// @Router       /images [get]
func FindImages(c *gin.Context) {
	var image []models.Image

	models.DB.Find(&image).Find(&image)

	c.JSON(http.StatusOK, gin.H{"data": image})
}

// Create new image
// PostBook             godoc
// @Summary      Store a new image
// @Description  Takes a image JSON and store in DB. Return saved JSON.
// @Tags         images
// @Produce      json
// @Param        image body  models.Image  true  "Image Directory"
// @Success      200   {object}  models.Image
// @Router       /images [post]
//
// Add the imagename to the function call coming from the process images function
// that has not been created yet
func CreateImage(c *gin.Context) {
	// Validate input
	var input CreateImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageData, err := processImage(input)

	if err != nil {
		log.Fatal(err)
	}

	// Create image
	image := models.Image{
		ImageFileName:    imageData.ImageFileName,
		ImageDirLocation: imageData.ImageDirLocation,
		ImageDateTime:    imageData.ImageDateTime,
		ImageYear:        imageData.ImageYear,
		ImageMonth:       imageData.ImageMonth,
		ImageDay:         imageData.ImageDay,
		ImageWidth:       imageData.ImageWidth,
		ImageHeight:      imageData.ImageHeight,
		ImageLat:         input.ImageLat,
		ImageLon:         input.ImageLon,
		ImageSize:        imageData.ImageSize,
		ImageType:        imageData.ImageType,
		ImageMegaPixels:  imageData.ImageMegaPixels,
		ImageFileSize:    imageData.ImageFileSize}

	models.DB.Create(&image)

	c.JSON(http.StatusOK, gin.H{"data": image})
}

// Delete an image
// DeleteImage                godoc
// @Summary      Delete single image by image_id
// @Description  Delete the image whose ID value matches the image_id.
// @Tags         images
// @Produce      json
// @Param        image_id  path      string  true  "delete image by image_id"
// @Success      200  {object}  models.Image
// @Router       /images/{image_id} [delete]
func DeleteImage(c *gin.Context) {
	// Get model if exist
	var image models.Image
	if err := models.DB.Where("image_id = ?", c.Param("image_id")).First(&image).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Where("image_id = ?", c.Param("image_id")).Delete(&image)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// Find an image
// FindBook                godoc
// @Summary      Get single image by ImageID
// @Description  Returns the image whose ID value matches the ImageID.
// @Tags         images
// @Produce      json
// @Param        image_id  path      string  true  "search image by ImageID"
// @Success      200  {object}  models.Image
// @Router       /images/{image_id} [get]
func FindImage(c *gin.Context) { // Get model if exist
	var image models.Image

	if err := models.DB.Where("image_id = ?", c.Param("image_id")).First(&image).Find(&image).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": image})
}

// Update an image
// UpdateImage                godoc
// @Summary      Update single image by image_id
// @Description  Updates the image whose ID value matches the image_id.
// @Tags         images
// @Produce      json
// @Param        image_id  path      int  true  "update image by image_id"
// @Param        image  body      models.Image  true  "Image JSON"
// @Success      200  {object}  models.Image
// @Router       /images/{image_id} [patch]
func UpdateImage(c *gin.Context) {
	// Get model if exist
	var image models.Image
	if err := models.DB.Where("image_id = ?", c.Param("image_id")).First(&image).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateImageInput := models.Image{
		ImageFileName:    input.ImageFileName,
		ImageDirLocation: input.ImageDirLocation,
		ImageDateTime:    input.ImageDateTime,
		ImageYear:        input.ImageYear,
		ImageMonth:       input.ImageMonth,
		ImageDay:         input.ImageDay,
		ImageWidth:       input.ImageWidth,
		ImageHeight:      input.ImageHeight,
		ImageLat:         input.ImageLat,
		ImageLon:         input.ImageLon,
		ImageSize:        input.ImageSize,
		ImageType:        input.ImageType,
		ImageMegaPixels:  input.ImageMegaPixels,
		ImageFileSize:    input.ImageFileSize}

	models.DB.Model(&image).Where("image_id = ?", c.Param("image_id")).Updates(&UpdateImageInput)

	c.JSON(http.StatusOK, gin.H{"data": UpdateImageInput})
}

func testImageFile(file *os.File) bool {

	resp, err := os.Open(file.Name())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Close()

	// Read the response body as a byte slice
	bytes, err := io.ReadAll(resp)
	if err != nil {
		log.Fatal(err)
	}

	mimeType := http.DetectContentType(bytes)
	fmt.Println(mimeType) // image/png

	test := strings.Split(mimeType, "/")

	fmt.Println(test)

	if test[0] != "image" {
		fmt.Println("**************************NOT AN IMAGE FILE************************")
		return false
	} else {
		return true
	}
}

func processImage(input CreateImageInput) (CreateImageInput, error) {

	fileName := strings.Split(input.ImageDirLocation, `\`)
	input.ImageFileName = fileName[len(fileName)-1]

	// Open the image file
	file, err := os.Open(input.ImageDirLocation)
	if err != nil {
		return input, fmt.Errorf("error opening file %s: %w", input.ImageDirLocation, err)
	}
	defer file.Close()

	imageFile := testImageFile(file)
	if !imageFile {
		return input, fmt.Errorf("Not an image file for file %s", input.ImageDirLocation)
	}

	et, err := exiftool.NewExiftool()
	if err != nil {
		return input, fmt.Errorf("error decoding EXIF data for file %s: %w", input.ImageDirLocation, err)
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(file.Name())

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Errorf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			fmt.Printf("[%v] %v\n", k, v)

			switch {
			case k == "CreateDate":
				input.ImageDateTime = v.(string)
			case k == "FileType":
				input.ImageType = v.(string)
			case k == "ImageWidth":
				input.ImageWidth = int(v.(float64))
			case k == "ImageHeight":
				input.ImageHeight = int(v.(float64))
			case k == "ImageSize":
				input.ImageSize = v.(string)
			case k == "Megapixels":
				input.ImageMegaPixels = v.(float64)
			case k == "FileSize":
				input.ImageFileSize = v.(string)
			}

		}
	}

	dateTimeSplit := strings.Split(input.ImageDateTime, " ")
	dateSplit := strings.Split(dateTimeSplit[0], ":")

	input.ImageYear, _ = strconv.Atoi(dateSplit[0])
	input.ImageMonth, _ = strconv.Atoi(dateSplit[1])
	input.ImageDay, _ = strconv.Atoi(dateSplit[2])

	return input, nil
}

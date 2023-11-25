package controllers

import (
	"imageApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateImageInput struct {
	ImageFileName    string `json:"imagefilename" binding:"required"`
	ImageDateTime    string `json:"imagedatetime" binding:"required"`
	ImageDirLocation string `json:"imagedirlocation" binding:"required"`
	ImageYear        int    `json:"imageyear"`
	ImageMonth       int    `json:"imagemonth"`
	ImageDay         int    `json:"imageday"`
	ImageWidth       int    `json:"imagewidth"`
	ImageHeight      int    `json:"imageheight"`
	ImageLat         string `json:"imagelat"`
	ImageLon         string `json:"imagelon"`
	ImageSize        string `json:"imagesize"`
	ImageType        string `json:"imagetype"`
}

type UpdateImageInput struct {
	ImageFileName    string `json:"imagefilename"`
	ImageDateTime    string `json:"imagedatetime"`
	ImageDirLocation string `json:"imagedirlocation"`
	ImageYear        int    `json:"imageyear"`
	ImageMonth       int    `json:"imagemonth"`
	ImageDay         int    `json:"imageday"`
	ImageWidth       int    `json:"imagewidth"`
	ImageHeight      int    `json:"imageheight"`
	ImageLat         string `json:"imagelat"`
	ImageLon         string `json:"imagelon"`
	ImageSize        string `json:"imagesize"`
	ImageType        string `json:"imagetype"`
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
// @Param        image  body      models.Image  true  "Image JSON"
// @Success      200   {object}  models.Image
// @Router       /images [post]
func CreateImage(c *gin.Context) {
	// Validate input
	var input CreateImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create image
	image := models.Image{
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
		ImageType:        input.ImageType}

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
		ImageType:        input.ImageType}

	models.DB.Model(&image).Where("image_id = ?", c.Param("image_id")).Updates(&UpdateImageInput)

	c.JSON(http.StatusOK, gin.H{"data": UpdateImageInput})
}

package models

type Image struct {
	ImageID          int    `gorm:"primaryKey"`
	ImageFileName    string `gorm:"unique"`
	ImageDateTime    string
	ImageYear        int
	ImageMonth       int
	ImageDay         int
	ImageDirLocation string
	ImageWidth       int
	ImageHeight      int
	ImageLat         string
	ImageLon         string
	ImageSize        string
	ImageType        string
}

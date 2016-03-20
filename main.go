package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aalpern/luminosity"
)

func main() {
	c, err := luminosity.OpenCatalog(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	lenses, _ := c.GetLenses()
	cameras, _ := c.GetCameras()
	distByDate, _ := c.GetPhotoCountsByDate()
	distByLens, _ := c.GetLensDistribution()
	distByFocalLength, _ := c.GetFocalLengthDistribution()
	distByCamera, _ := c.GetCameraDistribution()
	distByAperture, _ := c.GetApertureDistribution()
	distByExposureTime, _ := c.GetExposureTimeDistribution()

	data := map[string]interface{}{
		"cameras": cameras,
		"lenses":  lenses,
		"distribution": map[string]interface{}{
			"by_date":          distByDate,
			"by_lens":          distByLens,
			"by_focal_length":  distByFocalLength,
			"by_camera":        distByCamera,
			"by_aperture":      distByAperture,
			"by_exposure_time": distByExposureTime,
		},
	}

	print(data)
}

func print(data interface{}) {
	js, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(js))
}

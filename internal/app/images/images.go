package images

import (
	"bytes"
	"image/jpeg"
	"strconv"

	gim "github.com/ozankasikci/go-image-merge"
)

func NumToImage(dir string, num int) ([]byte, error) {
	strNum := strconv.Itoa(num)

	grids := []*gim.Grid{}

	for _, v := range strNum {
		grids = append(grids, &gim.Grid{
			ImageFilePath: dir + "/" + string(v) + ".jpg",
		})
	}

	rgba, err := gim.New(
		grids, len(strNum), 1,
	).Merge()

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, rgba, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

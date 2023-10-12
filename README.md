# Convert-To-BMP

App to convert a JPG/PNG/GIF image to a BMP file which then can be used with [Waveshare's ACeP 7-Color E-Paper Wood Frame (PhotoPainter)](https://www.waveshare.com/photopainter.htm)

    go install github.com/frifox/convert-to-bmp@latest
    convert-to-bmp -i image.jpg
    cp image-out.bmp /mnt/my-sd-card/pic

## Usage
    convert-to-bmp [options]

## Options
    -i string
        Input filename
    -d float
        Dithering strength 0.0 to 1.0 (default 1.0)
    -b float
        Adjust brightness -100 to 100. 0 for none (default 0)
    -c float
        Adjust contrast -100 to 100. 0 for none (default 0)


## Example
    
    # make image 20% brighter, dither at 0.8 strength
    convert-to-bmp -i 800x480.jpg -d 0.8 -b 20
    
    # copy generated image to `pic` folder on sdcard
    mkdir /mnt/my-sd-card/pic
    cp 800x480-out.bmp /mnt/my-sd-card/pic
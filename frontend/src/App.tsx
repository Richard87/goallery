import './App.css'
import PhotoAlbum from "react-photo-album";
import Lightbox from "yet-another-react-lightbox";
import {useState} from "react";

import "yet-another-react-lightbox/styles.css";

// import optional lightbox plugins
import Fullscreen from "yet-another-react-lightbox/plugins/fullscreen";
import Slideshow from "yet-another-react-lightbox/plugins/slideshow";
import Thumbnails from "yet-another-react-lightbox/plugins/thumbnails";
import Zoom from "yet-another-react-lightbox/plugins/zoom";
import "yet-another-react-lightbox/plugins/thumbnails.css";

const photos = [
  {
    fallback: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFEAIAAABSnW3xAAAAq0lEQVR4nACbAGT/BH+wew9ynww8BpQFfQf/9k7i7/47EnYooxFBDfkDawQBmwJhBMf2yO/N5xMywBJe8uHazvvN8kf2n/nxD0oE95L5V/u/EdMBOP+CArntEfFoCJ4VL/yy7Cr4dvrvAv9k/Z/7N/jH/hgBTemE60rtsfH1+yoADv7+/v7+/gMPkBEXCj/K/cnHyIQS9Q3JAQirOMah2zj5Yfth/N4BAAD//ynfVBBbCUUXAAAAAElFTkSuQmCC",
    src: "/20210720_0001.JPG?1",
    width: 6000,
    height: 4000,
  },
  {
    fallback: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFEAIAAABSnW3xAAAAq0lEQVR4nACbAGT/BH+wew9ynww8BpQFfQf/9k7i7/47EnYooxFBDfkDawQBmwJhBMf2yO/N5xMywBJe8uHazvvN8kf2n/nxD0oE95L5V/u/EdMBOP+CArntEfFoCJ4VL/yy7Cr4dvrvAv9k/Z/7N/jH/hgBTemE60rtsfH1+yoADv7+/v7+/gMPkBEXCj/K/cnHyIQS9Q3JAQirOMah2zj5Yfth/N4BAAD//ynfVBBbCUUXAAAAAElFTkSuQmCC",
    src: "/20210720_0001.JPG?12",
    width: 6000,
    height: 4000,
  },
  {
    fallback: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFEAIAAABSnW3xAAAAq0lEQVR4nACbAGT/BH+wew9ynww8BpQFfQf/9k7i7/47EnYooxFBDfkDawQBmwJhBMf2yO/N5xMywBJe8uHazvvN8kf2n/nxD0oE95L5V/u/EdMBOP+CArntEfFoCJ4VL/yy7Cr4dvrvAv9k/Z/7N/jH/hgBTemE60rtsfH1+yoADv7+/v7+/gMPkBEXCj/K/cnHyIQS9Q3JAQirOMah2zj5Yfth/N4BAAD//ynfVBBbCUUXAAAAAElFTkSuQmCC",
    src: "/20210720_0001.JPG?13",
    width: 6000,
    height: 4000,
  },
  {
    fallback: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFEAIAAABSnW3xAAAAq0lEQVR4nACbAGT/BH+wew9ynww8BpQFfQf/9k7i7/47EnYooxFBDfkDawQBmwJhBMf2yO/N5xMywBJe8uHazvvN8kf2n/nxD0oE95L5V/u/EdMBOP+CArntEfFoCJ4VL/yy7Cr4dvrvAv9k/Z/7N/jH/hgBTemE60rtsfH1+yoADv7+/v7+/gMPkBEXCj/K/cnHyIQS9Q3JAQirOMah2zj5Yfth/N4BAAD//ynfVBBbCUUXAAAAAElFTkSuQmCC",
    src: "/20210720_0001.JPG?14",
    width: 6000,
    height: 4000,
  },
  {
    fallback: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFEAIAAABSnW3xAAAAq0lEQVR4nACbAGT/BH+wew9ynww8BpQFfQf/9k7i7/47EnYooxFBDfkDawQBmwJhBMf2yO/N5xMywBJe8uHazvvN8kf2n/nxD0oE95L5V/u/EdMBOP+CArntEfFoCJ4VL/yy7Cr4dvrvAv9k/Z/7N/jH/hgBTemE60rtsfH1+yoADv7+/v7+/gMPkBEXCj/K/cnHyIQS9Q3JAQirOMah2zj5Yfth/N4BAAD//ynfVBBbCUUXAAAAAElFTkSuQmCC",
    src: "/20210720_0001.JPG?15",
    width: 6000,
    height: 4000,
  },
]

function App() {

  const [index, setIndex] = useState<number>(-1)

  return (<>
        <h1>Hello world</h1>
      <PhotoAlbum
          layout="rows"
          targetRowHeight={150}
          photos={photos}
          onClick={({ index }) => setIndex(index)}
          renderPhoto={({ imageProps: {style, ...imageProps} , photo }) => (
              <img
                  style={{
                    ...style,
                    backgroundImage: `url('${photo.fallback}')`,
                    backgroundRepeat: "no-repeat",
                    backgroundSize: "cover",
                    backgroundPosition: "center",
                  }}
                  loading="lazy"
                  {...imageProps}
              />
          )}
      />

        <Lightbox
            slides={photos}
            open={index >= 0}
            index={index}
            close={() => setIndex(-1)}
            // enable optional lightbox plugins
            plugins={[Fullscreen, Slideshow, Thumbnails, Zoom]}
        />
      </>
  )
}

export default App

import type {LinksFunction, MetaFunction } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { useState } from "react";
import PhotoAlbum from "react-photo-album";
import Lightbox from "yet-another-react-lightbox";
import { Fullscreen, Slideshow, Thumbnails, Zoom } from "yet-another-react-lightbox/plugins";
import "../../node_modules/yet-another-react-lightbox/dist/plugins/thumbnails/thumbnails.css";
import "../../node_modules/yet-another-react-lightbox/dist/styles.css";
import {Configuration, Image, ImagesApi} from "../../api";

export const meta: MetaFunction = () => {
  return [
    { title: "Goallery" },
    { name: "description", content: "Welcome to Goallery!" },
  ];
};

const api = new ImagesApi(new Configuration({
    basePath: "http://localhost:8000/api/v1",
    headers: {
        authorization: "Bearer hello-world"
    }
}))

export const loader = async () => {
    const photos = await api.getImages()
    return {photos}
}

export default function Index() {
    const {photos} = useLoaderData<{photos: Image[]}>()
    const [index, setIndex] = useState(-1)
  return (
    <div style={{ fontFamily: "system-ui, sans-serif", lineHeight: "1.8" }}>
      <h1>Welcome to Gaollery</h1>
      <PhotoAlbum
          layout="rows"
          targetRowHeight={150}
          photos={photos}
          onClick={({ index }) => setIndex(index)}
          renderPhoto={({ imageProps: {style, ...imageProps} , photo }) => (
              <img
                  style={{
                    ...style,
                    backgroundImage: `url('${photo.features["plugin.blurryimage"]}')`,
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
    </div>
  );
}

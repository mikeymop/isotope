import Nav from "@/components/ui/navbar";
import Album from "@/components/Album/Album";
// import GalleryListItem from "../../components/GalleryListItem/GalleryListItem";

import type { GalleryMeta } from "../../models/gallery";
import { listenNowAlbums } from "@/components/Album/data";

const makeGalleries = (): GalleryMeta[] =>
  Array.from({ length: 10 }, (_, i) => {
    return {
      id: i,
      title: `Test Gallery ${i}`,
      size: Math.floor(Math.random() * 100) + 1,
    };
  });

const galleries = makeGalleries();

/** The Gallery page, displays all created client galleries. */
const Gallery = () => {
  return (
    <>
      <div className="flex row m-8">
        {listenNowAlbums.map((album, index) => (
          <Album
            key={album.name}
            album={album}
            data={galleries[index]}
            className="w-[150px] m-4"
            aspectRatio="square"
            width={150}
            height={150}
          />
        ))}
        {/* {makeGalleries().map((gallery) => (
          <div className="column">
            <GalleryListItem key={gallery.id} data={gallery} />
          </div>
        ))} */}
      </div>
    </>
  );
};

export default Gallery;

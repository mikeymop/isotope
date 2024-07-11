import { Avatar, Flex, Text } from "@radix-ui/themes";
import { GalleryMeta } from "../../models/gallery";
import Card from "@/components/ui/card";
// import Card from "../Card/Card";

interface GalleryListItemProps {
  id: number;
  data: GalleryMeta;
}

/**
 * Represents a gallery within the main client gallery page.
 */
const GalleryListItem = ({ data }: GalleryListItemProps) => {
  return (
    <Card>
      <Flex gap="3" align="center">
        <Avatar
          size="3"
          src="https://images.unsplash.com/photo-1607346256330-dee7af15f7c5?&w=64&h=64&dpr=2&q=70&crop=focalpoint&fp-x=0.67&fp-y=0.5&fp-z=1.4&fit=crop"
          radius="full"
          fallback="T"
        />
        <div>
          <Text>{data.title}</Text>
          <p>{JSON.stringify(data)}</p>
        </div>
      </Flex>
    </Card>
  );
};

export default GalleryListItem;

import { User } from "@/domain/user";
import { ListItem, UnorderedList } from "@chakra-ui/react";

import React from "react";

type Props = {
  list: User[];
};

const GeneralList = ({ list }: Props) => {
  return (
    <UnorderedList>
      {list.map((item, i) => (
        <ListItem key={i}>
          {item.id}: {item.name}
        </ListItem>
      ))}
    </UnorderedList>
  );
};

export default GeneralList;

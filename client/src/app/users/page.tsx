"use client";

import { Container } from "@chakra-ui/react";
import GeneralList from "@/components/List.tsx/General";
import { UserApi } from "../../repositories/userApi";

export default async function users() {
  const client = new UserApi();
  const users = await client.getAll();

  return (
    <Container>
      <GeneralList list={users}></GeneralList>
    </Container>
  );
}

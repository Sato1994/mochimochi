"use client";
import { HttpClient } from "@/app/repositories/httpClient";
import { Button as _Button } from "@chakra-ui/react";

import React from "react";

type Props = {
  handleClick: () => Promise<void>;
  buttonText: string;
};


const Button = (
  { handleClick, buttonText }: Props
) => {
  
  return (
    <_Button colorScheme="blue" size="md" onClick={handleClick}>
      { buttonText }
    </_Button>
  );
};

export default Button;
"use client";
import { HttpClient } from "@/app/repositories/httpClient";
import { Button as _Button } from "@chakra-ui/react";

import React from "react";

const Button = () => {
  const handleButtonClick = async () => {
    const dom = document.documentElement.outerHTML;
    const client = new HttpClient()
    console.log(dom);
    console.log(await client.get("/todos/1"))
  };
  
  return (
    <_Button colorScheme="blue" size="md" onClick={handleButtonClick}>
      ドむ取得するよ
    </_Button>
  );
};

export default Button;
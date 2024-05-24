"use client";
// chakra uiを使用

import { Button as _Button } from "@chakra-ui/react";

import React from "react";

const Button = () => {
  const handleButtonClick = () => {
    const dom = document.documentElement.outerHTML;
    console.log(dom);
  };
  
  return (
    <_Button colorScheme="blue" size="md" onClick={handleButtonClick}>
      get DOM
    </_Button>
  );
};

export default Button;
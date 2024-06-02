import React from "react";
import { Button as _Button, ButtonProps } from "@chakra-ui/react";

const Button: React.FC<React.PropsWithChildren<ButtonProps>> = (props) => {
  return (
    <_Button colorScheme="blue" size="md" {...props}>
      {props.children}
    </_Button>
  );
};

export default Button;

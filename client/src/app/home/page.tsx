import Button from "@/components/Button";
import { Container } from "@chakra-ui/react";
import { HttpClient } from "../repositories/httpClient";

export default function home() {

   const handleButtonClick = async () => {
    const dom = document.documentElement.outerHTML;
    const client = new HttpClient()
    console.log(dom);
    console.log(await client.get("/todos/1"))
  };

  return (
    <Container>
      <Button handleClick={ handleButtonClick } buttonText="DOMの取得" />
    </Container>
  );
}

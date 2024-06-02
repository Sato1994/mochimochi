import Button from "@/components/Button";
import { Box, Container, Image } from "@chakra-ui/react";
import { useState } from "react";

export default function testActions() {
  const [screenShotPath, setScreenShotPath] = useState<string>("otameshi.jpg");

  async function testAction() {
    await autoBrowserTest();
  }

  async function autoBrowserTest() {
    // https://pptr.dev/
    // Puppeteerを使用する

    const browser = await  puppeteer.launch({ headless: true });

    const page = await browser.newPage();
    // await page.goto("https://github.com/Sato1994");
    const nowTimestamp = new Date().getTime();

    const r = await page.screenshot({ path: `${nowTimestamp}.png` });
    console.log({r});
    await browser.close();
    setScreenShotPath(`${nowTimestamp}.png`);
  }

  return (
    <Container>
      <Button onClick={testAction}>Click me</Button>

      <Box>
        {screenShotPath && <Image src={screenShotPath} alt="screenshot" />}
      </Box>
    </Container>
  );
}

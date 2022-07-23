import { Box, Flex } from "@chakra-ui/react";
import Link from "next/link";
import { FC } from "react";

export const Header: FC = () => {
  return (
    <Flex
      as="header"
      w="full"
      h="60px"
      verticalAlign={"center"}
      bg={"cyan.100"}
      alignItems={"center"}
      pl={"4"}
      position={"fixed"}
    >
      <Link href="/">
        <Box as="a">My Task App</Box>
      </Link>
    </Flex>
  );
};

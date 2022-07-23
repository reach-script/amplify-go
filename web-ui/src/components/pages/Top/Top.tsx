import { Box, chakra } from "@chakra-ui/react";
import Link from "next/link";
import { FC } from "react";

export const Top: FC = () => {
  return (
    <Box as="main">
      <chakra.h1>Top Page</chakra.h1>
      <Link href="/tasks">
        <chakra.a>Go Tasks Page</chakra.a>
      </Link>
    </Box>
  );
};

import { Container, Flex, Spinner } from "@chakra-ui/react";
import { FC } from "react";

export const PageLoading: FC = () => {
  return (
    <Container w="container.xl" h="full">
      <Flex
        w="full"
        h="full"
        justifyContent={"center"}
        alignItems={"center"}
        bg={"gray.50"}
      >
        <Spinner />
      </Flex>
    </Container>
  );
};

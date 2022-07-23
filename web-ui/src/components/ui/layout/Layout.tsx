import { Box, Container } from "@chakra-ui/react";
import { FC, ReactNode } from "react";

type Props = {
  header?: ReactNode;
  children: ReactNode;
};

export const Layout: FC<Props> = (props) => {
  const { children, header } = props;
  return (
    <Box h="full">
      {header}
      <Container w="container.xl" h="full" pt="60px">
        {children}
      </Container>
    </Box>
  );
};

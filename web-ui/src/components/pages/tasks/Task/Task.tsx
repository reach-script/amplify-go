import { Button, chakra } from "@chakra-ui/react";
import { useRouter } from "next/router";
import { FC } from "react";

export const Task: FC = () => {
  const router = useRouter();
  return (
    <chakra.main>
      <chakra.h2>task detail page : {router.query.id}</chakra.h2>
      <Button onClick={router.back}>back</Button>
    </chakra.main>
  );
};

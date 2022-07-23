import { Box, Button, chakra } from "@chakra-ui/react";
import Link from "next/link";
import { FC } from "react";
import useSWR, { Key } from "swr";

const url: Key = "/todos";
type Task = {
  userId: number;
  id: number;
  title: string;
  completed: boolean;
};

export const Tasks: FC = () => {
  const { data } = useSWR<Task[]>(url);

  if (!data) {
    return <Box>no data</Box>;
  }
  return (
    <Box as="main" minH={"full"} bg={"gray.50"}>
      <Link href="/tasks/create">
        <Button as="a">Create</Button>
      </Link>
      <h1>Tasks Page</h1>
      {data.map((item) => (
        <chakra.li key={item.id}>
          {item.title}
          <Link href={`/tasks/${item.id}`}>
            <Button as="a" size="xs">
              detail
            </Button>
          </Link>
        </chakra.li>
      ))}
    </Box>
  );
};

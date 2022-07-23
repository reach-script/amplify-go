import { NextPage } from "next";
import { Task } from "src/components/pages/tasks/Task/Task";
import { Header } from "src/components/ui/layout/Header";
import { Layout } from "src/components/ui/layout/Layout";

export const TaskPage: NextPage = () => {
  return (
    <Layout header={<Header />}>
      <Task />
    </Layout>
  );
};

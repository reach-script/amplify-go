import { NextPage } from "next";
import { TaskCreate } from "src/components/pages/tasks/TaskCreate/TaskCreate";
import { Header } from "src/components/ui/layout/Header";
import { Layout } from "src/components/ui/layout/Layout";

export const TaskCreatePage: NextPage = () => {
  return (
    <Layout header={<Header />}>
      <TaskCreate />
    </Layout>
  );
};

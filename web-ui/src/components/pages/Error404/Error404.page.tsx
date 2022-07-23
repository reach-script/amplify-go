import { NextPage } from "next";
import { Error404 } from "src/components/pages/Error404/Error404";
import { Header } from "src/components/ui/layout/Header";
import { Layout } from "src/components/ui/layout/Layout";

export const Error404Page: NextPage = () => {
  return (
    <Layout header={<Header />}>
      <Error404 />
    </Layout>
  );
};

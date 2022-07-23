import { NextPage } from "next";
import { Top } from "src/components/pages/Top/Top";
import { Header } from "src/components/ui/layout/Header";
import { Layout } from "src/components/ui/layout/Layout";

export const TopPage: NextPage = () => {
  return (
    <Layout header={<Header />}>
      <Top />
    </Layout>
  );
};

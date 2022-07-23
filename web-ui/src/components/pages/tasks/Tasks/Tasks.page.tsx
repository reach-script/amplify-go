import { NextPage } from "next";
import { Component, FC, ReactNode, Suspense } from "react";
import { Tasks } from "src/components/pages/tasks/Tasks/Tasks";
import { Layout } from "src/components/ui/layout/Layout";
import { Header } from "src/components/ui/layout/Header";
import { PageLoading } from "src/components/ui/loading/PageLoading";
import { TasksErrorFallback } from "src/components/pages/tasks/Tasks/TasksErrorFallback";
import { HttpErrorBoundary } from "src/components/functional/HttpErrorBoundary";

export const TasksPage: NextPage = () => {
  return (
    <Layout header={<Header />}>
      <HttpErrorBoundary FallbackComponent={TasksErrorFallback}>
        <Suspense fallback={<PageLoading />}>
          <Tasks />
        </Suspense>
      </HttpErrorBoundary>
    </Layout>
  );
};

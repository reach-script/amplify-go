import "src/styles/global.css";
import type { AppProps } from "next/app";
import { Amplify } from "aws-amplify";
import { ChakraProvider } from "@chakra-ui/react";

import { Authenticator } from "@aws-amplify/ui-react";
import "@aws-amplify/ui-react/styles.css";

import awsExports from "../aws-exports";
import { SWRConfig } from "src/components/functional/SWRConfig";
import { FC } from "react";
import { ErrorBoundary } from "src/components/functional/ErrorBoundary";
Amplify.configure({
  ...awsExports,
});

const MyApp: FC<AppProps> = ({ Component, pageProps }) => {
  return (
    <ChakraProvider>
      <SWRConfig>
        <ErrorBoundary FallbackComponent={() => <>Error</>}>
          <Authenticator>
            <Component {...pageProps} />
          </Authenticator>
        </ErrorBoundary>
      </SWRConfig>
    </ChakraProvider>
  );
};

export default MyApp;

import "../styles/globals.css";
import type { AppProps } from "next/app";
import { Amplify } from "aws-amplify";
import { Auth } from "@aws-amplify/auth";

import { Authenticator } from "@aws-amplify/ui-react";
import "@aws-amplify/ui-react/styles.css";

import awsExports from "../aws-exports";
import { UseAuthenticator } from "@aws-amplify/ui-react/dist/types/components/Authenticator/hooks/useAuthenticator";
Amplify.configure(awsExports);

// https://cognito-idp.ap-northeast-1.amazonaws.com/ap-northeast-1_Oxzc2wtHu/.well-known/jwks.json

function MyApp({ Component, pageProps }: AppProps) {
  const onClick = async () => {
    const jwt = await (await Auth.currentSession())
      .getAccessToken()
      .getJwtToken();
    console.log(jwt);
  };

  return (
    <Authenticator>
      {({
        signOut,
        user,
      }: {
        signOut?: UseAuthenticator["signOut"];
        user?: any;
      }) => (
        <main>
          <h1>Hello {user.username}</h1>
          <button onClick={signOut}>Sign out</button>
          <button onClick={onClick}>onClick</button>
          <Component {...pageProps} />
        </main>
      )}
    </Authenticator>
  );
}

export default MyApp;

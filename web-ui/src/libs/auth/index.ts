import { Auth } from "aws-amplify";

export const getAccessToken = async () => {
  const currentSession = await Auth.currentSession();
  const accessToken = currentSession.getAccessToken();
  const jwt = accessToken.getJwtToken();
  return jwt;
};

export const logout = async () => {
  await Auth.signOut();
};

import { API, Auth } from "aws-amplify";
import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";

const Home: NextPage = () => {
  const onClick = async () => {
    const currentSession = await Auth.currentSession();
    console.log("currentSession", currentSession);
    const idToken = currentSession.getAccessToken();
    console.log("idToken", idToken);

    const jwtToken = idToken.getJwtToken();
    console.log("jwtToken", jwtToken);

    const myInit = {
      headers: {
        Authorization: `Bearer ${jwtToken}`,
      },
      body: {
        text: "test",
      },
    };
    try {
      const response = await API.get("tasks", "/items", myInit);
      console.log(response);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className={styles.container}>
      <button onClick={onClick}>API Call</button>
    </div>
  );
};

export default Home;

import { useEffect, useState } from "react";
import { getMe } from "../apiService";

const Home = () => {
  const [me, setMe] = useState<string>("");

  useEffect(() => {
    async function fetchMessage() {
      const username = await getMe();
      setMe(username === undefined ? "" : username);
    }
    fetchMessage();
  }, []);
  return <h1>{`Home: Welcome to SimpleTM ${me}`}</h1>;
};

export default Home;

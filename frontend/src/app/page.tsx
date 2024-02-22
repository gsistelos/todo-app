"use client";

import styles from "./page.module.css";
import Header from "./components/Header";
import UserActions from "./components/UserActions";
import { UserContext, fetchUser } from "./contexts/UserContext";
import { User } from "./types";
import { useEffect, useState } from "react";

const Home = () => {
  const [loaded, setLoaded] = useState(false);
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    fetchUser().then((user) => {
      setUser(user);
      setLoaded(true);
    });
  }, []);

  const value = { user, setUser };

  return (
    <main className={styles.main}>
      {loaded ? (
        <UserContext.Provider value={value}>
          <Header
            title="Home"
            subtitle={
              user ? `Welcome, ${user.username}!` : "You are not logged in."
            }
          />
          <UserActions />
        </UserContext.Provider>
      ) : (
        <span className={styles.loading}>Loading...</span>
      )}
    </main>
  );
};

export default Home;

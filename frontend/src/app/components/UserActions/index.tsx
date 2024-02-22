"use client";

import styles from "./UserActions.module.css";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

const UserActions = () => {
  const [loggedIn, setLoggedIn] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setLoggedIn(true);
    }
  }, []);

  const router = useRouter();

  return loggedIn ? (
    <div className={styles.container}>
      <button className={styles.button} onClick={() => router.push("/profile")}>
        Profile
      </button>
      <button
        className={styles.button}
        onClick={() => {
          localStorage.removeItem("token");
          setLoggedIn(false);
        }}
      >
        Logout
      </button>
    </div>
  ) : (
    <div className={styles.container}>
      <button
        className={styles.button}
        onClick={() => router.push("/register")}
      >
        Register
      </button>
      <button className={styles.button} onClick={() => router.push("/login")}>
        Login
      </button>
    </div>
  );
};

export default UserActions;

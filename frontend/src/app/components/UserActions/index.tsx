"use client";

import styles from "./UserActions.module.css";
import { useRouter } from "next/navigation";
import { useContext } from "react";
import { UserContext } from "../../contexts/UserContext";

const UserActions = () => {
  const { user, setUser } = useContext(UserContext);

  const router = useRouter();

  return user ? (
    <div className={styles.container}>
      <button className={styles.button} onClick={() => router.push("/profile")}>
        Profile
      </button>
      <button
        className={styles.button}
        onClick={() => {
          localStorage.removeItem("token");
          setUser(null);
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

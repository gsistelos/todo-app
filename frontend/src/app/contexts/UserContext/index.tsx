import { createContext } from "react";
import { User } from "../../types";
import jwt from "jsonwebtoken";

export const fetchUser = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      return null;
    }

    const data = jwt.decode(token);
    if (!data || typeof data === "string") {
      localStorage.removeItem("token");
      return null;
    }

    const response = await fetch("http://localhost:8080/api/users/" + data.id, {
      headers: {
        Authorization: "Bearer " + token,
      },
    });

    if (!response.ok || response.status !== 200) {
      localStorage.removeItem("token");
      return null;
    }

    return await response.json();
  } catch (error) {
    alert("An error occurred while fetching the user: " + error);
    return null;
  }
};

export type UserContextType = {
  user: User | null;
  setUser: (user: User | null) => void;
};

export const UserContext = createContext<UserContextType>({
  user: null,
  setUser: () => {},
});
